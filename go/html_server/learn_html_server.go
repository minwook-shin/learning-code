package main

import (
    "flag"
    "html/template"
    "log"
    "net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

var templ = template.Must(template.New("").Parse(templateStr))

func main() {
    flag.Parse()
    http.Handle("/", http.HandlerFunc(index))
    err := http.ListenAndServe(*addr, nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

func index(w http.ResponseWriter, req *http.Request) {
    templ.Execute(w, req)
}

const templateStr = `
<html>
<head>
<title>hello world</title>
</head>
<body>
<p>hello world!</p>
</body>
</html>
`