package parser

import (
	"fmt"
	"strings"
)

type QueryType string

const (
	Select QueryType = "SELECT"
	Insert QueryType = "INSERT"
	Update QueryType = "UPDATE"
	Delete QueryType = "DELETE"
)

type CQLQuery struct {
	Type    QueryType
	Table   string
	Columns []string
	Values  []string
	Where   string
}

func ParseCQL(cql string) ([]CQLQuery, error) {
	queries := []CQLQuery{}
	lines := strings.Split(cql, ";")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid query: %s", line)
		}

		switch strings.ToUpper(parts[0]) {
		case "SELECT":
			query := CQLQuery{
				Type:  Select,
				Table: parts[3],
			}
			query.Columns = strings.Split(parts[1], ",")
			queries = append(queries, query)
		case "INSERT":
			query := CQLQuery{
				Type:  Insert,
				Table: parts[2],
			}
			queries = append(queries, query)
		// Add support for UPDATE, DELETE here
		default:
			return nil, fmt.Errorf("unsupported query type: %s", parts[0])
		}
	}

	return queries, nil
}
