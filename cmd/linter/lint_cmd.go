package linter

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/CrowdStrike/gql/pkg/linter"

	"github.com/CrowdStrike/gql/utils"

	"github.com/spf13/cobra"
)

var (
	schemaFilePath string
	passedRules    []string
)

// NewLintCmd creates new lint command
func NewLintCmd() *cobra.Command {
	// addCmd represents the add command
	lintCmd := &cobra.Command{
		Use:   "lint",
		Short: "lints given GraphQL schema",
		Long:  `lints given GraphQL schema`,
		Run: func(cmd *cobra.Command, args []string) {
			schemaFileContents := make(map[string][]byte)

			rulesToApply, err := findTheRulesToApply(passedRules)
			if err != nil {
				fmt.Printf("failed to parse the rules to apply string=%s, error:%v", passedRules, err)
				//Exit with error printed to stderr
				os.Exit(1)
			}

			if len(schemaFilePath) == 0 {
				content, err := io.ReadAll(os.Stdin)
				if err != nil {
					fmt.Printf("failed to input from stdin with error %v", err)
					os.Exit(1)
				}
				schemaFileContents[os.Stdin.Name()] = content
			} else {
				var err error
				schemaFileContents, err = utils.ReadFiles(schemaFilePath)
				if err != nil {
					fmt.Printf("failed to read schema file error %v", err)
					os.Exit(1)
				}
			}

			exitStatus := 0
			errorCount := 0
			for filename, schemaFileContent := range schemaFileContents {
				if lintErrors := Lint(filename, string(schemaFileContent), rulesToApply); len(lintErrors) != 0 {
					errorCount += len(lintErrors)
					errorPresenter(filename, lintErrors)
					exitStatus |= 1 // If there's error for any file, exit code should be 1
				}
			}

			if errorCount == 0 {
				fmt.Printf("Schema has no lint errors! 🎉\n")
			} else {
				fmt.Printf("❌ Total lint errors found: %d\n", errorCount)
			}

			os.Exit(exitStatus) // success
		},
	}
	lintCmd.PersistentFlags().StringVarP(&schemaFilePath, "filepath", "f", "", "Path to your GraphQL schema")
	lintCmd.PersistentFlags().StringSliceVarP(&passedRules, "rules", "r", []string{}, fmt.Sprintf("Rules you want linter to use e.g.(-r type-desc,field-desc); available rules:\n %s", linter.AvailableRulesWithDescription()))
	return lintCmd
}

func findTheRulesToApply(rulesString []string) ([]linter.LintRuleFunc, error) {
	rulesToApply := make([]linter.LintRuleFunc, 0)
	if len(rulesString) == 0 {
		for _, rule := range linter.AllTheRules {
			rulesToApply = append(rulesToApply, rule.RuleFunction)
		}
		return rulesToApply, nil
	}
	for _, ruleToken := range rulesString {
		inputRuleName := strings.TrimSpace(ruleToken)
		// Check whether the rule passed exists in our rule list
		matchFound := false
		for _, rule := range linter.AllTheRules {
			if strings.EqualFold(inputRuleName, string(rule.Name)) {
				matchFound = true
				rulesToApply = append(rulesToApply, rule.RuleFunction)
				break
			}
		}
		if !matchFound {
			return nil, fmt.Errorf("invalid rule[%s] passed", inputRuleName)
		}

	}
	return rulesToApply, nil
}

func errorPresenter(schemaFilePath string, errors []linter.LintErrorWithMetadata) {
	for _, err := range errors {
		fmt.Printf("%s:%d:%d %s\n", schemaFilePath, err.Line, err.Column, err.Err.Error())
	}
	fmt.Println("") // This is a separator between outputs of individual file
}
