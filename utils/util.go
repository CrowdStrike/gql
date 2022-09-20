package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
)

// ParseSchema parse schema files and combine their sources
func ParseSchema(schemaFileContents map[string][]byte) (*ast.SchemaDocument, error) {
	var sources []*ast.Source
	for fn, sf := range schemaFileContents {
		s := &ast.Source{
			Name:  fn,
			Input: string(sf),
		}
		sources = append(sources, s)
	}

	schema, parseErr := parser.ParseSchemas(sources...)
	if parseErr != nil {
		return nil, parseErr
	}
	return schema, nil
}

// ReadFiles read file contents from the give filepath
func ReadFiles(schemaFilePath string) (map[string][]byte, error) {
	schemaFiles, err := filepath.Glob(schemaFilePath)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, fmt.Errorf("matching files do not exist at path:%s, error:%v", schemaFilePath, err)
	}
	if len(schemaFiles) == 0 {
		return nil, fmt.Errorf("matching file does not exist at path:%s", schemaFilePath)
	}
	schemaFileContents := make(map[string][]byte)
	for _, filename := range schemaFiles {
		fileObject, fileErr := os.Open(filename) // nolint:gosec
		if fileErr != nil {
			return nil, fmt.Errorf("failed to open file:%s on path:%s, error:%v", filename, schemaFilePath, err)
		}
		content, err := io.ReadAll(fileObject)
		if err != nil {
			return nil, fmt.Errorf("failed to read file:%s on path:%s, error:%v", filename, schemaFilePath, err)
		}
		if len(content) == 0 {
			fmt.Printf("empty file=%s in path=%s", filename, schemaFilePath)
		}
		schemaFileContents[filename] = content
	}
	return schemaFileContents, nil
}
