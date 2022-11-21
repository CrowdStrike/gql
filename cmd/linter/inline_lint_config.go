package linter

import (
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/CrowdStrike/gqltools/cmd/linter/lexer"
)

type inlineLintConfigMetadata struct {
	pos   ast.Position
	value string
}

const (
	lintDisable     lintCommand = "lint-disable"
	lintEnable      lintCommand = "lint-enable"
	lintDisableLine lintCommand = "lint-disable-line"
)

func extractInlineLintConfiguration(source *ast.Source) []InlineLintConfig {
	inlineLintConfigs := make([]inlineLintConfigMetadata, 0)

	s := lexer.New(source)

	token, err := s.ReadToken()
	if err != nil { // If we reach here then, lexer should not return error now
		panic(err)
	}

	for {
		if token.Kind == lexer.Comment && strings.HasPrefix(token.Value, "#lint-") {
			inlineLintConfigs = append(inlineLintConfigs, inlineLintConfigMetadata{
				pos:   token.Pos,
				value: token.Value,
			})
		}
		token, err = s.ReadToken()
		if err != nil { // If we reach here then, lexer should not return error now
			panic(err)
		}
		if token.Kind == lexer.EOF {
			break
		}
	}

	inlineLintConfigRules := make([]InlineLintConfig, 0, len(inlineLintConfigs))

	// Now that we know which lines have lint configs in comments, lets parse them
	for _, linConfig := range inlineLintConfigs {
		matchGroups := inlineLintConfigurationRegex.FindStringSubmatch(linConfig.value)
		switch matchGroups[1] {
		case string(lintDisable):
			rulesString := matchGroups[2]
			rulesToApply := sanitizeRules(rulesString)
			inlineRule := InlineLintConfig{
				command: lintDisable,
				rules:   rulesToApply,
				pos:     linConfig.pos.Line,
			}
			inlineLintConfigRules = append(inlineLintConfigRules, inlineRule)
		case string(lintEnable):
			rulesString := matchGroups[2]
			rulesToApply := sanitizeRules(rulesString)
			inlineRule := InlineLintConfig{
				command: lintEnable,
				rules:   rulesToApply,
				pos:     linConfig.pos.Line,
			}
			inlineLintConfigRules = append(inlineLintConfigRules, inlineRule)
		case string(lintDisableLine):
			rulesString := matchGroups[2]
			rulesToApply := sanitizeRules(rulesString)
			inlineRule := InlineLintConfig{
				command: lintDisableLine,
				rules:   rulesToApply,
				pos:     linConfig.pos.Line,
			}
			inlineLintConfigRules = append(inlineLintConfigRules, inlineRule)
		}
	}
	return inlineLintConfigRules
}

func sanitizeRules(rulesString string) []string {
	rulesToApply := make([]string, 0)
	for _, rule := range strings.Split(rulesString, ",") {
		rulesToApply = append(rulesToApply, strings.TrimSpace(rule))
	}
	return rulesToApply
}
