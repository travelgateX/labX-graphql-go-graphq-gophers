package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type StarshipResolver struct {
	ss *starwars.Starship
}

func (r *StarshipResolver) ID() graphql.ID {
	return graphql.ID(r.ss.ID)
}

func (r *StarshipResolver) Name() string {
	return r.ss.Name
}

func (r *StarshipResolver) Length(args struct{ Unit string }) float64 {
	return r.ss.Length
}
