package linter

import (
	"testing"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

func TestArgumentsHaveDescription(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"arguments_without_description",
			`
			type Query {
  				todos(offset: Int, limit: Int): [Todo!]!
			}
			`,
			true,
		},
		{
			"arguments_in_extended_type_without_description",
			`
			extend type Query {
  				todos(offset: Int, limit: Int): [Todo!]!
			}
			`,
			true,
		},
		{
			"arguments_in_extended_type_with_description",
			`
			extend type Query {
  				todos(
					"some comment about offset"
					offset: Int, 
					"some comment about limit"
					limit: Int
				): [Todo!]!
			}
			`,
			false,
		},
		{
			"arguments_with_description",
			`
			type Query {
  				todos(
					"some comment about offset"
					offset: Int, 
					"some comment about limit"
					limit: Int
				): [Todo!]!
			}
			`,
			false,
		},
		{
			"one_argument_with_description",
			`	
			type Query {
  				todos(
					"some comment about offset"
					offset: Int, 
					limit: Int
				): [Todo!]!
			}
			`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("ArgumentsHaveDescription() invalid input; error = %v", parseErr)
			}
			if errs := ArgumentsHaveDescription(schemaDoc); (errs.Len() != 0) != tt.wantErr {
				t.Errorf("ArgumentsHaveDescription() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestEnumValuesAreAllCaps(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"enum_value_are_lowercase",
			`enum color { red, blue, green }`,
			true,
		},
		{
			"enum_value_are_titlecase",
			`enum color { Red, Blue, Green }`,
			true,
		},
		{
			"enum_value_are_uppercase",
			`enum color { RED, BLUE, GREEN }`,
			false,
		},
		{
			"enum_value_are_mixed",
			`enum color { RED, Blue, green }`,
			true,
		},
		{
			"extended_enum_value_are_uppercase",
			`extend enum color { RED, BLUE, GREEN }`,
			false,
		},
		{
			"extended_enum_value_are_mixed",
			`extend enum color { RED, Blue, green }`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("EnumValuesAreAllCaps() invalid input; error = %v", parseErr)
			}
			if errs := EnumValuesAreAllCaps(schemaDoc); (errs.Len() != 0) != tt.wantErr {
				t.Errorf("EnumValuesAreAllCaps() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestEnumValuesHaveDescriptions(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"enum_values_with_description",
			`
			enum color {
				"color red"
				red,  
				"color blue"
				blue, 
				"color green"
				green
			}`,
			false,
		},
		{
			"enum_values_without_description",
			`
			enum color {
				red,  
				blue, 
				green 
			}`,
			true,
		},
		{
			"some_enum_values_with_description",
			`
			enum color {
				red,
				"color blue"
				blue, 
				"color green"
				green
			}`,
			true,
		},
		{
			"extended_enum_values_with_description",
			`
			extend enum color {
				"crowdstrike is red "
				red,  
				"salesforce is blue"
				blue, 
				"splunk is green"
				green
			}`,
			false,
		},
		{
			"extended_enum_values_without_description",
			`
			extend enum color {
				red,  
				blue, 
				green 
			}`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("EnumValuesHaveDescriptions() invalid input; error = %v", parseErr)
			}
			if errs := EnumValuesHaveDescriptions(schemaDoc); (errs.Len() != 0) != tt.wantErr {
				t.Errorf("EnumValuesHaveDescriptions() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestFieldsAreCamelCased(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"fields_are_camelCase",
			`type Query {
  				todos: [Todo!]!
				todoList: [Todo!]!
			}`,
			false,
		},
		{
			"fields_are_TitleCase",
			`type Query {
				TodoList: [Todo!]!
			}`,
			true,
		},
		{
			"fields_are_UPPERCASE",
			`type Query {
				TODOS: [Todo!]!
			}`,
			true,
		},
		{
			"extended_type_fields_are_camelCase",
			`extend type Query {
				todoList: [Todo!]!
			}`,
			false,
		},
		{
			"extended_type_fields_are_TitleCase",
			`extend type Query {
				TodoList: [Todo!]!
			}`,
			true,
		},
		{
			"extended_type_fields_are_UPPERCASE",
			`extend type Query {
				TODOS: [Todo!]!
			}`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("FieldsAreCamelCased() invalid input; error = %v", parseErr)
			}
			if errs := FieldsAreCamelCased(schemaDoc); (errs.Len() != 0) != tt.wantErr {
				t.Errorf("FieldsAreCamelCased() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestFieldsHaveDescription(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"fields_without_description",
			`
			type Query {
  				todos(offset: Int, limit: Int): [Todo!]!
			}
			`,
			true,
		},
		{
			"fields_with_description",
			`
			type Query {
				"todos return list of todos"
  				todos(offset: Int, limit: Int): [Todo!]!
			}
			`,
			false,
		},
		{
			"fields_of_extended_type_without_description",
			`
			extend type Query {
  				todos(offset: Int, limit: Int): [Todo!]!
			}
		`,
			true,
		},
		{
			"fields_of_extended_type_with_description",
			`
			extend type Query {
				"todos return list of todos"
  				todos(offset: Int, limit: Int): [Todo!]!
			}
			`,
			false,
		},
		{
			"one_of_the_field_without_description",
			`	
			type Query {
  				"todos return list of todos"
  				todos(offset: Int, limit: Int): [Todo!]!
  				todoList(offset: Int, limit: Int): [Todo!]!
			}
			`,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("FieldsHaveDescription() invalid input; error = %v", parseErr)
			}
			if errs := FieldsHaveDescription(schemaDoc); (errs.Len() != 0) != tt.wantErr {
				t.Errorf("FieldsHaveDescription() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestTypesAreCapitalized(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"types_are_Capitalized",
			`type User {
  				name: String!
			}`,
			false,
		},
		{
			"types_are_lowercased",
			`type user {
  				name: String!
			}`,
			true,
		},
		{
			"types_are_mixed",
			`type User {
  				name: String!
			}
			type query {
  				me: User!
			}`,
			true,
		},
		{
			"extended_type_is_lowercased",
			`extend type user {
  				name: String!
			}`,
			true,
		},
		{
			"extended_types_is_Titlecased",
			`extend type User {
  				name: String!
			}`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("TypesAreCapitalized() invalid input; error = %v", parseErr)
			}
			if errs := TypesAreCapitalized(schemaDoc); (errs.Len() != 0) != tt.wantErr {
				t.Errorf("TypesAreCapitalized() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestTypesHaveDescription(t *testing.T) {
	t.Run("something", func(t *testing.T) {
		tests := []struct {
			name    string
			schema  string
			wantErr bool
		}{
			{
				"type_without_description",
				`
			type User {
  				name: String!
			}
			`,
				true,
			},
			{
				"type_with_description",
				`
			"User represent the person calling the endpoint"
			type User {
  				name: String!
			}
			`,
				false,
			},
			{
				"extended_type_without_description",
				`
			extend type User {
  				name: String!
			}
			`,
				false,
			},
			{
				"some_of_the_types_without_description",
				`
			"User represent the person calling the endpoint"
			type User {
  				name: String!
			}
			type Query {
  				me: User!
			}
			`,
				true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
					Input: tt.schema,
				})
				if parseErr != nil {
					t.Fatalf("TypesHaveDescription() invalid input; error = %v", parseErr)
				}
				if errs := TypesHaveDescription(schemaDoc); (errs.Len() != 0) != tt.wantErr {
					t.Errorf("TypesHaveDescription() error = %v, wantErr %v", errs, tt.wantErr)
				}
			})
		}
	})
}

func TestRelayConnectionTypesSpec(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"type_connection_non_object_type",
			`
			interface UserConnection {
				edges: [User]
				pageInfo: PageInfo!
			}
			`,
			true,
		},
		{
			"type_connection_without_edges_and_pageinfo_fields",
			`
			type UserConnection {
  				id: Int!
			}
			`,
			true,
		},
		{
			"type_connection_with_non_list_edges_type",
			`
			type UserConnection {
  				edges: String
				pageInfo: PageInfo!
			}
			`,
			true,
		},
		{
			"type_connection_with_nullable_pageinfo_type_for_pageinfo_field",
			`
			type UserConnection {
  				edges: [SomeObject]
				pageInfo: PageInfo
			}
			`,
			true,
		},
		{
			"type_connection_with_nonpageinfo_type_for_pageinfo_field",
			`
			type UserConnection {
  				edges: [SomeObject]
				pageInfo: String!
			}
			`,
			true,
		},
		{
			"type_connection_with_valid_edges_pageinfo_fields",
			`
			type UserConnection {
  				edges: [SomeObject]
				pageInfo: PageInfo!
			}
			`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("RelayConnectionTypesSpec() invalid input; error = %v", parseErr)
			}
			if errs := RelayConnectionTypesSpec(schemaDoc); (errs.Len() > 0) != tt.wantErr {
				t.Errorf("RelayConnectionTypesSpec() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}

func TestRelayConnectionArgumentsSpec(t *testing.T) {
	tests := []struct {
		name    string
		schema  string
		wantErr bool
	}{
		{
			"field_connection_without_forward_or_backward_pagination",
			`
			type User {
  				result: UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_forward_and_backward_pagination_and_non_nullable_first_argument",
			`
			type User {
  				result(first: Int!, after: String, last: Int, before: String): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_forward_and_backward_pagination_and_non_nullable_last_argument",
			`
			type User {
  				result(first: Int, after: String, last: Int!, before: String): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_non_int_first_argument",
			`
			type User {
  				result(first: String, after: String): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_non_int_last_argument",
			`
			type User {
  				result(last: String, before: String): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_only_first_argument",
			`
			type User {
  				result(first: Int): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_only_first_and_last_arguments",
			`
			type User {
  				result(first: Int, last: Int): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_only_first_and_before_arguments",
			`
			type User {
  				result(first: Int, before: String): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_only_last_and_after_arguments",
			`
			type User {
  				result(last: Int, after: String): UserConnection
			}
			`,
			true,
		},
		{
			"field_connection_with_valid_forward_pagination",
			`
			type User {
  				result(first: Int, after: String): UserConnection
			}
			`,
			false,
		},
		{
			"field_connection_with_valid_forward_pagination_with_nullable_first_argument",
			`
			type User {
  				result(first: Int!, after: String): UserConnection
			}
			`,
			false,
		},
		{
			"field_connection_with_valid_backward_pagination",
			`
			type User {
  				result(last: Int, before: String): UserConnection
			}
			`,
			false,
		},
		{
			"field_connection_with_valid_backward_pagination_with_nullable_last_argument",
			`
			type User {
  				result(last: Int!, before: String): UserConnection
			}
			`,
			false,
		},
		{
			"field_connection_with_valid_forward_and_backward_pagination",
			`
			type User {
  				result(first: Int, after: String, last: Int, before: String): UserConnection
			}
			`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaDoc, parseErr := parser.ParseSchema(&ast.Source{
				Input: tt.schema,
			})
			if parseErr != nil {
				t.Fatalf("RelayConnectionArgumentsSpec() invalid input; error = %v", parseErr)
			}
			if errs := RelayConnectionArgumentsSpec(schemaDoc); (errs.Len() > 0) != tt.wantErr {
				t.Errorf("RelayConnectionTypesSpec() error = %v, wantErr %v", errs, tt.wantErr)
			}
		})
	}
}
