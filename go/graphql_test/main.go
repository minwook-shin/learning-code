package main

import (
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (q *query) Search() string {
	return "VALUE"
}

func main() {
	s := `
                schema {
                        query: Query
                }
                type Query {
                        search: String!
                }
        `
	schema := graphql.MustParseSchema(s, &query{})

	http.Handle("/",
		&relay.Handler{
			Schema: schema},
	)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
