package main

import (
	"flag"
	"html/template"
	"net/http"
)

var templateStr = `
<html>
<head>
<title>hello world</title>
</head>
<body>
<p>hello world!</p>
</body>
</html>
`

var address = flag.String("addr", ":8080", "http service address")

var templateParser = template.Must(template.New("").Parse(templateStr))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(func (w http.ResponseWriter, req *http.Request)  {
		templateParser.Execute(w, req)
	}))

	http.ListenAndServe(*address, nil)
}
