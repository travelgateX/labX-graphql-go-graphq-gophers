package resolver

import "labX-graphql-go-graphq-gophers/pkg/starwars"

type DroidResolver struct {
	id              string
	name            string
	friends         *[]CharacterResolver
	appearsIn       []string
	primaryFunction *string
}

func (r *DroidResolver) ID() string {
	return r.id
}

func (r *DroidResolver) Name() string {
	return r.name
}

func (r *DroidResolver) Friends() *[]starwars.Character {
	if r.friends == nil || len(*r.friends) > 0 {
		return nil
	}
	friends := make([]starwars.Character, 0, len(*r.friends))
	for _, friend := range *r.friends {
		friends = append(friends, starwars.Character{
			Name:      friend.Name,
			ID:        friend.ID,
			AppearsIn: friend.AppearsIn,
			Friends:   friend.Friends,
		})
	}
	return &friends
}

func (r *DroidResolver) AppearsIn() []string {
	return r.appearsIn
}

func (r *DroidResolver) PrimaryFunction() *string {
	return r.primaryFunction
}
