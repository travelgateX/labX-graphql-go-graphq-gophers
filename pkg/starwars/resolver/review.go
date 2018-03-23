package resolver

import (
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type ReviewResolver struct {
	r *starwars.Review
}

func (r *ReviewResolver) Stars() int32 {
	return int32(r.r.Stars)
}

func (r *ReviewResolver) Commentary() *string {
	return r.r.Commentary
}
