package linter

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/CrowdStrike/gqltools/pkg/linter"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

var inlineLintConfigurationRegex, _ = regexp.Compile(`^#\s*(lint-[^\s ]+)(\s.*)?$`)

type lintCommand string

// InlineLintConfig represent config value defined in the schema for enabling/disabling certain rules for single/some Line/lines
type InlineLintConfig struct {
	command lintCommand
	rules   []string
	pos     int
}

// Lint lints federated GraphQL schema
func Lint(fileName string, schemaFileContents string, rules []linter.LintRuleFunc) []linter.LintErrorWithMetadata {
	source := &ast.Source{
		Name:  fileName,
		Input: schemaFileContents,
	}
	schema, parseErr := parser.ParseSchema(source)

	if parseErr != nil {
		fmt.Printf("failed to parse file=%s with error %v", fileName, parseErr)
		os.Exit(1)
	}

	allErrors := linter.LintErrorsWithMetadata{}
	for _, rule := range rules {
		errorsFromLintRule := rule(schema)
		allErrors = append(allErrors, errorsFromLintRule...)
	}

	sortedErrors := allErrors.GetSortedErrors()
	inlineLintConfigs := extractInlineLintConfiguration(source)
	filteredErrors := filterErrors(sortedErrors, inlineLintConfigs)
	return filteredErrors
}

func filterErrors(errors []linter.LintErrorWithMetadata, configs []InlineLintConfig) []linter.LintErrorWithMetadata {
	filteredErrors := make([]linter.LintErrorWithMetadata, 0)
	for _, lintErr := range errors {
		shouldApplyRule := true
		errorLine := lintErr.Line
		for _, config := range configs {
			// If the error for the lintRule isn't one of the specified rule then there's nothing to do for it
			if !contains(config.rules, lintErr.Rule) {
				continue
			}
			if config.command == "lint-disable-line" && config.pos == errorLine {
				shouldApplyRule = false
				break
			}
			if config.pos < errorLine {
				if config.command == "lint-enable" {
					shouldApplyRule = true
				} else if config.command == "lint-disable" {
					shouldApplyRule = false
				}
			}
		}
		if shouldApplyRule {
			filteredErrors = append(filteredErrors, lintErr)
		}
	}

	return filteredErrors
}

func contains(list []string, first linter.LintRule) bool {
	for _, second := range list {
		if strings.EqualFold(string(first), second) {
			return true
		}
	}
	return false
}
