package mock

import (
	"encoding/json"
	"labX/labX-graphql-go-graphq-gophers/pkg/starwars"
	"os"
)

type service struct {
	droids    []*starwars.Droid
	films     []*starwars.Film
	humans    []*starwars.Human
	starships []*starwars.Starships
}

func NewService() (starwars.service, error) {
	s := service{}

	err := decode("resource/droids.json", &s.droids)
	if err != nil {
		return nil, err
	}
	err = decode("resource/films.json", &s.films)
	if err != nil {
		return nil, err
	}
	err = decode("resource/humans.json", &s.humans)
	if err != nil {
		return nil, err
	}
	err = decode("resource/starships.json", &s.starships)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func decode(path string, target interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(target)
	return err
}

func (s *service) Hero(episode string) *Character {
	return nil
}
func (s *service) Reviews(episode string) []*Review {
	return nil
}
func (s *service) Search(text string) []*SearchResult {
	return nil
}
func (s *service) Character(id string) *Character {
	return nil
}
func (s *service) Droid(id string) *Droid {
	return nil
}
func (s *service) Human(id string) *Human {
	return nil
}
func (s *service) Starship(id string) *Starship {
	return nil
}
func (s *service) CreateReview(episode string, ri *ReviewInput) *Review {
	return nil
}
