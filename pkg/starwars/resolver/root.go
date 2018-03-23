package resolver

import "labX/labX-graphql-go-graphq-gophers/pkg/starwars"

type QueryResolver struct {
	service starwars.Service
}

func (r *QueryResolver) StarWars() *StarWarsQueryResolver {
	return &StarWarsQueryResolver{
		service: r.service,
	}
}
