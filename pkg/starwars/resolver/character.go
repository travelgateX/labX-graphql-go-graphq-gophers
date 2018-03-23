package resolver

import "labX-graphql-go-graphq-gophers/pkg/starwars"

type CharacterResolver struct {
	id        string
	name      string
	friends   *[]*starwars.Character
	appearsIn []string
}

func (r *CharacterResolver) ID() string {
	return r.id
}

func (r *CharacterResolver) Name() string {
	return r.name
}

func (r *CharacterResolver) Friends() *[]starwars.Character {
	if r.friends == nil || len(*r.friends) > 0 {
		return nil
	}
	friends := make([]starwars.Character, 0, len(*r.friends))
	for _, friend := range *r.friends {
		friends = append(friends, starwars.Character{
			ID:        friend.ID,
			Name:      friend.Name,
			AppearsIn: friend.AppearsIn,
			Friends:   friend.Friends,
		})
	}
	return &friends
}

func (r *CharacterResolver) AppearsIn() []string {
	return r.appearsIn
}
