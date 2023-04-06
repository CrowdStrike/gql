package compare

import (
	"fmt"
	"sort"

	"github.com/vektah/gqlparser/v2/ast"
)

//const (
//	breakingIcon, dangerousIcon, nonbreakingIcon = ":x:", ":warning:", ":white_check_mark:"
//)

// ChangeType enum to list all type of breaking/non-breaking/dangerous changes
type ChangeType string

const (
	// FieldArgumentDescriptionChanged Field Argument Description Changed
	FieldArgumentDescriptionChanged ChangeType = "FIELD_ARGUMENT_DESCRIPTION_CHANGED"
	// FieldArgumentDefaultChanged Field Argument Default Changed
	FieldArgumentDefaultChanged ChangeType = "FIELD_ARGUMENT_DEFAULT_CHANGED"
	// FieldArgumentTypeChanged Field Argument Type Changed
	FieldArgumentTypeChanged ChangeType = "FIELD_ARGUMENT_TYPE_CHANGED"
	// DirectiveRemoved Directive Removed
	DirectiveRemoved ChangeType = "DIRECTIVE_REMOVED"
	// DirectiveChanged Directive changed
	DirectiveChanged ChangeType = "DIRECTIVE_CHANGED"
	// DirectiveAdded Directive Added
	DirectiveAdded ChangeType = "DIRECTIVE_ADDED"
	// DirectiveDescriptionChanged Directive Description Changed
	DirectiveDescriptionChanged ChangeType = "DIRECTIVE_DESCRIPTION_CHANGED"
	// DirectiveLocationAdded Directive Location Added
	DirectiveLocationAdded ChangeType = "DIRECTIVE_LOCATION_ADDED"
	// DirectiveLocationRemoved Directive Location Removed
	DirectiveLocationRemoved ChangeType = "DIRECTIVE_LOCATION_REMOVED"
	// DirectiveArgumentAdded Directive Argument Added
	DirectiveArgumentAdded ChangeType = "DIRECTIVE_ARGUMENT_ADDED"
	// DirectiveArgumentRemoved Directive Argument Removed
	DirectiveArgumentRemoved ChangeType = "DIRECTIVE_ARGUMENT_REMOVED"
	// DirectiveArgumentDescriptionChanged Directive Argument Description Changed
	DirectiveArgumentDescriptionChanged ChangeType = "DIRECTIVE_ARGUMENT_DESCRIPTION_CHANGED"
	// DirectiveArgumentDefaultValueChanged Directive Argument Default Value Changed
	DirectiveArgumentDefaultValueChanged ChangeType = "DIRECTIVE_ARGUMENT_DEFAULT_VALUE_CHANGED"
	// DirectiveArgumentTypeChanged Directive Argument Type Changed
	DirectiveArgumentTypeChanged ChangeType = "DIRECTIVE_ARGUMENT_TYPE_CHANGED"
	// DirectiveRepeatableRemoved Directive Repeatable Removed
	DirectiveRepeatableRemoved ChangeType = "DIRECTIVE_REPEATABLE_REMOVED"
	// DirectiveRepeatableAdded Directive Repeatable Added
	DirectiveRepeatableAdded ChangeType = "DIRECTIVE_REPEATABLE_ADDED"
	// DirectiveArgumentValueChanged Directive Argument Value Changed
	DirectiveArgumentValueChanged ChangeType = "DIRECTIVE_ARGUMENT_VALUE_CHANGED"
	// EnumValueRemoved Enum Value Removed
	EnumValueRemoved ChangeType = "ENUM_VALUE_REMOVED"
	// EnumValueAdded Enum Value Added
	EnumValueAdded ChangeType = "ENUM_VALUE_ADDED"
	// EnumValueDescriptionChanged Enum Value Description Changed
	EnumValueDescriptionChanged ChangeType = "ENUM_VALUE_DESCRIPTION_CHANGED"
	// EnumValueDeprecationReasonChanged Enum Value Deprecation Reason Changed
	EnumValueDeprecationReasonChanged ChangeType = "ENUM_VALUE_DEPRECATION_REASON_CHANGED"
	// EnumValueDeprecationAdded Enum Value Deprecation Added
	EnumValueDeprecationAdded ChangeType = "ENUM_VALUE_DEPRECATION_ADDED"
	// FieldRemoved Field Removed
	FieldRemoved ChangeType = "FIELD_REMOVED"
	// FieldAdded Field Added
	FieldAdded ChangeType = "FIELD_ADDED"
	// FieldDescriptionChanged Field Description Changed
	FieldDescriptionChanged ChangeType = "FIELD_DESCRIPTION_CHANGED"
	// FieldDeprecationAdded Field Deprecation Added
	FieldDeprecationAdded ChangeType = "FIELD_DEPRECATION_ADDED"
	// FieldDeprecationRemoved Field Deprecation Removed
	FieldDeprecationRemoved ChangeType = "FIELD_DEPRECATION_REMOVED"
	// FieldDeprecationReasonChanged Field Deprecation Reason Changed
	FieldDeprecationReasonChanged ChangeType = "FIELD_DEPRECATION_REASON_CHANGED"
	// FieldTypeChanged Field Type Changed
	FieldTypeChanged ChangeType = "FIELD_TYPE_CHANGED"
	// FieldArgumentAdded Field Argument Added
	FieldArgumentAdded ChangeType = "FIELD_ARGUMENT_ADDED"
	// FieldArgumentRemoved Field Argument Removed
	FieldArgumentRemoved ChangeType = "FIELD_ARGUMENT_REMOVED"
	// InputFieldRemoved Input Field Removed
	InputFieldRemoved ChangeType = "INPUT_FIELD_REMOVED"
	// InputFieldAdded Input Field Added
	InputFieldAdded ChangeType = "INPUT_FIELD_ADDED"
	// InputFieldDescriptionChanged Input Field Description Changed
	InputFieldDescriptionChanged ChangeType = "INPUT_FIELD_DESCRIPTION_CHANGED"
	// InputFieldDefaultValueChanged Input Field Default Value Changed
	InputFieldDefaultValueChanged ChangeType = "INPUT_FIELD_DEFAULT_VALUE_CHANGED"
	// InputFieldTypeChanged Input Field Type Changed
	InputFieldTypeChanged ChangeType = "INPUT_FIELD_TYPE_CHANGED"
	// ObjectTypeInterfaceAdded Object Type Interface Added
	ObjectTypeInterfaceAdded ChangeType = "OBJECT_TYPE_INTERFACE_ADDED"
	// ObjectTypeInterfaceRemoved Object Type Interface Removed
	ObjectTypeInterfaceRemoved ChangeType = "OBJECT_TYPE_INTERFACE_REMOVED"
	// SchemaQueryTypeChanged Schema Query Type Changed
	SchemaQueryTypeChanged ChangeType = "SCHEMA_QUERY_TYPE_CHANGED"
	// SchemaMutationTypeChanged Schema Mutation Type Changed
	SchemaMutationTypeChanged ChangeType = "SCHEMA_MUTATION_TYPE_CHANGED"
	// SchemaSubscriptionTypeChanged Schema Subscription Type Changed
	SchemaSubscriptionTypeChanged ChangeType = "SCHEMA_SUBSCRIPTION_TYPE_CHANGED"
	// TypeRemoved Type Removed
	TypeRemoved ChangeType = "TYPE_REMOVED"
	// TypeAdded Type Added
	TypeAdded ChangeType = "TYPE_ADDED"
	// TypeKindChanged Type Kind Changed
	TypeKindChanged ChangeType = "TYPE_KIND_CHANGED"
	// TypeDescriptionChanged Type Description Changed
	TypeDescriptionChanged ChangeType = "TYPE_DESCRIPTION_CHANGED"
	// UnionMemberRemoved Union Member Removed
	UnionMemberRemoved ChangeType = "UNION_MEMBER_REMOVED"
	// UnionMemberAdded Union Member Added
	UnionMemberAdded ChangeType = "UNION_MEMBER_ADDED"
)

// Criticality severity of a change in schema
type Criticality int

const (
	// NonBreaking Change is compatible with previous version
	NonBreaking Criticality = 0
	// Dangerous Change is compatible with previous version but can result in unexpected behavior for consumer
	Dangerous Criticality = 1
	// Breaking Change is incompatible with previous version
	Breaking Criticality = 2
)

const deprecatedDirective = "deprecated"

// Change defines a change in schema
type Change struct {
	message          string
	changeType       ChangeType
	criticalityLevel Criticality
	path             string
	position         *ast.Position
}

// GetPosition get change position
func (c *Change) GetPosition() *ast.Position {
	return c.position
}

// GetMessage get change message
func (c *Change) GetMessage() string {
	return c.message
}

// GetChangeCriticalityLevel get change criticality level
func (c *Change) GetChangeCriticalityLevel() Criticality {
	return c.criticalityLevel
}

// GetChangeType get change type
func (c *Change) GetChangeType() ChangeType {
	return c.changeType
}

// GetPath get change path
func (c *Change) GetPath() string {
	return c.path
}

// string get change criticality level string
func (c Criticality) String() string {
	switch c {
	case Breaking:
		return "Breaking"
	case Dangerous:
		return "Dangerous"
	case NonBreaking:
		return "NonBreaking"
	default:
		return ""
	}
}

// FindChangesInSchemas compares two schemas, returns the list of all changes made in the second schema
func FindChangesInSchemas(oldSchema *ast.SchemaDocument, newSchema *ast.SchemaDocument) []*Change {
	var changes []*Change
	changes = []*Change{}
	changes = append(changes, changeInSchema(oldSchema.Schema, newSchema.Schema)...)
	changes = append(changes, changeInSchema(oldSchema.SchemaExtension, newSchema.SchemaExtension)...)
	changes = append(changes, changeInTypes(oldSchema, newSchema)...)
	changes = append(changes, changeInDirective(oldSchema.Directives, newSchema.Directives)...)
	return changes
}

// changeInSchema change in schema root operations
func changeInSchema(oldSchemaDefs ast.SchemaDefinitionList, newSchemaDefs ast.SchemaDefinitionList) []*Change {
	var changes []*Change
	if len(oldSchemaDefs) == 0 && len(newSchemaDefs) == 0 {
		return changes
	}
	if len(oldSchemaDefs) > 0 && len(newSchemaDefs) == 0 {
		oldQuery := getOperationForName(oldSchemaDefs[0].OperationTypes, ast.Query)
		changes = append(changes, changesInSchemaOperation(oldQuery, nil, ast.Query)...)

		oldMutation := getOperationForName(oldSchemaDefs[0].OperationTypes, ast.Mutation)
		changes = append(changes, changesInSchemaOperation(oldMutation, nil, ast.Mutation)...)

		oldSubscription := getOperationForName(oldSchemaDefs[0].OperationTypes, ast.Subscription)
		changes = append(changes, changesInSchemaOperation(oldSubscription, nil, ast.Subscription)...)
		return changes
	}
	if len(oldSchemaDefs) == 0 && len(newSchemaDefs) > 0 {
		newQuery := getOperationForName(newSchemaDefs[0].OperationTypes, ast.Query)
		changes = append(changes, changesInSchemaOperation(nil, newQuery, ast.Query)...)

		newMutation := getOperationForName(newSchemaDefs[0].OperationTypes, ast.Mutation)
		changes = append(changes, changesInSchemaOperation(nil, newMutation, ast.Mutation)...)

		newSubscription := getOperationForName(newSchemaDefs[0].OperationTypes, ast.Subscription)
		changes = append(changes, changesInSchemaOperation(nil, newSubscription, ast.Subscription)...)
		return changes
	}

	oldQuery := getOperationForName(oldSchemaDefs[0].OperationTypes, ast.Query)
	newQuery := getOperationForName(newSchemaDefs[0].OperationTypes, ast.Query)
	changes = append(changes, changesInSchemaOperation(oldQuery, newQuery, ast.Query)...)

	oldMutation := getOperationForName(oldSchemaDefs[0].OperationTypes, ast.Mutation)
	newMutation := getOperationForName(newSchemaDefs[0].OperationTypes, ast.Mutation)
	changes = append(changes, changesInSchemaOperation(oldMutation, newMutation, ast.Mutation)...)

	oldSubscription := getOperationForName(oldSchemaDefs[0].OperationTypes, ast.Subscription)
	newSubscription := getOperationForName(newSchemaDefs[0].OperationTypes, ast.Subscription)
	changes = append(changes, changesInSchemaOperation(oldSubscription, newSubscription, ast.Subscription)...)

	return changes
}

func changesInSchemaOperation(oldOp *ast.OperationTypeDefinition, newOp *ast.OperationTypeDefinition, op ast.Operation) []*Change {
	var changes []*Change
	var changeType ChangeType
	switch op {
	case ast.Query:
		changeType = SchemaQueryTypeChanged
	case ast.Mutation:
		changeType = SchemaMutationTypeChanged
	case ast.Subscription:
		changeType = SchemaSubscriptionTypeChanged
	}
	if oldOp == nil && newOp != nil {
		changes = append(changes, &Change{
			changeType:       changeType,
			criticalityLevel: NonBreaking,
			message:          fmt.Sprintf("Schema %s root has added '%s'", op, newOp.Operation),
			position:         newOp.Position,
		})
	}
	if oldOp != nil && newOp == nil {
		changes = append(changes, &Change{
			changeType:       changeType,
			criticalityLevel: Breaking,
			message:          fmt.Sprintf("Schema %s root has removed '%s'", op, oldOp.Operation),
			position:         oldOp.Position,
		})
	}
	if oldOp != nil && newOp != nil && oldOp.Type != newOp.Type {
		changes = append(changes, &Change{
			changeType:       changeType,
			criticalityLevel: Breaking,
			message:          fmt.Sprintf("Schema %s root has changed from '%s' to '%s'", op, oldOp.Operation, newOp.Operation),
			position:         newOp.Position,
		})
	}
	return changes
}

// changeInTypes change in al the types
func changeInTypes(oldSchema *ast.SchemaDocument, newSchema *ast.SchemaDocument) []*Change {
	var changes []*Change
	persistedType := map[string][]*ast.Definition{}
	//Check if types added/removed/persisted
	changes = append(changes, checkTypeRemoved(oldSchema.Definitions, newSchema.Definitions, persistedType, false)...)
	changes = append(changes, checkTypeAdded(oldSchema.Definitions, newSchema.Definitions, false)...)

	//Check if extended types added/removed/persisted
	changes = append(changes, checkTypeRemoved(oldSchema.Extensions, newSchema.Extensions, persistedType, true)...)
	changes = append(changes, checkTypeAdded(oldSchema.Extensions, newSchema.Extensions, true)...)

	//compare persisted types for changes in fields/directives
	for _, defs := range persistedType {
		ot := defs[0]
		nt := defs[1]
		if ot.Kind == ast.Enum && nt.Kind == ast.Enum {
			changes = append(changes, changeInEnum(ot, nt)...)
		}
		if ot.Kind == ast.InputObject && nt.Kind == ast.InputObject {
			changes = append(changes, changeInInputFields(ot, nt)...)
		}
		if ot.Kind == ast.Interface && nt.Kind == ast.Interface {
			changes = append(changes, changeInTypeFieldDirectives(ot.Directives, nt.Directives, nt.Name, nt.Position)...)
			changes = append(changes, changeInFields(ot, nt)...)
		}
		if ot.Kind == ast.Object && nt.Kind == ast.Object {
			changes = append(changes, changeInObject(ot, nt)...)
		}
		if ot.Kind == ast.Union && nt.Kind == ast.Union {
			changes = append(changes, changeInUnion(ot, nt)...)
		}
		if ot.Kind != nt.Kind {
			//Changing the kind of a type is a breaking change because it can cause existing queries to error.
			//For example, turning an object type to a scalar type would break queries that define a selection set for this type.
			changes = append(changes, &Change{
				changeType:       TypeKindChanged,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("Type '%s' kind changed from '%s' to '%s'", ot.Name, ot.Kind, nt.Kind),
				path:             nt.Name,
				position:         nt.Position,
			})
		}
		if ot.Description != nt.Description {
			changes = append(changes, &Change{
				changeType:       TypeDescriptionChanged,
				criticalityLevel: NonBreaking,
				message:          fmt.Sprintf("Type '%s' description changed", ot.Name),
				path:             nt.Name,
				position:         nt.Position,
			})
		}
	}
	return changes
}

func checkTypeRemoved(oldSchemaDefs ast.DefinitionList, newSchemaDefs ast.DefinitionList, persistedType map[string][]*ast.Definition, isExtended bool) []*Change {
	var changes []*Change
	for _, ot := range oldSchemaDefs {
		nt := newSchemaDefs.ForName(ot.Name)
		if nt == nil {
			msg := fmt.Sprintf("Type '%s' was removed", ot.Name)
			if isExtended {
				msg = fmt.Sprintf("Extended type '%s' was removed", ot.Name)
			}
			//removing a type from schema is a breaking change
			changes = append(changes, &Change{
				changeType:       TypeRemoved,
				criticalityLevel: Breaking,
				message:          msg,
				path:             ot.Name,
				position:         ot.Position,
			})
		} else {
			persistedType[ot.Name] = []*ast.Definition{ot, nt}
		}
	}
	return changes
}

func checkTypeAdded(oldSchemaDefs ast.DefinitionList, newSchemaDefs ast.DefinitionList, isExtended bool) []*Change {
	var changes []*Change
	for _, nt := range newSchemaDefs {
		ot := oldSchemaDefs.ForName(nt.Name)
		if ot == nil {
			msg := fmt.Sprintf("Type '%s' was added", nt.Name)
			if isExtended {
				msg = fmt.Sprintf("Extended type '%s' was added", nt.Name)
			}
			//type added to new schema
			changes = append(changes, &Change{
				changeType:       TypeAdded,
				criticalityLevel: NonBreaking,
				message:          msg,
				path:             nt.Name,
				position:         nt.Position,
			})
		}
	}
	return changes
}

// changeInDirective change in the directive definitions
func changeInDirective(oDirs ast.DirectiveDefinitionList, nDirs ast.DirectiveDefinitionList) []*Change {
	var changes []*Change
	for _, od := range oDirs {
		nd := nDirs.ForName(od.Name)
		if nd == nil {
			changes = append(changes, &Change{
				changeType:       DirectiveRemoved,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("Directive '@%s' was removed ", od.Name),
				path:             fmt.Sprintf("@%s", od.Name),
				position:         od.Position,
			})
		} else {
			//description changed
			if od.Description != nd.Description {
				changes = append(changes, &Change{
					changeType:       DirectiveDescriptionChanged,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Directive '@%s' description changed ", od.Name),
					path:             fmt.Sprintf("@%s", nd.Name),
					position:         nd.Position,
				})
			}
			changes = append(changes, checkDirectiveLocationChanged(od, nd)...)
			changes = append(changes, checkDirectiveRepeatableChanged(od, nd)...)
			//argument changed
			changes = append(changes, changeInDirectiveArguments(od, nd)...)
			changes = append(changes, checkDirectiveArgumentAdded(od, nd)...)
		}
	}
	changes = append(changes, checkDirectiveAdded(oDirs, nDirs)...)
	return changes
}

func checkDirectiveArgumentAdded(od *ast.DirectiveDefinition, nd *ast.DirectiveDefinition) []*Change {
	var changes []*Change
	for _, nArg := range nd.Arguments {
		oArg := od.Arguments.ForName(nArg.Name)
		if oArg == nil {
			//argument added to the field
			//Adding non-nullable argument is a breaking change
			if nArg.Type.NonNull {
				changes = append(changes, &Change{
					changeType:       DirectiveArgumentAdded,
					criticalityLevel: Breaking,
					message:          fmt.Sprintf("Non-nullable argument '%s:%s' was added to directive '@%s'", nArg.Name, nArg.Type.String(), nd.Name),
					path:             fmt.Sprintf("@%s.%s", od.Name, nArg.Name),
					position:         nArg.Position,
				})
			} else {
				changes = append(changes, &Change{
					changeType:       DirectiveArgumentAdded,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Argument '%s:%s' was added to to directive '@%s'", nArg.Name, nArg.Type.String(), nd.Name),
					path:             fmt.Sprintf("@%s.%s", od.Name, nArg.Name),
					position:         nArg.Position,
				})
			}
		}
	}
	return changes
}

func changeInDirectiveArguments(od *ast.DirectiveDefinition, nd *ast.DirectiveDefinition) []*Change {
	var changes []*Change
	for _, oArg := range od.Arguments {
		nArg := nd.Arguments.ForName(oArg.Name)
		if nArg == nil {
			//argument is removed
			changes = append(changes, &Change{
				changeType:       DirectiveArgumentRemoved,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("Argument '%s' was removed from directive '@%s'", oArg.Name, od.Name),
				path:             fmt.Sprintf("@%s.%s", od.Name, oArg.Name),
				position:         nd.Position,
			})
		} else {
			//check argument type change
			changes = append(changes, checkDirectiveArgumentTypeChanged(od, oArg, nArg)...)

			//check description change
			if oArg.Description != nArg.Description {
				changes = append(changes, &Change{
					changeType:       DirectiveArgumentDescriptionChanged,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Argument '%s' description changed in directive '@%s' ", oArg.Name, od.Name),
					path:             fmt.Sprintf("@%s.%s", od.Name, oArg.Name),
					position:         nArg.Position,
				})
			}
		}
	}
	return changes
}

func checkDirectiveArgumentTypeChanged(od *ast.DirectiveDefinition, oArg *ast.ArgumentDefinition, nArg *ast.ArgumentDefinition) []*Change {
	var changes []*Change
	if oArg.Type.String() != nArg.Type.String() {
		//Changing an input field from non-null to null is considered non-breaking.
		cl := NonBreaking
		if !isSafeChangeForInputValue(oArg.Type, nArg.Type) {
			cl = Breaking
		}
		changes = append(changes, &Change{
			changeType:       DirectiveArgumentTypeChanged,
			criticalityLevel: cl,
			message:          fmt.Sprintf("Argument '%s' type changed from '%s' to '%s' in directive '@%s' ", oArg.Name, oArg.Type.String(), nArg.Type.String(), od.Name),
			path:             fmt.Sprintf("@%s.%s", od.Name, oArg.Name),
			position:         nArg.Position,
		})
	}
	//Changing the default value for an argument may change the runtime behaviour of a field if it was never provided.
	if oArg.DefaultValue.String() != nArg.DefaultValue.String() {
		changes = append(changes, &Change{
			changeType:       DirectiveArgumentDefaultValueChanged,
			criticalityLevel: Dangerous,
			message:          fmt.Sprintf("Argument '%s' default value changed from '%s' to '%s' in directive '@%s' ", oArg.Name, oArg.DefaultValue.String(), nArg.DefaultValue.String(), od.Name),
			path:             fmt.Sprintf("@%s.%s", od.Name, oArg.Name),
			position:         nArg.Position,
		})
	}
	return changes
}

func checkDirectiveRepeatableChanged(od *ast.DirectiveDefinition, nd *ast.DirectiveDefinition) []*Change {
	var changes []*Change
	//isRepeatable removed
	if od.IsRepeatable && !nd.IsRepeatable {
		changes = append(changes, &Change{
			changeType:       DirectiveRepeatableRemoved,
			criticalityLevel: Breaking,
			message:          fmt.Sprintf("Repeatable flag was removed from '@%s' directive", od.Name),
			path:             fmt.Sprintf("@%s", nd.Name),
			position:         nd.Position,
		})
	}
	//isRepeatable added
	if !od.IsRepeatable && nd.IsRepeatable {
		changes = append(changes, &Change{
			changeType:       DirectiveRepeatableAdded,
			criticalityLevel: NonBreaking,
			message:          fmt.Sprintf("Repeatable flag was removed from '@%s' directive", od.Name),
			path:             fmt.Sprintf("@%s", nd.Name),
			position:         nd.Position,
		})
	}
	return changes
}

func checkDirectiveLocationChanged(od *ast.DirectiveDefinition, nd *ast.DirectiveDefinition) []*Change {
	var changes []*Change
	//location changed
	found := false
	for _, ol := range od.Locations {
		found = false
		for _, nl := range nd.Locations {
			if ol == nl {
				found = true
			}
		}
		if !found {
			changes = append(changes, &Change{
				changeType:       DirectiveLocationRemoved,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("Location '%s' was removed from '@%s' directive", ol, od.Name),
				path:             fmt.Sprintf("@%s", nd.Name),
				position:         nd.Position,
			})
		}
	}
	for _, nl := range nd.Locations {
		found = false
		for _, ol := range od.Locations {
			if nl == ol {
				found = true
			}
		}
		if !found {
			changes = append(changes, &Change{
				changeType:       DirectiveLocationAdded,
				criticalityLevel: NonBreaking,
				message:          fmt.Sprintf("Location '%s' was added to '@%s' directive", nl, nd.Name),
				path:             fmt.Sprintf("@%s", nd.Name),
				position:         nd.Position,
			})
		}
	}
	return changes
}

func checkDirectiveAdded(oDirs ast.DirectiveDefinitionList, nDirs ast.DirectiveDefinitionList) []*Change {
	var changes []*Change
	for _, nd := range nDirs {
		od := oDirs.ForName(nd.Name)
		if od == nil {
			changes = append(changes, &Change{
				changeType:       DirectiveAdded,
				criticalityLevel: NonBreaking,
				message:          fmt.Sprintf("Directive '@%s' was added ", nd.Name),
				path:             fmt.Sprintf("@%s", nd.Name),
				position:         nd.Position,
			})
		}
	}
	return changes
}

func changeInEnum(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	for _, ov := range oDef.EnumValues {
		nv := nDef.EnumValues.ForName(ov.Name)
		oDep := ov.Directives.ForName(deprecatedDirective)
		if nv == nil {
			msg := fmt.Sprintf("Enum value '%s' was removed from enum '%s'", ov.Name, oDef.Name)
			if oDep != nil {
				msg = fmt.Sprintf("Enum value '%s'(deprecated) was removed from enum '%s' ", ov.Name, oDef.Name)
			}
			//Removing an enum value will cause existing queries that use this enum value to error.
			changes = append(changes, &Change{
				changeType:       EnumValueRemoved,
				criticalityLevel: Breaking,
				message:          msg,
				path:             fmt.Sprintf("%s.%s", oDef.Name, ov.Name),
				position:         ov.Position,
			})
		} else {
			if ov.Description != nv.Description {
				changes = append(changes, &Change{
					changeType:       EnumValueDescriptionChanged,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Enum value '%s' description changed in  enum '%s' ", ov.Name, oDef.Name),
					path:             fmt.Sprintf("%s.%s", oDef.Name, ov.Name),
					position:         nv.Position,
				})
			}
			changes = append(changes, checkEnumValueDeprecationChanged(oDef, nv, ov)...)
		}
	}
	changes = append(changes, checkEnumValuesAdded(oDef, nDef)...)
	return changes
}

func checkEnumValueDeprecationChanged(oDef *ast.Definition, nv *ast.EnumValueDefinition, ov *ast.EnumValueDefinition) []*Change {
	var changes []*Change
	oDep := ov.Directives.ForName(deprecatedDirective)
	nDep := nv.Directives.ForName(deprecatedDirective)
	if oDep == nil && nDep != nil {
		changes = append(changes, &Change{
			changeType:       EnumValueDeprecationAdded,
			criticalityLevel: Dangerous,
			message:          fmt.Sprintf("Enum value '%s' deprecated in enum '%s' ", ov.Name, oDef.Name),
			path:             fmt.Sprintf("%s.%s", oDef.Name, ov.Name),
			position:         nv.Position,
		})
	}
	if oDep != nil && nDep != nil && oDep.Arguments.ForName("reason") != nDep.Arguments.ForName("reason") {
		changes = append(changes, &Change{
			changeType:       EnumValueDeprecationReasonChanged,
			criticalityLevel: NonBreaking,
			message:          fmt.Sprintf("Enum value '%s' deprecation reason changed in enum '%s' ", ov.Name, oDef.Name),
			path:             fmt.Sprintf("%s.%s", oDef.Name, ov.Name),
			position:         nv.Position,
		})
	}
	return changes
}

func checkEnumValuesAdded(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	for _, nv := range nDef.EnumValues {
		ov := oDef.EnumValues.ForName(nv.Name)
		if ov == nil {
			//Adding an enum value may break existing clients that were not programming defensively against an added case when querying an enum.
			changes = append(changes, &Change{
				changeType:       EnumValueAdded,
				criticalityLevel: Dangerous,
				message:          fmt.Sprintf("Enum value '%s' was added to enum '%s'", nv.Name, nDef.Name),
				path:             fmt.Sprintf("%s.%s", oDef.Name, nv.Name),
				position:         nv.Position,
			})
		}
	}
	return changes
}

func changeInObject(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	//check implementing interfaces
	changes = append(changes, checkTypeInterfaceRemoved(oDef, nDef)...)
	changes = append(changes, checkTypeInterfacesAdded(oDef, nDef)...)
	changes = append(changes, changeInTypeFieldDirectives(oDef.Directives, nDef.Directives, nDef.Name, nDef.Position)...)
	changes = append(changes, changeInFields(oDef, nDef)...)
	return changes
}

func checkTypeInterfaceRemoved(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	found := false
	for _, oInt := range oDef.Interfaces {
		found = false
		for _, nInt := range nDef.Interfaces {
			if oInt == nInt {
				found = true
			}
		}
		if !found {
			//Removing an interface from an object type can cause existing queries that use this in a fragment spread to error.
			changes = append(changes, &Change{
				changeType:       ObjectTypeInterfaceRemoved,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("'%s' object type no longer implements '%s' interface", oDef.Name, oInt),
				path:             oDef.Name,
				position:         nDef.Position,
			})
		}
	}
	return changes
}

func checkTypeInterfacesAdded(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	found := false
	for _, nInt := range nDef.Interfaces {
		found = false
		for _, oInt := range oDef.Interfaces {
			if oInt == nInt {
				found = true
			}
		}
		if !found {
			//Adding an interface to an object type may break existing clients that were not programming defensively against a new possible type.
			changes = append(changes, &Change{
				changeType:       ObjectTypeInterfaceAdded,
				criticalityLevel: Dangerous,
				message:          fmt.Sprintf("'%s' object type implements '%s' interface", nDef.Name, nInt),
				path:             oDef.Name,
				position:         nDef.Position,
			})
		}
	}
	return changes
}

func changeInUnion(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	//Check if union types added/removed
	changes = append(changes, checkUnionMemberRemoved(oDef, nDef)...)
	changes = append(changes, checkUnionMemberAdded(oDef, nDef)...)
	return changes
}

func checkUnionMemberRemoved(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	found := false
	for _, ot := range oDef.Types {
		found = false
		for _, nt := range nDef.Types {
			if ot == nt {
				found = true
			}
		}
		if !found {
			//Removing a union member from a union can cause existing queries that use this union member in a fragment spread to error.
			changes = append(changes, &Change{
				changeType:       UnionMemberRemoved,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("Member '%s' was removed from Union type '%s'", ot, oDef.Name),
				path:             oDef.Name,
				position:         nDef.Position,
			})
		}
	}
	return changes
}

func checkUnionMemberAdded(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	found := false
	for _, nt := range nDef.Types {
		found = false
		for _, ot := range oDef.Types {
			if ot == nt {
				found = true
			}
		}
		if !found {
			//Adding a possible type to Unions may break existing clients that were not programming defensively against a new possible type.
			changes = append(changes, &Change{
				changeType:       UnionMemberAdded,
				criticalityLevel: Dangerous,
				message:          fmt.Sprintf("Member '%s' was added to Union type '%s'", nt, nDef.Name),
				path:             oDef.Name,
				position:         nDef.Position,
			})
		}
	}
	return changes
}

func changeInFields(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	for _, of := range oDef.Fields {
		nf := nDef.Fields.ForName(of.Name)
		oDep := of.Directives.ForName(deprecatedDirective)
		if nf == nil {
			//Removing a field is a breaking change. It is preferable to deprecate the field before removing it.
			msg := fmt.Sprintf("Field '%s.%s' was removed from %s", oDef.Name, of.Name, oDef.Kind)
			if oDep != nil {
				//Removing a deprecated field is a breaking change.
				//Before removing it, you may want to look at the field's usage to see the impact of removing the field.
				msg = fmt.Sprintf("Field '%s.%s'(deprecated) was removed from %s", oDef.Name, of.Name, oDef.Kind)
			}
			changes = append(changes, &Change{
				changeType:       FieldRemoved,
				criticalityLevel: Breaking,
				message:          msg,
				path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
				position:         nDef.Position,
			})
		} else {
			//Check field type
			changes = append(changes, checkFieldTypeChanged(oDef, of, nf)...)
			//Check description change
			if of.Description != nf.Description {
				changes = append(changes, &Change{
					changeType:       FieldDescriptionChanged,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Field '%s.%s' description changed in %s", oDef.Name, of.Name, oDef.Kind),
					path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
					position:         nf.Position,
				})
			}
			//Check deprecation changes
			changes = append(changes, checkFieldDeprecationChanged(oDef, nf, of)...)
			//check argument changes
			changes = append(changes, changeInArgument(of, nf, oDef.Name)...)
			changes = append(changes, changeInTypeFieldDirectives(of.Directives, nf.Directives, fmt.Sprintf("%s.%s", nDef.Name, nf.Name), nf.Position)...)
		}
	}
	changes = append(changes, checkFieldsAdded(oDef, nDef)...)
	return changes
}

func checkFieldTypeChanged(oDef *ast.Definition, of *ast.FieldDefinition, nf *ast.FieldDefinition) []*Change {
	var changes []*Change
	if of.Type.String() != nf.Type.String() {
		cl := NonBreaking
		if !isSafeChangeForFieldType(of.Type, nf.Type) {
			cl = Breaking
		}
		changes = append(changes, &Change{
			changeType:       FieldTypeChanged,
			criticalityLevel: cl,
			message:          fmt.Sprintf("Field '%s.%s' type changed from '%s' to '%s' in %s ", oDef.Name, of.Name, of.Type.String(), nf.Type.String(), oDef.Kind),
			path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
			position:         nf.Position,
		})
	}
	return changes
}

func checkFieldDeprecationChanged(oDef *ast.Definition, nf *ast.FieldDefinition, of *ast.FieldDefinition) []*Change {
	var changes []*Change
	oDep := of.Directives.ForName(deprecatedDirective)
	nDep := nf.Directives.ForName(deprecatedDirective)
	if oDep == nil && nDep != nil {
		changes = append(changes, &Change{
			changeType:       FieldDeprecationAdded,
			criticalityLevel: Dangerous,
			message:          fmt.Sprintf("Field '%s.%s' deprecated in %s ", oDef.Name, of.Name, oDef.Kind),
			path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
			position:         nf.Position,
		})
	}
	if oDep != nil && nDep == nil {
		changes = append(changes, &Change{
			changeType:       FieldDeprecationRemoved,
			criticalityLevel: Dangerous,
			message:          fmt.Sprintf("Field '%s.%s' deprecation removed in %s ", oDef.Name, of.Name, oDef.Kind),
			path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
			position:         nf.Position,
		})
	}
	if oDep != nil && nDep != nil && oDep.Arguments.ForName("reason") != nDep.Arguments.ForName("reason") {
		changes = append(changes, &Change{
			changeType:       FieldDeprecationReasonChanged,
			criticalityLevel: NonBreaking,
			message:          fmt.Sprintf("Field '%s.%s' deprecation reason changed in %s ", oDef.Name, of.Name, oDef.Kind),
			path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
			position:         nf.Position,
		})
	}
	return changes
}

func checkFieldsAdded(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	for _, nf := range nDef.Fields {
		if oDef.Fields.ForName(nf.Name) == nil {
			//Field added to the type
			changes = append(changes, &Change{
				changeType:       FieldAdded,
				criticalityLevel: NonBreaking,
				message:          fmt.Sprintf("Field '%s.%s' was added to %s", nDef.Name, nf.Name, nDef.Kind),
				path:             fmt.Sprintf("%s.%s", oDef.Name, nf.Name),
				position:         nf.Position,
			})
		}
	}
	return changes
}

func changeInArgument(oDef *ast.FieldDefinition, nDef *ast.FieldDefinition, typeName string) []*Change {
	var changes []*Change
	for _, oArg := range oDef.Arguments {
		nArg := nDef.Arguments.ForName(oArg.Name)
		if nArg == nil {
			//Removing a field argument is a breaking change because it will cause existing queries that use this argument to error.
			changes = append(changes, &Change{
				changeType:       FieldArgumentRemoved,
				criticalityLevel: Breaking,
				message:          fmt.Sprintf("Argument '%s:%s' was removed from field '%s.%s'", oArg.Name, oArg.Type.String(), typeName, nDef.Name),
				path:             fmt.Sprintf("%s.%s.%s", typeName, oDef.Name, oArg.Name),
				position:         nDef.Position,
			})
		} else {
			//check argument type change
			changes = append(changes, checkFieldArgumentTypeChanged(oArg, nArg, typeName, oDef.Name)...)
			//Changing the default value for an argument may change the runtime behaviour of a field if it was never provided.
			if oArg.DefaultValue.String() != nArg.DefaultValue.String() {
				changes = append(changes, &Change{
					changeType:       FieldArgumentDefaultChanged,
					criticalityLevel: Dangerous,
					message:          fmt.Sprintf("Argument '%s' default value changed from '%s' to '%s' in '%s.%s' ", oArg.Name, oArg.DefaultValue.String(), nArg.DefaultValue.String(), typeName, oDef.Name),
					path:             fmt.Sprintf("%s.%s.%s", typeName, oDef.Name, oArg.Name),
					position:         nArg.Position,
				})
			}

			//check description change
			if oArg.Description != nArg.Description {
				changes = append(changes, &Change{
					changeType:       FieldArgumentDescriptionChanged,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Argument '%s' description changed in '%s.%s' ", oArg.Name, typeName, oDef.Name),
					path:             fmt.Sprintf("%s.%s.%s", typeName, oDef.Name, oArg.Name),
					position:         nArg.Position,
				})
			}
		}
	}
	changes = append(changes, checkFieldArgumentAdded(oDef, nDef, typeName)...)
	return changes
}

func checkFieldArgumentTypeChanged(oArg *ast.ArgumentDefinition, nArg *ast.ArgumentDefinition, typeName string, fieldName string) []*Change {
	var changes []*Change
	if oArg.Type.String() != nArg.Type.String() {
		//Changing an input field from non-null to null is considered non-breaking.
		cl := NonBreaking
		if !isSafeChangeForInputValue(oArg.Type, nArg.Type) {
			//Changing the type of a field's argument can cause existing queries that use this argument to error.
			cl = Breaking
		}
		changes = append(changes, &Change{
			changeType:       FieldArgumentTypeChanged,
			criticalityLevel: cl,
			message:          fmt.Sprintf("Argument '%s' type changed from '%s' to '%s' in '%s.%s' ", oArg.Name, oArg.Type.String(), nArg.Type.String(), typeName, fieldName),
			path:             fmt.Sprintf("%s.%s.%s", typeName, fieldName, oArg.Name),
			position:         nArg.Position,
		})
	}
	return changes
}

func checkFieldArgumentAdded(oDef *ast.FieldDefinition, nDef *ast.FieldDefinition, typeName string) []*Change {
	var changes []*Change
	for _, nArg := range nDef.Arguments {
		oArg := oDef.Arguments.ForName(nArg.Name)
		if oArg == nil {
			//Adding a required argument to an existing field is a breaking change because it will cause existing uses of this field to error.
			if nArg.Type.NonNull {
				changes = append(changes, &Change{
					changeType:       FieldArgumentAdded,
					criticalityLevel: Breaking,
					message:          fmt.Sprintf("Required argument '%s:%s' was added to field '%s.%s'", nArg.Name, nArg.Type.String(), typeName, nDef.Name),
					path:             fmt.Sprintf("%s.%s.%s", typeName, oDef.Name, nArg.Name),
					position:         nArg.Position,
				})
			} else {
				//Adding a new argument to an existing field may involve a change in resolve function logic that potentially may cause some side effects.
				changes = append(changes, &Change{
					changeType:       FieldArgumentAdded,
					criticalityLevel: Dangerous,
					message:          fmt.Sprintf("Argument '%s:%s' was added to field '%s.%s'", nArg.Name, nArg.Type.String(), typeName, nDef.Name),
					path:             fmt.Sprintf("%s.%s.%s", typeName, oDef.Name, nArg.Name),
					position:         nArg.Position,
				})
			}
		}
	}
	return changes
}

func changeInInputFields(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	for _, of := range oDef.Fields {
		nf := nDef.Fields.ForName(of.Name)
		oDep := of.Directives.ForName(deprecatedDirective)
		if nf == nil {
			//Removing an input field will cause existing queries that use this input field to error.
			msg := fmt.Sprintf("Input field '%s.%s' was removed from input object type", oDef.Name, of.Name)
			if oDep != nil {
				msg = fmt.Sprintf("input field '%s.%s'(deprecated) was removed from input object type", oDef.Name, of.Name)
			}
			changes = append(changes, &Change{
				changeType:       InputFieldRemoved,
				criticalityLevel: Breaking,
				message:          msg,
				path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
				position:         nDef.Position,
			})
		} else {
			//Check input field type
			changes = append(changes, checkInputFieldTypeValueChanged(oDef, of, nf)...)
			//Check description change
			if of.Description != nf.Description {
				changes = append(changes, &Change{
					changeType:       InputFieldDescriptionChanged,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Input field '%s.%s' description changed in input object type", oDef.Name, of.Name),
					path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
					position:         nf.Position,
				})
			}
			//check change in field directives
			changes = append(changes, changeInTypeFieldDirectives(of.Directives, nf.Directives, fmt.Sprintf("%s.%s", nDef.Name, nf.Name), nf.Position)...)
		}
	}
	changes = append(changes, checkInputFieldsAdded(oDef, nDef)...)
	return changes
}

func checkInputFieldTypeValueChanged(oDef *ast.Definition, of *ast.FieldDefinition, nf *ast.FieldDefinition) []*Change {
	var changes []*Change
	if of.Type.String() != nf.Type.String() {
		//Changing an input field from non-null to null is considered non-breaking.
		cl := NonBreaking
		if !isSafeChangeForInputValue(of.Type, nf.Type) {
			//Changing the type of an input field can cause existing queries that use this field to error.
			cl = Breaking
		}
		changes = append(changes, &Change{
			changeType:       InputFieldTypeChanged,
			criticalityLevel: cl,
			message:          fmt.Sprintf("Input field '%s.%s' type changed from '%s' to '%s' in input object type", oDef.Name, of.Name, of.Type.String(), nf.Type.String()),
			path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
			position:         nf.Position,
		})
	}
	//Changing the default value for an argument may change the runtime behaviour of a field if it was never provided.
	if of.DefaultValue.String() != nf.DefaultValue.String() {
		changes = append(changes, &Change{
			changeType:       InputFieldDefaultValueChanged,
			criticalityLevel: Dangerous,
			message:          fmt.Sprintf("Input field '%s.%s' default value changed from '%s' to '%s' in input object type", oDef.Name, of.Name, of.DefaultValue.String(), nf.DefaultValue.String()),
			path:             fmt.Sprintf("%s.%s", oDef.Name, of.Name),
			position:         nf.Position,
		})
	}
	return changes
}

func checkInputFieldsAdded(oDef *ast.Definition, nDef *ast.Definition) []*Change {
	var changes []*Change
	for _, nf := range nDef.Fields {
		if oDef.Fields.ForName(nf.Name) == nil {
			//input field added to the type
			if nf.Type.NonNull {
				//Adding a required input field to an existing input object type is a breaking change because it will cause existing uses of this input object type to error.
				changes = append(changes, &Change{
					changeType:       InputFieldAdded,
					criticalityLevel: Breaking,
					message:          fmt.Sprintf("Required field '%s' was added to input object type '%s'", nf.Name, nDef.Name),
					path:             fmt.Sprintf("%s.%s", oDef.Name, nf.Name),
					position:         nf.Position,
				})
			} else {
				changes = append(changes, &Change{
					changeType:       InputFieldAdded,
					criticalityLevel: Dangerous,
					message:          fmt.Sprintf("Field '%s' was added to input object type '%s'", nf.Name, nDef.Name),
					path:             fmt.Sprintf("%s.%s", oDef.Name, nf.Name),
					position:         nf.Position,
				})
			}
		}
	}
	return changes
}

func changeInTypeFieldDirectives(oDirs ast.DirectiveList, nDirs ast.DirectiveList, typeName string, pos *ast.Position) []*Change {
	var changes []*Change
	for _, od := range oDirs {
		oDirList := oDirs.ForNames(od.Name) //if the directive is repetitive
		if od.Name != deprecatedDirective {
			nDirList := nDirs.ForNames(od.Name)
			switch {
			case len(nDirList) == 0:
				changes = append(changes, &Change{
					changeType:       DirectiveRemoved,
					criticalityLevel: Dangerous,
					message:          fmt.Sprintf("Directive '@%s' was removed from '%s'", od.Name, typeName),
					path:             typeName,
					position:         pos,
				})
			case len(nDirList) == 1 && len(oDirList) == 1:
				//means there is only one directive, check for the argument changes
				nd := nDirList[0]
				changes = append(changes, checkFieldDirectiveArgumentChanged(od, nd, typeName, pos)...)
				changes = append(changes, checkFieldDirectiveArgumentAdded(od, nd, typeName)...)
			default:
				//check if at least one directive from the list matches arguments
				haveSameArgVals := false
				for _, nd := range nDirList {
					if len(od.Arguments) == len(nd.Arguments) {
						for _, oArg := range od.Arguments {
							nArg := nd.Arguments.ForName(oArg.Name)
							if nArg != nil && oArg.Value.String() == nArg.Value.String() {
								haveSameArgVals = true // args checked till now are same
								continue
							}
							// We found first arg not matching so not need to check more args
							haveSameArgVals = false
							break
						}
						if haveSameArgVals {
							// means the directive have all the arguments matching so no need to check other directives in the list
							break
						}
					}
				}
				if !haveSameArgVals {
					changes = append(changes, &Change{
						changeType:       DirectiveChanged,
						criticalityLevel: Dangerous,
						message:          fmt.Sprintf("Directive '@%s' was changed on '%s'", od.Name, typeName),
						path:             typeName,
						position:         pos,
					})
				}
			}
		}
	}
	changes = append(changes, checkFieldDirectiveAdded(oDirs, nDirs, typeName)...)
	return changes
}

func checkFieldDirectiveArgumentAdded(od *ast.Directive, nd *ast.Directive, typeName string) []*Change {
	var changes []*Change
	for _, nArg := range nd.Arguments {
		oArg := od.Arguments.ForName(nArg.Name)
		if oArg == nil {
			//argument added to the field
			changes = append(changes, &Change{
				changeType:       DirectiveArgumentAdded,
				criticalityLevel: NonBreaking,
				message:          fmt.Sprintf("Directive '@%s' argument '%s' was added to in '%s'", nd.Name, nArg.Name, typeName),
				path:             fmt.Sprintf("%s.@%s", typeName, od.Name),
				position:         nArg.Position,
			})
		}
	}
	return changes
}

func checkFieldDirectiveArgumentChanged(od *ast.Directive, nd *ast.Directive, typeName string, pos *ast.Position) []*Change {
	var changes []*Change
	for _, oArg := range od.Arguments {
		nArg := nd.Arguments.ForName(oArg.Name)
		if nArg == nil {
			//argument is removed
			changes = append(changes, &Change{
				changeType:       DirectiveArgumentRemoved,
				criticalityLevel: Dangerous,
				message:          fmt.Sprintf("Directive '@%s' argument '%s' was removed in '%s'", od.Name, oArg.Name, typeName),
				path:             fmt.Sprintf("@%s.%s", od.Name, oArg.Name),
				position:         pos,
			})
		} else if oArg.Value.String() != nArg.Value.String() {
			changes = append(changes, &Change{
				changeType:       DirectiveArgumentValueChanged,
				criticalityLevel: Dangerous,
				message:          fmt.Sprintf("Directive '@%s' argument '%s' value changed from '%s' to '%s' in '%s' ", od.Name, oArg.Name, oArg.Value.String(), nArg.Value.String(), typeName),
				path:             fmt.Sprintf("@%s.%s", od.Name, oArg.Name),
				position:         nArg.Position,
			})
		}
	}
	return changes
}

func checkFieldDirectiveAdded(oDirs ast.DirectiveList, nDirs ast.DirectiveList, typeName string) []*Change {
	var changes []*Change
	for _, nd := range nDirs {
		if nd.Name != deprecatedDirective {
			od := oDirs.ForName(nd.Name)
			if od == nil {
				changes = append(changes, &Change{
					changeType:       DirectiveAdded,
					criticalityLevel: NonBreaking,
					message:          fmt.Sprintf("Directive '@%s' was added in '%s'", nd.Name, typeName),
					path:             typeName,
					position:         nd.Position,
				})
			}
		}
	}
	return changes
}

func isSafeChangeForFieldType(otyp *ast.Type, ntyp *ast.Type) bool {
	if !isWrappingType(otyp) && !isWrappingType(ntyp) {
		//if they're both named types, see if their names are equivalent
		return otyp.String() == ntyp.String()
	}
	if !ntyp.NonNull && otyp.NonNull {
		return false
	}
	if ntyp.NonNull {
		if isListType(ntyp) {
			//if they're both lists, make sure underlying types are compatible
			return isListType(otyp) && isSafeChangeForFieldType(otyp.Elem, ntyp.Elem)
		}
		//moving from nullable to non-nullable is safe change
		return otyp.NamedType == ntyp.NamedType
	}
	if isListType(otyp) {
		//if they're both lists, make sure underlying types are compatible
		return isListType(ntyp) && isSafeChangeForFieldType(otyp.Elem, ntyp.Elem)
	}
	return false
}

func isSafeChangeForInputValue(otyp *ast.Type, ntyp *ast.Type) bool {
	if !isWrappingType(otyp) && !isWrappingType(ntyp) {
		// if they're both named types, see if their names are equivalent
		return otyp.String() == ntyp.String()
	}
	if !otyp.NonNull && ntyp.NonNull {
		return false
	}
	if otyp.NonNull {
		if isListType(otyp) {
			//if they're both lists, make sure underlying types are compatible
			return isListType(ntyp) && isSafeChangeForInputValue(otyp.Elem, ntyp.Elem)
		}
		//moving from non-nullable to nullable is safe change
		return otyp.NamedType == ntyp.NamedType
	}
	// if they're both lists, make sure underlying types are compatible
	if isListType(otyp) && isListType(ntyp) {
		return isSafeChangeForInputValue(otyp.Elem, ntyp.Elem)
	}
	return false
}

//IsListType checks if a type is a list
func isListType(typ *ast.Type) bool {
	return typ != nil && typ.Elem != nil && typ.NamedType == ""
}

//IsNonNullType checks if a type can be null or not
func isNonNullType(typ *ast.Type) bool {
	return typ != nil && typ.NonNull
}

func isWrappingType(typ *ast.Type) bool {
	return isListType(typ) || isNonNullType(typ)
}

func getOperationForName(ops ast.OperationTypeDefinitionList, name ast.Operation) *ast.OperationTypeDefinition {
	for _, op := range ops {
		if op.Operation == name {
			return op
		}
	}
	return nil
}

// GroupChanges group all changes on their criticality level
func GroupChanges(changes []*Change) map[Criticality][]*Change {
	groupChanges := map[Criticality][]*Change{}
	for _, c := range changes {
		if _, ok := groupChanges[c.criticalityLevel]; !ok {
			groupChanges[c.criticalityLevel] = []*Change{}
		}
		groupChanges[c.criticalityLevel] = append(groupChanges[c.criticalityLevel], c)
	}
	return groupChanges
}

// ReportBreakingChanges print only breaking changes in output
func ReportBreakingChanges(changes []*Change, withFilepath bool) int {
	if len(changes) == 0 {
		return 0
	}
	sort.Slice(changes, less(changes))
	for _, c := range changes {
		if pos := getPosition(c); withFilepath && len(pos) > 0 {
			fmt.Printf("%s  %s %s\n", "❌", pos, c.message)
			continue
		}
		fmt.Printf("%s  %s\n", "❌", c.message)
	}
	return len(changes)
}

// ReportDangerousChanges print only breaking changes in output
func ReportDangerousChanges(changes []*Change, withFilepath bool) int {
	if len(changes) == 0 {
		return 0
	}
	sort.Slice(changes, less(changes))
	for _, c := range changes {
		if pos := getPosition(c); withFilepath && len(pos) > 0 {
			fmt.Printf("%s  %s %s\n", "✋️", pos, c.message)
			continue
		}
		fmt.Printf("%s  %s\n", "✋️", c.message)
	}
	return len(changes)
}

// ReportNonBreakingChanges print only breaking changes in output
func ReportNonBreakingChanges(changes []*Change, withFilepath bool) int {
	if len(changes) == 0 {
		return 0
	}
	sort.Slice(changes, less(changes))
	for _, c := range changes {
		if pos := getPosition(c); withFilepath && len(pos) > 0 {
			fmt.Printf("%s  %s %s\n", "✅", pos, c.message)
			continue
		}
		fmt.Printf("%s  %s\n", "✅", c.message)
	}
	return len(changes)
}

func less(changes []*Change) func(i int, j int) bool {
	return func(i, j int) bool {
		if changes[i].position.Src.Name != changes[j].position.Src.Name {
			return changes[i].position.Src.Name < changes[j].position.Src.Name
		}
		return changes[i].position.Line < changes[j].position.Line
	}
}

func getPosition(c *Change) string {
	position := ""
	if c.position != nil {
		fileName := ""
		if c.position.Src.Name != "" {
			fileName = c.position.Src.Name
		}
		position = fmt.Sprintf("%s:%d", fileName, c.position.Line)
	}
	return position
}
