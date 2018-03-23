package main

import (
	"labX-graphql-go-graphq-gophers/cmd"
	"labX-graphql-go-graphq-gophers/pkg/graphqlgo"
	starwarsGraphql "labX-graphql-go-graphq-gophers/pkg/starwars/graphqlgo"
	"labX-graphql-go-graphq-gophers/pkg/starwars/mock"
	"net/http"
)

func main() {
	s, _ := mock.NewService()
	// TODO: Resolver
	schema := starwarsGraphql.NewSchema(s)

	routes := []cmd.Route{
		{
			Name:          "ServiceGraphQL",
			Method:        "POST",
			Pattern:       "/query",
			GzipMandatory: false,
			HandlerFunc:   graphqlgo.SchemaHandler(schema),
		},
	}

	handler := cmd.NewRouter(routes)
	panic(http.ListenAndServe(":8080", handler))
}
