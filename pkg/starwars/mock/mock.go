package mock

import (
	"encoding/json"
	"labX/labX-graphql-go-graphq-gophers/pkg/starwars"
	"os"
	"strings"
)

type service struct {
	droids    *starwars.Droids
	films     []*starwars.Film
	humans    *starwars.Humans
	starships *starwars.Starships
	reviews   map[string][]*starwars.Review
}

func NewService() (starwars.Service, error) {
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

func (s *service) Hero(episode string) *starwars.Character {
	if episode == "EMPIRE" {
		c, _ := s.humans.HumanList["1000"]
		return &c.Character
	}
	c, _ := s.droids.DroidList["2001"]
	return &c.Character
}

func (s *service) Reviews(episode string) []*starwars.Review {
	r, _ := s.reviews[episode]
	return r
}

func (s *service) Search(text string) []starwars.SearchResult {
	var l []starwars.SearchResult

	for _, v := range s.droids.DroidList {
		if strings.Contains(v.Name, text) {
			l = append(l, starwars.SearchResult(v))
		}
	}

	for _, v := range s.humans.HumanList {
		if strings.Contains(v.Name, text) {
			l = append(l, starwars.SearchResult(v))
		}
	}

	for _, v := range s.starships.StarshipList {
		if strings.Contains(v.Name, text) {
			l = append(l, starwars.SearchResult(v))
		}
	}
	return l
}

func (s *service) Character(id string) *starwars.Character {
	if d, ok := s.droids.DroidList[id]; ok {
		return &d.Character
	}
	if h, ok := s.humans.HumanList[id]; ok {
		return &h.Character
	}
	return nil
}

func (s *service) Droid(id string) *starwars.Droid {
	d, _ := s.droids.DroidList[id]
	return d
}

func (s *service) Human(id string) *starwars.Human {
	h, _ := s.humans.HumanList[id]
	return h
}

func (s *service) Starship(id string) *starwars.Starship {
	ss, _ := s.starships.StarshipList[id]
	return ss
}

func (s *service) CreateReview(episode string, ri *starwars.ReviewInput) *starwars.Review {
	r := &ri.Review
	sr, ok := s.reviews[episode]
	if !ok {
		tmp := []*starwars.Review{}
		sr = tmp
		s.reviews[episode] = tmp
	}
	sr = append(sr, r)
	return r
}
