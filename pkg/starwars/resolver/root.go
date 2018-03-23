package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type Resolver struct {
	Service starwars.Service
}

type Character interface {
	ID() graphql.ID
	Name() string
	Friends() *[]*CharacterResolver
	AppearsIn() []string
}

func (r *Resolver) Hero(args struct{ Episode string }) *CharacterResolver {
	c := r.Service.Hero(args.Episode)
	if c == nil {
		return nil
	}
	switch v := c.(type) {
	case *starwars.Droid:
		dr := DroidResolver{d: v, s: r.Service}
		return &CharacterResolver{&dr}
	case *starwars.Human:
		hr := HumanResolver{h: v, s: r.Service}
		return &CharacterResolver{&hr}
	}

	return nil
}

func (r *Resolver) Reviews(args struct{ Episode string }) []*ReviewResolver {
	reviews := r.Service.Reviews(args.Episode)
	if len(reviews) == 0 {
		return nil
	}

	rr := make([]*ReviewResolver, 0, len(reviews))
	for _, review := range reviews {
		rr = append(rr, &ReviewResolver{review})
	}
	return rr
}

func (r *Resolver) Search(args struct{ Text string }) []*SearchResultResolver {
	srs := r.Service.Search(args.Text)
	if len(srs) == 0 {
		return nil
	}

	ret := make([]*SearchResultResolver, 0, len(srs))
	for _, sr := range srs {
		switch t := sr.(type) {
		case *starwars.Human:
			hr := &HumanResolver{h: t, s: r.Service}
			ret = append(ret, &SearchResultResolver{human: hr})
		case *starwars.Droid:
			dr := &DroidResolver{d: t}
			ret = append(ret, &SearchResultResolver{droid: dr})
		case *starwars.Starship:
			ssr := &StarshipResolver{ss: t}
			ret = append(ret, &SearchResultResolver{starship: ssr})
		}
	}
	return ret
}

func (r *Resolver) Character(args struct{ ID string }) *CharacterResolver {
	c := r.Service.Character(args.ID)
	if c == nil {
		return nil
	}
	switch v := c.(type) {
	case *starwars.Droid:
		dr := DroidResolver{d: v, s: r.Service}
		return &CharacterResolver{&dr}
	case *starwars.Human:
		hr := HumanResolver{h: v, s: r.Service}
		return &CharacterResolver{&hr}
	}
	return nil
}

func (r *Resolver) Droid(args struct{ ID string }) *DroidResolver {
	droid := r.Service.Droid(args.ID)
	if droid == nil {
		return nil
	}
	return &DroidResolver{d: droid, s: r.Service}
}

func (r *Resolver) Human(args struct{ ID string }) *HumanResolver {
	human := r.Service.Human(args.ID)
	if human == nil {
		return nil
	}
	return &HumanResolver{h: human, s: r.Service}
}

func (r *Resolver) Starship(args struct{ ID string }) *StarshipResolver {
	ss := r.Service.Starship(args.ID)
	if ss == nil {
		return nil
	}
	return &StarshipResolver{ss}
}

func (r *Resolver) CreateReview(args struct {
	Episode string
	Review  *starwars.ReviewInput
}) *ReviewResolver {
	review := r.Service.CreateReview(args.Episode, args.Review)
	if review == nil {
		return nil
	}
	return &ReviewResolver{review}
}
