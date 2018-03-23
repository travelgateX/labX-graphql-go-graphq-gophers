package resolver

type ReviewsResolver struct {
	stars int
	commentary *string
}

func(r *ReviewsResolver) Stars() int{
	return r.stars
}

func (r *ReviewsResolver) Commentary() *string{
	return r.commentary
}