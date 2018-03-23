package resolver

type StarWarsQueryResolver struct{}

func(r *StarWarsQueryResolver)Hero() *CharacterResolver{

}

func(r *StarWarsQueryResolver)Reviews() *ReviewsResolver{

}

func (r *StarWarsQueryResolver) Search() *[]SearchResultResolver{

}

func (r *StarWarsQueryResolver) Character() *[]CharacterResolver{

}

func(r *StarWarsQueryResolver)Droid() *DroidResoler{

}

func (r *StarWarsQueryResolver) Human() *HumanResolver{

}

func(r *StarWarsQueryResolver) Starship() *StarshipResolver{

}

type StarWarsMutationResolver struct{}
