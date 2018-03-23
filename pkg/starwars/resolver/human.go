package resolver

import (
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type HumanResolver struct {
	id        string
	name      string
	height    float64
	mass      *float64
	friends   *[]*starwars.Character
	appearsIn []string
	starships *[]StarshipResolver
}

func (r *HumanResolver) ID() string {
	return r.id
}

func (r *HumanResolver) Name() string {
	return r.name
}

type LengthUnit string

const (
	CENTIMETER LengthUnit = "CENTIMETER"
	METER      LengthUnit = "METER"
	KILOMETER  LengthUnit = "KILOMETER"
)

func (r *HumanResolver) Height(unit LengthUnit) float64 {
	switch unit {
	case CENTIMETER:
		return r.height / 1000
	case KILOMETER:
		return r.height * 1000
	default:
		return r.height
	}
}

func (r *HumanResolver) Mass() *float64 {
	return r.mass
}
func (r *HumanResolver) Friends() *[]starwars.Character {
	if r.friends == nil || len(*r.friends) > 0 {
		return nil
	}
	friends := make([]starwars.Character, 0, len(*r.friends))
	for _, friend := range *r.friends {
		friends = append(friends, starwars.Character{
			ID:        friend.ID,
			Friends:   friend.Friends,
			AppearsIn: friend.AppearsIn,
			Name:      friend.Name,
		})
	}
	return &friends
}

func (r *HumanResolver) Starships() *[]StarshipResolver {
	if r.starships == nil || len(*r.starships) > 0 {
		return nil
	}
	starships := make([]StarshipResolver, 0, len(*r.starships))
	for _, starship := range *r.starships {
		starships = append(starships,
			StarshipResolver{
				name:   starship.name,
				id:     starship.id,
				length: starship.length,
			})
	}
	return &starships
}
