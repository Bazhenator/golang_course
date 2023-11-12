package main

//go:generate go run generator.go

import (
	"fmt"
	"os"
	"text/template"
)

func main() {

	codeTemplate := ` package main
		type {{.StructName}} struct {
    	{{range .Fields}}
    	{{.FieldName}} {{.FieldType}} ` + "`json:\"{{.JSONTag}}\"`" + `{{end}}
		}
	`
	data := struct {
		StructName string
		Fields     []struct {
			FieldName string
			FieldType string
			JSONTag   string
		}
	}{
		StructName: "Person",
		Fields: []struct {
			FieldName string
			FieldType string
			JSONTag   string
		}{
			{"Name", "string", "name"},
			{"Age", "int", "age"},
		},
	}

	f, err := os.Create("generated_code.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := template.Must(template.New("code").Parse(codeTemplate))
	err = t.Execute(f, data)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated JSON model code.")
}
