package resolver

type SearchResultResolver struct {
	human    *HumanResolver
	droid    *DroidResolver
	starship *StarshipResolver
}

func (r *SearchResultResolver) ToHuman() (*HumanResolver, bool) {
	return r.human, r.human != nil
}

func (r *SearchResultResolver) ToDroid() (*DroidResolver, bool) {
	return r.droid, r.droid != nil
}

func (r *SearchResultResolver) ToStarship() (*StarshipResolver, bool) {
	return r.starship, r.starship != nil
}
