package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type DroidResolver struct {
	d *starwars.Droid
	s starwars.Service
}

func (r *DroidResolver) ID() graphql.ID {
	return graphql.ID(r.d.ID)
}

func (r *DroidResolver) Name() string {
	return r.d.Name
}

func (r *DroidResolver) Friends() *[]*CharacterResolver {
	if r.d.Friends == nil || len(*r.d.Friends) == 0 {
		return nil
	}

	friends := make([]*CharacterResolver, 0, len(*r.d.Friends))
	for _, friend := range *r.d.Friends {
		c := r.s.Character(friend)
		if c != nil {
			dr := DroidResolver{d: r.d, s: r.s}
			friends = append(friends, &CharacterResolver{&dr})
		}
	}
	return &friends
}

func (r *DroidResolver) AppearsIn() []string {
	return r.d.AppearsIn
}

func (r *DroidResolver) PrimaryFunction() *string {
	if r.d.PrimaryFunction == "" {
		return nil
	}
	return &r.d.PrimaryFunction
}
