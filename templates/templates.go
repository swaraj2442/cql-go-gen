package templates

import (
	"os"
	"text/template"
)

const goTemplate = `package {{.Package}}

import (
	"github.com/gocql/gocql"
)

func {{.FuncName}}(session *gocql.Session{{if .Params}}, {{.Params}}{{end}}) ([]map[string]interface{}, error) {
	query := "{{.Query}}"
	iter := session.Query(query).Iter()
	var results []map[string]interface{}
	m := map[string]interface{}{}
	for iter.MapScan(m) {
		results = append(results, m)
		m = map[string]interface{}{}
	}
	if err := iter.Close(); err != nil {
		return nil, err
	}
	return results, nil
}
`

type TemplateData struct {
	Package  string
	FuncName string
	Query    string
	Params   string
}

func GenerateCode(data TemplateData, outputFile string) error {
	tmpl, err := template.New("goCode").Parse(goTemplate)
	if err != nil {
		return err
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, data)
	if err != nil {
		return err
	}

	return nil
}
