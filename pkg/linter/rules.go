package linter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

var camelCaseRegex, _ = regexp.Compile("^[a-z][a-zA-Z0-9]*$")

// LintRuleFunc is a short-form for the function signature for every lint Rule function should have
type LintRuleFunc = func(schema *ast.SchemaDocument) LintErrorsWithMetadata

// LintRule is the name of the lint Rule internally
type LintRule string

// LintRuleMetadata holds information for a given lint Rule
type LintRuleMetadata struct {
	Name         LintRule
	description  string
	RuleFunction LintRuleFunc
}

// AvailableRulesWithDescription returns the comma separated list of rules with description
func AvailableRulesWithDescription() string {
	availableRulesWithDescription := make([]string, 0)
	for _, rule := range AllTheRules {
		ruleWithDescription := fmt.Sprintf("	%s => %s", rule.Name, rule.description)
		availableRulesWithDescription = append(availableRulesWithDescription, ruleWithDescription)
	}
	return strings.Join(availableRulesWithDescription, "\n")
}

const (
	typeDesc      = "type-desc"
	argsDesc      = "args-desc"
	fieldDesc     = "field-desc"
	enumCaps      = "enum-caps"
	enumDesc      = "enum-desc"
	fieldCamel    = "field-camel"
	typeCaps      = "type-caps"
	relayConnType = "relay-conn-type"
	relayConnArgs = "relay-conn-args"
)

// AllTheRules is a list of all the lint rules available
var AllTheRules = []LintRuleMetadata{
	{
		typeDesc,
		"type-desc checks whether all the types defined have description",
		TypesHaveDescription,
	},
	{
		argsDesc,
		"args-desc checks whether arguments have description",
		ArgumentsHaveDescription,
	},
	{
		fieldDesc,
		"field-desc checks whether fields have description",
		FieldsHaveDescription,
	},
	{
		enumCaps,
		"enum-caps checks whether Enum values are all UPPER_CASE",
		EnumValuesAreAllCaps,
	},
	{
		enumDesc,
		"enum-desc checks whether Enum values have description",
		EnumValuesHaveDescriptions,
	},
	{
		fieldCamel,
		"field-camel checks whether fields defined are all camelCase",
		FieldsAreCamelCased,
	},
	{
		typeCaps,
		"type-caps checks whether types defined are Capitalized",
		TypesAreCapitalized,
	},
	{
		relayConnType,
		"relay-conn-type checks if Connection Types follow the Relay Cursor Connections Specification",
		RelayConnectionTypesSpec,
	},
	{
		relayConnArgs,
		"relay-conn-args checks if Connection Args follow of the Relay Cursor Connections Specification",
		RelayConnectionArgumentsSpec,
	},
}

// TypesHaveDescription checks whether all the types defined have description
func TypesHaveDescription(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, definition := range schema.Definitions {
		if len(definition.Description) == 0 {
			lintError := LintErrorWithMetadata{
				Rule:   typeDesc,
				Line:   definition.Position.Line,
				Column: definition.Position.Column,
				Err:    fmt.Errorf("type %s does not have description", definition.Name),
			}
			errors = append(errors, lintError)
		}
	}
	// extended types should not have descriptions since that can collide with type being extended
	return errors
}

// ArgumentsHaveDescription checks whether arguments have description
func ArgumentsHaveDescription(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, definition := range schema.Definitions {
		if definition.IsCompositeType() {
			for _, field := range definition.Fields {
				for _, argument := range field.Arguments {
					if len(argument.Description) == 0 {
						lintError := LintErrorWithMetadata{
							Rule:   argsDesc,
							Line:   argument.Position.Line,
							Column: argument.Position.Column,
							Err:    fmt.Errorf("argument %s.%s.%s does not have description", definition.Name, field.Name, argument.Name),
						}
						errors = append(errors, lintError)
					}
				}
			}
		}
	}
	// extended types are not included in schema.definitions but schema.extensions
	for _, definition := range schema.Extensions {
		if definition.IsCompositeType() {
			for _, field := range definition.Fields {
				for _, argument := range field.Arguments {
					if len(argument.Description) == 0 {
						lintError := LintErrorWithMetadata{
							Rule:   argsDesc,
							Line:   argument.Position.Line,
							Column: argument.Position.Column,
							Err:    fmt.Errorf("argument %s.%s.%s does not have description", definition.Name, field.Name, argument.Name),
						}
						errors = append(errors, lintError)
					}
				}
			}
		}
	}
	return errors
}

// FieldsHaveDescription checks whether fields have description
func FieldsHaveDescription(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, definition := range schema.Definitions {
		for _, fieldDefinition := range definition.Fields {
			if len(fieldDefinition.Description) == 0 {
				lintError := LintErrorWithMetadata{
					Rule:   fieldDesc,
					Line:   fieldDefinition.Type.Position.Line,
					Column: fieldDefinition.Type.Position.Column,
					Err:    fmt.Errorf("field %s.%s does not have description", definition.Name, fieldDefinition.Name),
				}
				errors = append(errors, lintError)
			}
		}
	}
	// extended types are not included in schema.definitions but schema.extensions
	for _, definition := range schema.Extensions {
		for _, fieldDefinition := range definition.Fields {
			if len(fieldDefinition.Description) == 0 {
				lintError := LintErrorWithMetadata{
					Rule:   fieldDesc,
					Line:   fieldDefinition.Type.Position.Line,
					Column: fieldDefinition.Type.Position.Column,
					Err:    fmt.Errorf("field %s.%s does not have description", definition.Name, fieldDefinition.Name),
				}
				errors = append(errors, lintError)
			}
		}
	}
	// ToDo: we should not allow comment on fields with @external directive as well. This is inline with gqlparser not allowing descriptions for extended types.
	return errors
}

// EnumValuesAreAllCaps checks whether Enum values are all UPPER_CASE
func EnumValuesAreAllCaps(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, definition := range schema.Definitions {
		if definition.Kind == ast.Enum {
			for _, enumValue := range definition.EnumValues {
				if strings.ToUpper(enumValue.Name) != enumValue.Name {
					lintError := LintErrorWithMetadata{
						Rule:   enumCaps,
						Line:   enumValue.Position.Line,
						Column: enumValue.Position.Column,
						Err:    fmt.Errorf("enum value %s.%s is not uppercase", definition.Name, enumValue.Name),
					}
					errors = append(errors, lintError)
				}
			}
		}
	}
	// extended types are not included in schema.definitions but schema.extensions
	for _, definition := range schema.Extensions {
		if definition.Kind == ast.Enum {
			for _, enumValue := range definition.EnumValues {
				if strings.ToUpper(enumValue.Name) != enumValue.Name {
					lintError := LintErrorWithMetadata{
						Rule:   enumCaps,
						Line:   enumValue.Position.Line,
						Column: enumValue.Position.Column,
						Err:    fmt.Errorf("extended enum value %s.%s is not uppercase", definition.Name, enumValue.Name),
					}
					errors = append(errors, lintError)
				}
			}
		}
	}
	return errors
}

// EnumValuesHaveDescriptions checks whether Enum values have description
func EnumValuesHaveDescriptions(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, definition := range schema.Definitions {
		typeDefinition := schema.Definitions.ForName(definition.Name)
		if typeDefinition.Kind == ast.Enum {
			for _, enumValue := range typeDefinition.EnumValues {
				if len(enumValue.Description) == 0 {
					lintError := LintErrorWithMetadata{
						Rule:   enumDesc,
						Line:   enumValue.Position.Line,
						Column: enumValue.Position.Column,
						Err:    fmt.Errorf("enum value %s.%s does not have description", typeDefinition.Name, enumValue.Name),
					}
					errors = append(errors, lintError)
				}
			}
		}
	}
	// extended types are not included in schema.definitions but schema.extensions
	for _, definition := range schema.Extensions {
		if definition.Kind == ast.Enum {
			for _, enumValue := range definition.EnumValues {
				if len(enumValue.Description) == 0 {
					lintError := LintErrorWithMetadata{
						Rule:   enumDesc,
						Line:   enumValue.Position.Line,
						Column: enumValue.Position.Column,
						Err:    fmt.Errorf("extended enum value %s.%s does not have description", definition.Name, enumValue.Name),
					}
					errors = append(errors, lintError)
				}
			}
		}
	}
	return errors
}

// FieldsAreCamelCased checks whether fields defined are all camelCase
func FieldsAreCamelCased(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, definition := range schema.Definitions {
		for _, fieldDefinition := range definition.Fields {
			if !camelCaseRegex.MatchString(fieldDefinition.Name) {
				lintError := LintErrorWithMetadata{
					Rule:   fieldCamel,
					Line:   fieldDefinition.Type.Position.Line,
					Column: fieldDefinition.Type.Position.Column,
					Err:    fmt.Errorf("field %s.%s is not camelcased", definition.Name, fieldDefinition.Name),
				}
				errors = append(errors, lintError)
			}
		}
	}
	// extended types are not included in schema.definitions but schema.extensions
	for _, definition := range schema.Extensions {
		for _, fieldDefinition := range definition.Fields {
			if !camelCaseRegex.MatchString(fieldDefinition.Name) {
				lintError := LintErrorWithMetadata{
					Rule:   fieldCamel,
					Line:   fieldDefinition.Type.Position.Line,
					Column: fieldDefinition.Type.Position.Column,
					Err:    fmt.Errorf("field %s.%s is not camelcased", definition.Name, fieldDefinition.Name),
				}
				errors = append(errors, lintError)
			}
		}
	}
	return errors
}

// TypesAreCapitalized checks whether types defined are Capitalized
func TypesAreCapitalized(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)
	for _, typeDefinition := range schema.Definitions {
		if typeDefinition.Name[0] > 97 && typeDefinition.Name[0] <= 122 {
			lintError := LintErrorWithMetadata{
				Rule:   typeCaps,
				Line:   typeDefinition.Position.Line,
				Column: typeDefinition.Position.Column,
				Err:    fmt.Errorf("type %s is not capitalized", typeDefinition.Name),
			}
			errors = append(errors, lintError)
		}
	}
	// extended types are not included in schema.definitions but schema.extensions
	for _, typeDefinition := range schema.Extensions {
		if typeDefinition.Name[0] > 97 && typeDefinition.Name[0] <= 122 {
			lintError := LintErrorWithMetadata{
				Rule:   typeCaps,
				Line:   typeDefinition.Position.Line,
				Column: typeDefinition.Position.Column,
				Err:    fmt.Errorf("extended type %s is not capitalized", typeDefinition.Name),
			}
			errors = append(errors, lintError)
		}
	}
	return errors
}

// RelayConnectionTypesSpec will validate the schema adheres to section 2 (Connection Types) of the Relay Cursor Connections Specification.
// See https://relay.dev/graphql/connections.htm#sec-Connection-Types and https://relay.dev/graphql/connections.htm
func RelayConnectionTypesSpec(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)

	for _, typeDefinition := range schema.Definitions {
		if strings.HasSuffix(typeDefinition.Name, "Connection") {
			if typeDefinition.Kind != ast.Object {
				lintError := LintErrorWithMetadata{
					Rule:   relayConnType,
					Line:   typeDefinition.Position.Line,
					Column: typeDefinition.Position.Column,
					Err:    fmt.Errorf("%d:%d type %s cannot end with Connection as that is reserved for entities", typeDefinition.Position.Line, typeDefinition.Position.Column, typeDefinition.Name),
				}
				errors = append(errors, lintError)
				continue
			}

			var foundEdgesField, foundPageInfoField bool
			for _, fieldDefinition := range typeDefinition.Fields {
				if fieldDefinition.Name == "edges" {
					foundEdgesField = true
					if !isFieldListType(fieldDefinition) {
						lintError := LintErrorWithMetadata{
							Rule:   relayConnType,
							Line:   fieldDefinition.Type.Position.Line,
							Column: fieldDefinition.Type.Position.Column,
							Err:    fmt.Errorf("%d:%d edges field from Connection type %s needs to return a list type", fieldDefinition.Type.Position.Line, fieldDefinition.Type.Position.Column, typeDefinition.Name),
						}
						errors = append(errors, lintError)
					}

				} else if fieldDefinition.Name == "pageInfo" {
					foundPageInfoField = true

					// this is to account for extra spaces such as PageInfo !
					if fieldDefinition.Type.Name() != "PageInfo" || !fieldDefinition.Type.NonNull || isFieldListType(fieldDefinition) {
						lintError := LintErrorWithMetadata{
							Rule:   relayConnType,
							Line:   fieldDefinition.Type.Position.Line,
							Column: fieldDefinition.Type.Position.Column,
							Err:    fmt.Errorf("%d:%d pageInfo field from Connection type %s needs to return a non-null PageInfo object", fieldDefinition.Type.Position.Line, fieldDefinition.Type.Position.Column, typeDefinition.Name),
						}
						errors = append(errors, lintError)
					}
				}
			}

			if !foundEdgesField {
				lintError := LintErrorWithMetadata{
					Rule:   relayConnType,
					Line:   typeDefinition.Position.Line,
					Column: typeDefinition.Position.Column,
					Err:    fmt.Errorf("%d:%d type %s is a Connection type and therefore needs to have a field named 'edges' that returns a list type", typeDefinition.Position.Line, typeDefinition.Position.Column, typeDefinition.Name),
				}
				errors = append(errors, lintError)
			}

			if !foundPageInfoField {
				lintError := LintErrorWithMetadata{
					Rule:   relayConnType,
					Line:   typeDefinition.Position.Line,
					Column: typeDefinition.Position.Column,
					Err:    fmt.Errorf("%d:%d type %s is a Connection type and therefore needs to have a field named 'pageInfo' that returns a non-null PageInfo object", typeDefinition.Position.Line, typeDefinition.Position.Column, typeDefinition.Name),
				}
				errors = append(errors, lintError)
			}
		}
	}

	return errors
}

// RelayConnectionArgumentsSpec will validate the schema adheres to section 4 (Arguments) of the Relay Cursor Connections Specification.
// See https://relay.dev/graphql/connections.htm#sec-Arguments and https://relay.dev/graphql/connections.htm
func RelayConnectionArgumentsSpec(schema *ast.SchemaDocument) LintErrorsWithMetadata {
	errors := make([]LintErrorWithMetadata, 0)

	for _, typeDefinition := range schema.Definitions {
		for _, fieldDefinition := range typeDefinition.Fields {
			var firstArgument, afterArgument, lastArgument, beforeArgument *ast.ArgumentDefinition
			if strings.HasSuffix(fieldDefinition.Type.Name(), "Connection") {
				for _, argumentDefinition := range fieldDefinition.Arguments {
					switch argumentDefinition.Name {
					case "first":
						firstArgument = argumentDefinition
					case "after":
						afterArgument = argumentDefinition
					case "last":
						lastArgument = argumentDefinition
					case "before":
						beforeArgument = argumentDefinition
					}
				}

				hasForwardPagination := firstArgument != nil && afterArgument != nil
				hasBackwardPagination := lastArgument != nil && beforeArgument != nil

				if !hasForwardPagination && !hasBackwardPagination {
					lintError := LintErrorWithMetadata{
						Rule:   relayConnArgs,
						Line:   fieldDefinition.Type.Position.Line,
						Column: fieldDefinition.Type.Position.Column,
						Err:    fmt.Errorf("%d:%d field %s returns a Connection type and therefore must include forward pagination arguments (`first` and `after`) and/or backward pagination arguments (`last` and `before`) as per the Relay spec", fieldDefinition.Type.Position.Line, fieldDefinition.Type.Position.Column, fieldDefinition.Name), // nolint: lll
					}
					errors = append(errors, lintError)
				}

				if firstArgument != nil {
					if hasBackwardPagination {
						if firstArgument.Type.NamedType == "" || firstArgument.Type.NonNull || firstArgument.Type.Name() != "Int" {
							lintError := LintErrorWithMetadata{
								Rule:   relayConnArgs,
								Line:   firstArgument.Position.Line,
								Column: firstArgument.Position.Column,
								Err:    fmt.Errorf("%d:%d field %s is returns a Connection type that has both forward and backward pagination and therefore `first` argument should take a nullable non-negative integer as per the Relay spec", firstArgument.Position.Line, firstArgument.Position.Column, fieldDefinition.Name),
							}
							errors = append(errors, lintError)
						}
					} else {
						if isArgListType(firstArgument) || firstArgument.Type.Name() != "Int" {
							lintError := LintErrorWithMetadata{
								Rule:   relayConnArgs,
								Line:   firstArgument.Position.Line,
								Column: firstArgument.Position.Column,
								Err:    fmt.Errorf("%d:%d field %s is returns a Connection type and has forward pagination and therefore `first` argument should take a non-negative integer as per the Relay spec", firstArgument.Position.Line, firstArgument.Position.Column, fieldDefinition.Name),
							}
							errors = append(errors, lintError)
						}
					}
				}

				if lastArgument != nil {
					if hasForwardPagination {
						if isArgListType(lastArgument) || lastArgument.Type.NonNull || lastArgument.Type.Name() != "Int" {
							lintError := LintErrorWithMetadata{
								Rule:   relayConnArgs,
								Line:   lastArgument.Position.Line,
								Column: lastArgument.Position.Column,
								Err:    fmt.Errorf("%d:%d field %s is returns a Connection type that has both forward and backward pagination and therefore `last` argument should take a nullable non-negative integer as per the Relay spec", lastArgument.Position.Line, lastArgument.Position.Column, fieldDefinition.Name),
							}
							errors = append(errors, lintError)
						}
					} else {
						if isArgListType(lastArgument) || lastArgument.Type.Name() != "Int" {
							lintError := LintErrorWithMetadata{
								Rule:   relayConnArgs,
								Line:   lastArgument.Position.Line,
								Column: lastArgument.Position.Column,
								Err:    fmt.Errorf("%d:%d field %s is returns a Connection type and has backward pagination and therefore `last` argument should take a non-negative integer as per the Relay spec", lastArgument.Position.Line, lastArgument.Position.Column, fieldDefinition.Name),
							}
							errors = append(errors, lintError)
						}
					}
				}
			}
		}

	}

	return errors
}

func isFieldListType(fieldDefinition *ast.FieldDefinition) bool {
	return fieldDefinition.Type.NamedType == "" && fieldDefinition.Type.Elem != nil
}

func isArgListType(fieldArgument *ast.ArgumentDefinition) bool {
	return fieldArgument.Type.NamedType == "" && fieldArgument.Type.Elem != nil
}
