package resolver

import "github.com/travelgateX/labX-graphql-go-graphq-gophers/pkg/starwars"


type StarWarsQueryResolver struct{
	service starwars.Service
}

func(r *StarWarsQueryResolver)Hero(episode string) *CharacterResolver{
	character := r.service.Hero(episode)
	return &CharacterResolver{

	}
}

func(r *StarWarsQueryResolver)Reviews(episode string) *[]ReviewResolver{
	reviews := r.service.Reviews(episode)
	if len(reviews)== 0{
		return nil
	}
	resReviews := make([]ReviewResolver, 0, len(reviews))
	for _, review := range reviews{
		resReviews = append(resReviews, ReviewResolver{
			stars: review.stars,
			commentary: review.commentary,
		})
	}
	return &resReviews
}

func (r *StarWarsQueryResolver) Search(text string) *[]SearchResultResolver{
	searchResults := r.service.Search(text)
	resSearchResults := make([]SearchResultResolver, 0, len(searchResults))
	for _, result := range searchResults{
		resSearchResults = append(resSearchResults, SearchResultResolver{

		})
	}
	return &resSearchResults
}

func (r *StarWarsQueryResolver) Character(id string) *CharacterResolver{
	character := r.service.Character(id)
	return &CharacterResolver{
		id: character.id,
		appearsIn: character.appearsIn,

	}
}

func(r *StarWarsQueryResolver)Droid() *DroidResolver{

}

func (r *StarWarsQueryResolver) Human() *HumanResolver{

}

func(r *StarWarsQueryResolver) Starship() *StarshipResolver{

}

type StarWarsMutationResolver struct{}
