package linter

import (
	"sort"
)

// LintErrorWithMetadata represent a lint error. It stores metadata such as Rule for which error happened, position and actual error.
type LintErrorWithMetadata struct {
	Rule         LintRule
	Line, Column int
	Err          error
}

// LintErrorsWithMetadata represent collection of lint errors.
type LintErrorsWithMetadata []LintErrorWithMetadata

func (e LintErrorsWithMetadata) Len() int {
	return len(e)
}

// Less returns true if position at which error[i] occurred is before error[j]
func (e LintErrorsWithMetadata) Less(i, j int) bool {
	if e[i].Line < e[j].Line {
		return true
	}
	if e[i].Line > e[j].Line {
		return false
	}
	return e[i].Column < e[j].Column
}

// Swap swaps values in error list for two given indices
func (e LintErrorsWithMetadata) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// GetSortedErrors returns sorted list of errors with metadata. Sorting is done for Line-Column position for lint error
func (e LintErrorsWithMetadata) GetSortedErrors() []LintErrorWithMetadata {
	sort.Sort(e)
	return e
}
