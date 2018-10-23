package main

import (
	"html/template"
	"os"
)

func main() {
	var source = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			{{range .Items}}<div>{{ . }}</div>{{end}}
		</body>
	</html>`

	webTemplate, _ := template.New("webpage").Parse(source)

	structData := struct {
		Title string
		Body  []string
	}{
		Title: "title",
		Body: []string{
			"list_1",
			"list_2",
		},
	}
	webTemplate.Execute(os.Stdout, structData)
}
