package resolver

type ReviewResolver struct {
	stars      int
	commentary *string
}

func (r *ReviewResolver) Stars() int {
	return r.stars
}

func (r *ReviewResolver) Commentary() *string {
	return r.commentary
}
