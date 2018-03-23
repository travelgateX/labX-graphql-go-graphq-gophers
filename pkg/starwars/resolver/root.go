package resolver

type QueryResolver struct{}
type MutationResolver struct{}

func (r *QueryResolver) StarWars() *StarWarsQueryResolver {
	return &StarWarsQueryResolver{}
}
func (r *MutationResolver) StarWars() *StarWarsMutationResolver{
	return &StarWarsMutationResolver{}
}
