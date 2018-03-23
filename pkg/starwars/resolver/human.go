package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type HumanResolver struct {
	h *starwars.Human
	s starwars.Service
}

func (r *HumanResolver) ID() graphql.ID {
	return graphql.ID(r.h.ID)
}

func (r *HumanResolver) Name() string {
	return r.h.Name
}

type LengthUnit string

const (
	Centimeter string = "CENTIMETER"
	Meter      string = "METER"
	Kilometer  string = "KILOMETER"
)

func (r *HumanResolver) Height(args struct{ Unit string }) float64 {
	switch args.Unit {
	case Centimeter:
		return r.h.Height / 1000
	case Kilometer:
		return r.h.Height * 1000
	default:
		return r.h.Height
	}
}

func (r *HumanResolver) Mass() *float64 {
	if r.h.Mass == 0 {
		return nil
	}
	return &r.h.Mass
}
func (r *HumanResolver) Friends() *[]*CharacterResolver {
	if r.h.Friends == nil || len(*r.h.Friends) == 0 {
		return nil
	}

	friends := make([]*CharacterResolver, 0, len(*r.h.Friends))
	for _, friend := range *r.h.Friends {
		c := r.s.Character(friend)
		if c != nil {
			hr := HumanResolver{h: r.h, s: r.s}
			friends = append(friends, &CharacterResolver{&hr})
		}
	}
	return &friends
}

func (r *HumanResolver) Starships() *[]*StarshipResolver {
	if len(r.h.Starships) == 0 {
		return nil
	}

	starships := make([]*StarshipResolver, 0, len(r.h.Starships))
	for _, v := range r.h.Starships {
		ss := r.s.Starship(v)
		if ss != nil {
			starships = append(starships, &StarshipResolver{ss})
		}
	}
	return &starships
}

func (r *HumanResolver) AppearsIn() []string {
	return r.h.AppearsIn
}
