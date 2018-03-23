package resolver

import (
	"labX-graphql-go-graphq-gophers/pkg/starwars"
)

type StarWarsQueryResolver struct {
	service starwars.Service
}

func (r *StarWarsQueryResolver) Hero(episode string) *CharacterResolver {
	character := r.service.Hero(episode)
	charRes := &CharacterResolver{
		appearsIn: character.AppearsIn,
		id:        character.ID,
		name:      character.Name,
	}
	if character.Friends != nil && len(*character.Friends) > 0 {
		friendsRes := make([]CharacterResolver, 0, len(*character.Friends))
		for _, friend := range *character.Friends {
			friendsRes = append(friendsRes, CharacterResolver{
				id:        friend.ID,
				name:      friend.Name,
				appearsIn: friend.AppearsIn,
				friends:   friend.Friends,
			})
		}
	}
	return charRes
}

func (r *StarWarsQueryResolver) Reviews(episode string) *[]ReviewResolver {
	reviews := r.service.Reviews(episode)
	if len(reviews) == 0 {
		return nil
	}
	resReviews := make([]ReviewResolver, 0, len(reviews))
	for _, review := range reviews {
		resReviews = append(resReviews, ReviewResolver{
			stars:      review.Stars,
			commentary: review.Commentary,
		})
	}
	return &resReviews
}

func (r *StarWarsQueryResolver) Search(text string) *[]SearchResultResolver {
	return nil
}

func (r *StarWarsQueryResolver) Character(id string) *CharacterResolver {
	character := r.service.Character(id)
	return &CharacterResolver{
		id:        character.ID,
		appearsIn: character.AppearsIn,
		name:      character.Name,
		friends:   character.Friends,
	}
}

func (r *StarWarsQueryResolver) Droid(id string) *DroidResolver {
	droid := r.service.Droid(id)
	return &DroidResolver{
		id:              droid.ID,
		name:            droid.Name,
		appearsIn:       droid.AppearsIn,
		primaryFunction: &droid.PrimaryFunction,
		friends:         droid.Friends,
	}
}

func (r *StarWarsQueryResolver) Human(id string) *HumanResolver {
	human := r.service.Human(id)
	return &HumanResolver{
		id:        human.ID,
		name:      human.Name,
		appearsIn: human.AppearsIn,
		friends:   human.Friends,
	}
}

func (r *StarWarsQueryResolver) Starship(id string) *StarshipResolver {
	starship := r.service.Starship(id)
	return &StarshipResolver{
		id:     starship.ID,
		name:   starship.Name,
		length: starship.Length,
	}
}

func (r *StarWarsQueryResolver) CreateReview(episode string, reviewInput *starwars.ReviewInput) *ReviewResolver {
	review := r.service.CreateReview(episode, reviewInput)
	return &ReviewResolver{
		commentary: review.Commentary,
		stars:      reviewInput.Stars,
	}
}
