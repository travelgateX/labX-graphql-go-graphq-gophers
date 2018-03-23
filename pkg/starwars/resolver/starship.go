package resolver

type StarshipResolver struct {
	id     string
	name   string
	length float64
}

func (r *StarshipResolver) ID() string {
	return r.id
}

func (r *StarshipResolver) Name() string {
	return r.name
}

func (r *StarshipResolver) Length() float64{
	return r.length
}
