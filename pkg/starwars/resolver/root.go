package resolver

import "labX-graphql-go-graphq-gophers/pkg/starwars"

type QueryResolver struct {
	Service starwars.Service
}

func (r *QueryResolver) StarWars() *StarWarsQueryResolver {
	return &StarWarsQueryResolver{
		service: r.Service,
	}
}
