package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
{{.Name}}님께,
{{if .Option}}
이 텍스트는 선택된 사람들만 보입니다.
{{- else}}
안녕하세요.
{{- end}}
{{with .Gift -}}
{{.}}를 구입해주셔서 감사합니다.
{{end}}
감사합니다,
gopher.
`
	type textTemplate struct {
		Name, Gift string
		Option   bool
	}
	var auto = textTemplate{
		"익명의 구매자", 
		"banana computer", 
		true}

	t := template.Must(template.New("letter").Parse(letter))
	t.Execute(os.Stdout, auto)
}
