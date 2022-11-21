package compare

import (
	"path"
	"testing"

	"github.com/CrowdStrike/gqltools/utils"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

func TestCompareSchemaRoot(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Schema root query type changed",
			oldSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutation
			}`,
			newSchema: `
			schema {
			  query: RootQueryChanged
			  mutation: RootMutation
			}`,
			criticality: Breaking,
			ChangeType:  SchemaQueryTypeChanged,
		},
		{
			name:      "Schema with root query type added",
			oldSchema: ``,
			newSchema: `
			schema {
			  query: RootQuery
			}`,
			criticality: NonBreaking,
			ChangeType:  SchemaQueryTypeChanged,
		},
		{
			name: "Schema with root query type removed",
			oldSchema: `
			schema {
			  query: RootQuery
			}`,
			newSchema:   ``,
			criticality: Breaking,
			ChangeType:  SchemaQueryTypeChanged,
		},
		{
			name: "Schema root query type added",
			oldSchema: `
			schema {
			  mutation: RootMutation
			}`,
			newSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutation
			}`,
			criticality: NonBreaking,
			ChangeType:  SchemaQueryTypeChanged,
		},
		{
			name: "Schema root query type removed",
			oldSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutation
			}`,
			newSchema: `
			schema {
			  mutation: RootMutation
			}`,
			criticality: Breaking,
			ChangeType:  SchemaQueryTypeChanged,
		},
		{
			name: "Schema root mutation type changed",
			oldSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutation
			}
			`,
			newSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutationChanged
			}`,
			criticality: Breaking,
			ChangeType:  SchemaMutationTypeChanged,
		},
		{
			name: "Schema root mutation type added",
			oldSchema: `
			schema {
			  query: RootQuery
			}`,
			newSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutation
			}`,
			criticality: NonBreaking,
			ChangeType:  SchemaMutationTypeChanged,
		},
		{
			name: "Schema root mutation type removed",
			oldSchema: `
			schema {
			  query: RootQuery
			  mutation: RootMutation
			}`,
			newSchema: `
			schema {
			  query: RootQuery
			}`,
			criticality: Breaking,
			ChangeType:  SchemaMutationTypeChanged,
		},
		{
			name: "Schema root subscription type changed",
			oldSchema: `
			schema {
			  query: RootQuery
			  subscription: RootSubscription
			}`,
			newSchema: `
			schema {
			  query: RootQuery
			  subscription: RootSubscriptionChanged
			}`,
			criticality: Breaking,
			ChangeType:  SchemaSubscriptionTypeChanged,
		},
		{
			name: "Schema root subscription type added",
			oldSchema: `
			schema {
			  query: RootQuery
			}`,
			newSchema: `
			schema {
			  query: RootQuery
			  subscription: RootSubscription
			}`,
			criticality: NonBreaking,
			ChangeType:  SchemaSubscriptionTypeChanged,
		},
		{
			name: "Schema root subscription type removed",
			oldSchema: `
			schema {
			  query: RootQuery
			  subscription: RootSubscription
			}`,
			newSchema: `
			schema {
			  query: RootQuery
			}`,
			criticality: Breaking,
			ChangeType:  SchemaSubscriptionTypeChanged,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareTypes(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "A new type added",
			oldSchema: `
			type User {
				name : String!
			}
			`,
			newSchema: `
			type User {
				name : String!
			}
			type Book {
				name: String
			}
			`,
			criticality: NonBreaking,
			ChangeType:  TypeAdded,
		},
		{
			name: "An existing type removed",
			oldSchema: `
			type User {
				name : String!
			}
			type Book {
				name: String
			}
			`,
			newSchema: `
			type User {
				name : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  TypeRemoved,
		},
		{
			name: "An extended type added",
			oldSchema: `
			type User {
				name : String!
			}
			`,
			newSchema: `
			type User {
				name : String!
			}
			extend type Book {
				name: String
			}
			`,
			criticality: NonBreaking,
			ChangeType:  TypeAdded,
		},
		{
			name: "An existing extended type removed",
			oldSchema: `
			type User {
				name : String!
			}
			extend type Book {
				name: String
			}
			`,
			newSchema: `
			type User {
				name : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  TypeRemoved,
		},
		{
			name: "A type kind changed",
			oldSchema: `
			type User {
				name : String!
			}
			`,
			newSchema: `
			interface User {
				name : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  TypeKindChanged,
		},
		{
			name: "A type description changed",
			oldSchema: `
			"""
			User description
			"""
			type User {
				name : String!
			}
			`,
			newSchema: `
			"""
			User description changed
			"""
			type User {
				name : String!
			}
			`,
			criticality: NonBreaking,
			ChangeType:  TypeDescriptionChanged,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareObjectTypes(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Object type no longer implements interface",
			oldSchema: `
			type Employee implements User{
				name : String!
			}
			`,
			newSchema: `
			type Employee {
				name : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  ObjectTypeInterfaceRemoved,
		},
		{
			name: "Object type implements new interface",
			oldSchema: `
			type Employee {
				name : String!
			}
			`,
			newSchema: `
			type Employee implements User{
				name : String!
			}
			`,
			criticality: Dangerous,
			ChangeType:  ObjectTypeInterfaceAdded,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareFields(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Field type changed",
			oldSchema: `
			type Employee {
				name : String!
			}
			`,
			newSchema: `
			type Employee {
				name : Int!
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldTypeChanged,
		},
		{
			name: "Field type change from optional to required ",
			oldSchema: `
			type Employee {
				name : String
			}
			`,
			newSchema: `
			type Employee {
				name : String!
			}
			`,
			criticality: NonBreaking,
			ChangeType:  FieldTypeChanged,
		},
		{
			name: "Field type change from required to optional",
			oldSchema: `
			type Employee {
				name : String!
			}
			`,
			newSchema: `
			type Employee {
				name : String
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldTypeChanged,
		},
		{
			name: "Field type change from required list to optional list",
			oldSchema: `
			type Employee {
				name : [String!]
			}
			`,
			newSchema: `
			type Employee {
				name : [String]
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldTypeChanged,
		},
		{
			name: "Field description changed",
			oldSchema: `
			type Employee {
				"""
				field description
				"""
				name : String
			}
			`,
			newSchema: `
			type Employee {
				"""
				field description changed
				"""
				name : String
			}
			`,
			criticality: NonBreaking,
			ChangeType:  FieldDescriptionChanged,
		},
		{
			name: "Field deprecation added",
			oldSchema: `
			type Employee {
				name : String 
				newName: String!
			}
			`,
			newSchema: `
			type Employee {
				name : String @deprecated(reason: "use newName")
				newName: String!
			}
			`,
			criticality: Dangerous,
			ChangeType:  FieldDeprecationAdded,
		},
		{
			name: "Field deprecation removed",
			oldSchema: `
			type Employee {
				name : String @deprecated(reason: "some reason")
				newName: String!
			}
			`,
			newSchema: `
			type Employee {
				name : String 
				newName: String!
			}
			`,
			criticality: Dangerous,
			ChangeType:  FieldDeprecationRemoved,
		},
		{
			name: "Field deprecation reason changed",
			oldSchema: `
			type Employee {
				name : String @deprecated(reason: "some reason")
				newName: String!
			}
			`,
			newSchema: `
			type Employee {
				name : String  @deprecated(reason: "some reason changed")
				newName: String!
			}
			`,
			criticality: NonBreaking,
			ChangeType:  FieldDeprecationReasonChanged,
		},
		{
			name: "Field removed",
			oldSchema: `
			type Employee {
				name : String @deprecated(reason: "some reason")
				newName: String!
			}
			`,
			newSchema: `
			type Employee {
				newName: String!
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldRemoved,
		},
		{
			name: "Field added",
			oldSchema: `
			type Employee {
				name : String 
			}
			`,
			newSchema: `
			type Employee {
				name : String 
				newName: String!
			}
			`,
			criticality: NonBreaking,
			ChangeType:  FieldAdded,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareFieldArguments(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Field argument removed",
			oldSchema: `
			type Query {
				reviews(offset:Int, limit:Int) : String!
			}
			`,
			newSchema: `
			type Query {
				reviews(offset:Int) : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldArgumentRemoved,
		},
		{
			name: "Optional Argument added to field",
			oldSchema: `
			type Query {
				reviews(offset:Int) : String!
			}
			`,
			newSchema: `
			type Query {
				reviews(offset:Int, limit:Int) : String!
			}
			`,
			criticality: Dangerous,
			ChangeType:  FieldArgumentAdded,
		},
		{
			name: "required Argument added to field",
			oldSchema: `
			type Query {
				reviews(offset:Int) : String!
			}
			`,
			newSchema: `
			type Query {
				reviews(offset:Int, limit:Int!) : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldArgumentAdded,
		},
		{
			name: "Field Argument type changed",
			oldSchema: `
			type Query {
				reviews(offset:Int, limit:Int!) : String!
			}
			`,
			newSchema: `
			type Query {
				reviews(offset:Int, limit:Float!) : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldArgumentTypeChanged,
		},
		{
			name: "Field Argument type changed to required",
			oldSchema: `
			type Query {
				reviews(offset:Int, limit:Int) : String!
			}
			`,
			newSchema: `
			type Query {
				reviews(offset:Int, limit:Int!) : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  FieldArgumentTypeChanged,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareInputObjectFields(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Input object field removed",
			oldSchema: `
			input UserInput{
				id : ID!
				name : String!
			}
			`,
			newSchema: `
			input UserInput{
				id : ID!
			}
			`,
			criticality: Breaking,
			ChangeType:  InputFieldRemoved,
		},
		{
			name: "Optional input field Added",
			oldSchema: `
			input UserInput{
				id : ID!
			}
			`,
			newSchema: `
			input UserInput{
				id : ID!
				name : String
			}
			`,
			criticality: Dangerous,
			ChangeType:  InputFieldAdded,
		},
		{
			name: "Required input field added",
			oldSchema: `
			input UserInput{
				id : ID!
			}
			`,
			newSchema: `
			input UserInput{
				id : ID!
				name : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  InputFieldAdded,
		},
		{
			name: "Input field type changed to required",
			oldSchema: `
			input UserInput{
				id : ID!
				name : String
			}
			`,
			newSchema: `
			input UserInput{
				id : ID!
				name : String!
			}
			`,
			criticality: Breaking,
			ChangeType:  InputFieldTypeChanged,
		},
		{
			name: "Input field type changed",
			oldSchema: `
			input UserInput{
				id : ID!
				name : String
			}
			`,
			newSchema: `
			input UserInput{
				id : ID!
				name : Name!
			}
			`,
			criticality: Breaking,
			ChangeType:  InputFieldTypeChanged,
		},
		{
			name: "Input field type changed to optional",
			oldSchema: `
			input UserInput{
				addresses : [String]!
			}
			`,
			newSchema: `
			input UserInput{
				addresses : [String]
			}
			`,
			criticality: NonBreaking,
			ChangeType:  InputFieldTypeChanged,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareEnumType(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Enum value removed",
			oldSchema: `
			enum Color{
				RED
				GREEN
			}
			`,
			newSchema: `
			enum Color{
				RED
			}
			`,
			criticality: Breaking,
			ChangeType:  EnumValueRemoved,
		},
		{
			name: "Enum value added",
			oldSchema: `
			enum Color{
				RED
			}
			`,
			newSchema: `
			enum Color{
				RED
				GREEN
			}
			`,
			criticality: Dangerous,
			ChangeType:  EnumValueAdded,
		},
		{
			name: "Enum description changed",
			oldSchema: `
			enum Color{
				"""
				Enum value description
				"""
				RED
			}
			`,
			newSchema: `
			enum Color{
				"""
				Enum value description changed
				"""
				RED
			}
			`,
			criticality: NonBreaking,
			ChangeType:  EnumValueDescriptionChanged,
		},
		{
			name: "Enum deprecation added",
			oldSchema: `
			enum Color{
				RED
				GREEN 
			}
			`,
			newSchema: `
			enum Color{
				RED
				GREEN @deprecated(reason:"some reason")
			}
			`,
			criticality: Dangerous,
			ChangeType:  EnumValueDeprecationAdded,
		},
		{
			name: "Enum deprecation reason changed",
			oldSchema: `
			enum Color{
				RED
				GREEN @deprecated(reason:"some reason")
			}
			`,
			newSchema: `
			enum Color{
				RED
				GREEN @deprecated(reason:"some reason changed")
			}
			`,
			criticality: NonBreaking,
			ChangeType:  EnumValueDeprecationReasonChanged,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareUnions(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Union member removed",
			oldSchema: `
			union Body = Image | Text
			`,
			newSchema: `
			union Body = Image
			`,
			criticality: Breaking,
			ChangeType:  UnionMemberRemoved,
		},
		{
			name: "Object type implements new interface",
			oldSchema: `
			union Body = Image
			`,
			newSchema: `
			union Body = Image | Text
			`,
			criticality: Dangerous,
			ChangeType:  UnionMemberAdded,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareDirectives(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Directive removed",
			oldSchema: `
			directive @stream on FIELD
			directive @transform(from: String!) on FIELD
			`,
			newSchema: `
			directive @transform(from: String!) on FIELD
			`,
			criticality: Breaking,
			ChangeType:  DirectiveRemoved,
		},
		{
			name: "Directive added",
			oldSchema: `
			directive @stream on FIELD
			`,
			newSchema: `
			directive @stream on FIELD
			directive @transform(from: String!) on FIELD
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveAdded,
		},
		{
			name: "Directive description changed",
			oldSchema: `
			"""
			Some description
			"""
			directive @transform(from: String!) on FIELD
			`,
			newSchema: `
			"""
			Some description changed
			"""
			directive @transform(from: String!) on FIELD
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveDescriptionChanged,
		},
		{
			name: "Directive location removed",
			oldSchema: `
			directive @transform(from: String!) on OBJECT | INTERFACE
			`,
			newSchema: `
			directive @transform(from: String!) on OBJECT
			`,
			criticality: Breaking,
			ChangeType:  DirectiveLocationRemoved,
		},
		{
			name: "Directive location added",
			oldSchema: `
			directive @transform(from: String!) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String!) on OBJECT | INTERFACE
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveLocationAdded,
		},
		{
			name: "Directive repeatable removed",
			oldSchema: `
			directive @transform(from: String!) repeatable on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String!) on OBJECT 
			`,
			criticality: Breaking,
			ChangeType:  DirectiveRepeatableRemoved,
		},
		{
			name: "Directive repeatable added",
			oldSchema: `
			directive @transform(from: String!) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String!) repeatable on OBJECT 
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveRepeatableAdded,
		},
		{
			name: "Directive argument removed",
			oldSchema: `
			directive @transform(from: String, to:string) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String) on OBJECT 
			`,
			criticality: Breaking,
			ChangeType:  DirectiveArgumentRemoved,
		},
		{
			name: "Directive argument added",
			oldSchema: `
			directive @transform(from: String) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String, to:string) on OBJECT 
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveArgumentAdded,
		},
		{
			name: "Directive required argument added",
			oldSchema: `
			directive @transform(from: String) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String, to:string!) on OBJECT 
			`,
			criticality: Breaking,
			ChangeType:  DirectiveArgumentAdded,
		},
		{
			name: "Directive argument type changed",
			oldSchema: `
			directive @transform(from: String) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: Int) on OBJECT 
			`,
			criticality: Breaking,
			ChangeType:  DirectiveArgumentTypeChanged,
		},
		{
			name: "Directive argument type changed from required to optional",
			oldSchema: `
			directive @transform(from: String!) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String) on OBJECT 
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveArgumentTypeChanged,
		},
		{
			name: "Directive argument type changed from optional to required",
			oldSchema: `
			directive @transform(from: String) on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String!) on OBJECT 
			`,
			criticality: Breaking,
			ChangeType:  DirectiveArgumentTypeChanged,
		},
		{
			name: "Directive argument default value changed",
			oldSchema: `
			directive @transform(from: String = "value") on OBJECT 
			`,
			newSchema: `
			directive @transform(from: String = "value changed") on OBJECT 
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveArgumentDefaultValueChanged,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareTypeDirectives(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Type directive removed",
			oldSchema: `
			type Book @key(fields: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book {
				isbn: String!
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveRemoved,
		},
		{
			name: "Type directive added",
			oldSchema: `
			type Book  {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book @key(field1: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveAdded,
		},
		{
			name: "Type directive argument value changed",
			oldSchema: `
			type Book @key(fields: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book @key(fields: "isbn title") {
				isbn: String!
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveArgumentValueChanged,
		},
		{
			name: "type directive argument removed",
			oldSchema: `
			type Book @key(field1: "isbn", field2: "title") {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book @key(field1: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveArgumentRemoved,
		},
		{
			name: "Type directive argument added",
			oldSchema: `
			type Book @key(field1: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book @key(field1: "isbn", field2: "title") {
				isbn: String!
				title: String 
			}
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveArgumentAdded,
		},
		{
			name: "one of the repetitive directive is removed ",
			oldSchema: `
			type Book 
			  @graph(type:"book", key: "isbn") 
			  @graph(type:"library", key: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book 
			  @graph(type:"book", key: "isbn")  {
				isbn: String!
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveChanged,
		},
		{
			name: "repetitive directive is changed ",
			oldSchema: `
			type Book 
			  @graph(type:"book", key: "isbn") 
			  @graph(type:"library", key: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book 
			  @graph(type:"book", key: "isbn")  
			  @graph(type:"user", key: "isbn") {
				isbn: String!
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveChanged,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareFieldDirectives(t *testing.T) {
	tests := []struct {
		name        string
		oldSchema   string
		newSchema   string
		criticality Criticality
		ChangeType  ChangeType
	}{
		{
			name: "Field directive removed",
			oldSchema: `
			type Book {
				isbn: String! @exposure(scope: [PARTNER])
				title: String 
			}
			`,
			newSchema: `
			type Book {
				isbn: String!
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveRemoved,
		},
		{
			name: "Field directive added",
			oldSchema: `
			type Book  {
				isbn: String!
				title: String 
			}
			`,
			newSchema: `
			type Book {
				isbn: String! @exposure(scope: [PARTNER])
				title: String 
			}
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveAdded,
		},
		{
			name: "Field directive argument value changed",
			oldSchema: `
			type Book {
				isbn: String! @exposure(scope: [PARTNER, PUBLIC])
				title: String 
			}
			`,
			newSchema: `
			type Book  {
				isbn: String! @exposure(scope: [PARTNER])
				title: String
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveArgumentValueChanged,
		},
		{
			name: "Field directive argument removed",
			oldSchema: `
			type Book {
				isbn: String! @directive(field1: "name", field2: "title")
				title: String 
			}
			`,
			newSchema: `
			type Book {
				isbn: String! @directive(field1: "name")
				title: String 
			}
			`,
			criticality: Dangerous,
			ChangeType:  DirectiveArgumentRemoved,
		},
		{
			name: "Field directive argument added",
			oldSchema: `
			type Book {
				isbn: String! @directive(field1: "name")
				title: String 
			}
			`,
			newSchema: `
			type Book {
				isbn: String! @directive(field1: "name", field2: "title")
				title: String 
			}
			`,
			criticality: NonBreaking,
			ChangeType:  DirectiveArgumentAdded,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			oldSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.oldSchema,
			})
			if err != nil {
				t.Fatalf("error parsing old schema, error = %v", err)
			}
			newSchema, err := parser.ParseSchema(&ast.Source{
				Input: tt.newSchema,
			})
			if err != nil {
				t.Fatalf("error parsing new schema, error = %v", err)
			}
			changes := FindChangesInSchemas(oldSchema, newSchema)
			if len(changes) != 1 {
				t.Errorf("Unexpected changes added = %v", changes)
			}
			if changes[0].criticalityLevel != tt.criticality || changes[0].changeType != tt.ChangeType {
				t.Errorf("Object type changes = %v", changes[0])
			}
		})
	}
}

func TestCompareSchemaFiles(t *testing.T) {
	t.Run("Compare schema files", func(t *testing.T) {
		sourceDir := "./test_schema"
		schemaOldContents, err := utils.ReadFiles(path.Join(sourceDir, "oldSchema.graphql"))
		if err != nil {
			t.Errorf("error reading oldSchema = %v", err)
		}

		schemaNewContents, err := utils.ReadFiles(path.Join(sourceDir, "newSchema.graphql"))
		if err != nil {
			t.Errorf("error reading newSchema = %v", err)
		}

		schemaOld, parseErr := utils.ParseSchema(schemaOldContents)
		if parseErr != nil {
			t.Errorf("error parsing oldSchema = %v", err)
		}

		schemaNew, parseErr := utils.ParseSchema(schemaNewContents)
		if parseErr != nil {
			t.Errorf("error reading newSchema = %v", err)
		}

		changes := FindChangesInSchemas(schemaOld, schemaNew)
		if len(changes) != 7 {
			t.Errorf("Unexpected changes added = %v", changes)
		}
	})
}
