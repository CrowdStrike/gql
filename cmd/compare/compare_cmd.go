package compare

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/CrowdStrike/gqltools/pkg/compare"
	"github.com/CrowdStrike/gqltools/utils"
)

var (
	oldSchemaPath      string
	newSchemaPath      string
	onlyBreakingChange bool
	excludeFilePath    bool
)

//NewCompareCmd creates new compare command
func NewCompareCmd() *cobra.Command {
	compareCmd := &cobra.Command{
		Use:   "compare",
		Short: "compare two graphql schemas",
		Long:  "compare two graphql schemas",
		Run: func(cmd *cobra.Command, args []string) {
			if len(oldSchemaPath) == 0 || len(newSchemaPath) == 0 {
				fmt.Print("compare expects two version of schemas in the arguments\n")
				os.Exit(1)
			}
			if oldSchemaPath == newSchemaPath {
				fmt.Printf("Both old '%s' and new '%s' schema path are same\n", oldSchemaPath, newSchemaPath)
				os.Exit(1)
			}

			schemaOldContents, err := utils.ReadFiles(oldSchemaPath)
			if err != nil {
				fmt.Printf("failed to read schema files on filepath:%s, error:%v", oldSchemaPath, err)
				os.Exit(1)
			}

			schemaNewContents, err := utils.ReadFiles(newSchemaPath)
			if err != nil {
				fmt.Printf("failed to read schema files on filepath:%s, error:%v", newSchemaPath, err)
				os.Exit(1)
			}

			schemaOld, parseErr := utils.ParseSchema(schemaOldContents)
			if parseErr != nil {
				fmt.Printf("Error parsing schema content on path=%s, error:%v", oldSchemaPath, parseErr)
				os.Exit(1)
			}

			schemaNew, parseErr := utils.ParseSchema(schemaNewContents)
			if parseErr != nil {
				fmt.Printf("Error parsing schema content on path=%s, error:%v", newSchemaPath, parseErr)
				os.Exit(1)
			}
			exitStatus := 0
			changes := compare.FindChangesInSchemas(schemaOld, schemaNew)
			if len(changes) == 0 {
				fmt.Println("No changes found on schema compare!")
			} else {
				changeCriticalityMap := compare.GroupChanges(changes)
				errorCount := 0
				//print changes
				if onlyBreakingChange {
					errorCount = compare.ReportBreakingChanges(changeCriticalityMap[compare.Breaking], !excludeFilePath)
				} else {
					errorCount = compare.ReportBreakingChanges(changeCriticalityMap[compare.Breaking], !excludeFilePath)
					compare.ReportDangerousChanges(changeCriticalityMap[compare.Dangerous], !excludeFilePath)
					compare.ReportNonBreakingChanges(changeCriticalityMap[compare.NonBreaking], !excludeFilePath)
				}
				if errorCount == 0 {
					fmt.Println("No breaking changes found 🎉")
				} else {
					fmt.Printf("\n❌ Breaking changes in schema: %d\n", errorCount)
					exitStatus |= 1
				}
			}
			os.Exit(exitStatus)
		},
	}
	compareCmd.PersistentFlags().StringVarP(&oldSchemaPath, "oldversion", "o", "", "Path to your older version of GraphQL schema")
	compareCmd.PersistentFlags().StringVarP(&newSchemaPath, "newversion", "n", "", "Path to your new version of GraphQL schema")
	compareCmd.PersistentFlags().BoolVarP(&onlyBreakingChange, "breaking-change-only", "b", false, "Get breaking change only")
	compareCmd.PersistentFlags().BoolVarP(&excludeFilePath, "exclude-print-filepath", "e", false, "Exclude printing schema filepath positions")
	return compareCmd
}
