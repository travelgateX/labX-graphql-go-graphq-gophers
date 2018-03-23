package main

import (
	"io/ioutil"
	"labX-graphql-go-graphq-gophers/cmd"
	"labX-graphql-go-graphq-gophers/pkg/gopher"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
	"labX-graphql-go-graphq-gophers/pkg/starwars/resolver"
	"net/http"
	"os"

	"github.com/graph-gophers/graphql-go"
)

func main() {
	f, err := os.Open("resource/task1-schema.graphql")
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	f.Close()
	if err != nil {
		panic(err)
	}

	// TODO: coger implementacion por ficheros
	s := new(starwars.Service)
	// TODO: Resolver
	schema, err := graphql.ParseSchema(string(b), resolver.QueryResolver{*s})
	if err != nil {
		panic(err)
	}

	routes := []cmd.Route{
		{
			Name:          "ServiceGraphQL",
			Method:        "POST",
			Pattern:       "/query",
			GzipMandatory: true,
			HandlerFunc:   gopher.SchemaHandler(schema),
		}, {
			Name:          "ServiceGraphiQL",
			Method:        "GET",
			Pattern:       "/",
			GzipMandatory: false,
			HandlerFunc:   gopher.GraphiQLHandler(),
		},
	}

	handler := cmd.NewRouter(routes)
	panic(http.ListenAndServe(":8080", handler))
}
