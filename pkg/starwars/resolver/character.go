package resolver

type CharacterResolver struct {
	id        string
	name      string
	friends   *[]CharacterResolver
	appearsIn []string
}

func (r *CharacterResolver) ID() string {
	return r.id
}

func (r *CharacterResolver) Name() string {
	return r.name
}

func (r *CharacterResolver) Friends() *[]CharacterResolver {
	if r.friends == nil || len(*r.friends) > 0 {
		return nil
	}
	friends := make([]CharacterResolver, 0, len(*r.friends))
	for _, friend := range *r.friends {
		friends = append(friends, friend)
	}
	return &friends
}

func (r *CharacterResolver) AppearsIn() []string {
	return r.appearsIn
}
