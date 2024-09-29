package generator

import (
	"cql-gen/parser"
	"cql-gen/templates"
	"fmt"
	"strings"
)

func GenerateGoCode(queries []parser.CQLQuery, outputFile string) error {
	for _, query := range queries {
		funcName := fmt.Sprintf("%s%s", strings.Title(string(query.Type)), strings.Title(query.Table))
		queryStr := formatQuery(query)
		data := templates.TemplateData{
			Package:  "main", // You can make this configurable
			FuncName: funcName,
			Query:    queryStr,
			Params:   "", // Handle query parameters here if needed
		}
		err := templates.GenerateCode(data, outputFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func formatQuery(query parser.CQLQuery) string {
	switch query.Type {
	case parser.Select:
		return fmt.Sprintf("SELECT %s FROM %s", strings.Join(query.Columns, ", "), query.Table)
	case parser.Insert:
		// Implement for INSERT
	}
	return ""
}
