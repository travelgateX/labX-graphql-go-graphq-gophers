package resolver

type CharacterResolver struct {
	Character
}

func (r *CharacterResolver) ToHuman() (*HumanResolver, bool) {
	c, ok := r.Character.(*HumanResolver)
	return c, ok
}

func (r *CharacterResolver) ToDroid() (*DroidResolver, bool) {
	c, ok := r.Character.(*DroidResolver)
	return c, ok
}

// func (r *CharacterResolver) ID() graphql.ID {
// 	return r.Character.ID()
// }

// func (r *CharacterResolver) Name() string {
// 	return r.c.Name
// }

// func (r *CharacterResolver) Friends() *[]*CharacterResolver {
// 	if r.c.Friends == nil || len(*r.c.Friends) == 0 {
// 		return nil
// 	}

// 	friends := make([]*CharacterResolver, 0, len(*r.c.Friends))
// 	for _, friend := range *r.c.Friends {
// 		c := r.s.Character(friend)
// 		if c != nil {
// 			friends = append(friends, &CharacterResolver{c: c, s: r.s})
// 		}
// 	}
// 	return &friends
// }

// func (r *CharacterResolver) AppearsIn() []string {
// 	return r.c.AppearsIn
// }

// func (r *CharacterResolver) ToDroid() (*DroidResolver, bool) {
// 	c, ok := r.c.(*droidResolver)
// 	return c, ok
// }
