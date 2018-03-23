package main

import (
	"io/ioutil"
	"labX/labX-graphql-go-graphq-gophers/cmd"
	"labX/labX-graphql-go-graphq-gophers/pkg/gopher"
	"labX/labX-graphql-go-graphq-gophers/pkg/starwars/mock"
	"labX/labX-graphql-go-graphq-gophers/pkg/starwars/resolver"
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

	s, err := mock.NewService()
	schema, err := graphql.ParseSchema(string(b), resolver.QueryResolver{})
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
