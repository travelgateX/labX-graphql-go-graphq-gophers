package mock

import (
	"encoding/json"
	"labX-graphql-go-graphq-gophers/pkg/starwars"
	"os"
	"strings"
	"sync"
)

type service struct {
	droids    starwars.DroidMap
	films     starwars.Films
	humans    starwars.HumanMap
	starships starwars.StarshipMap
	reviewsMu sync.RWMutex
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

func (s *service) Hero(episode string) interface{} {
	if episode == "EMPIRE" {
		c, _ := s.humans.M["1000"]
		return c
	}
	d, _ := s.droids.M["2001"]
	return d
}

func (s *service) Reviews(episode string) []*starwars.Review {
	s.reviewsMu.RLock()
	r, _ := s.reviews[episode]
	s.reviewsMu.RUnlock()
	return r
}

func (s *service) Search(text string) []starwars.SearchResult {
	var l []starwars.SearchResult

	for _, v := range s.droids.M {
		if strings.Contains(v.Name, text) {
			l = append(l, starwars.SearchResult(v))
		}
	}

	for _, v := range s.humans.M {
		if strings.Contains(v.Name, text) {
			l = append(l, starwars.SearchResult(v))
		}
	}

	for _, v := range s.starships.M {
		if strings.Contains(v.Name, text) {
			l = append(l, starwars.SearchResult(v))
		}
	}
	return l
}

func (s *service) Character(id string) interface{} {
	if d, ok := s.droids.M[id]; ok {
		return d
	}
	if h, ok := s.humans.M[id]; ok {
		return h
	}
	return nil
}

func (s *service) Droid(id string) *starwars.Droid {
	d, _ := s.droids.M[id]
	return d
}

func (s *service) Human(id string) *starwars.Human {
	h, _ := s.humans.M[id]
	return h
}

func (s *service) Starship(id string) *starwars.Starship {
	ss, _ := s.starships.M[id]
	return ss
}

func (s *service) CreateReview(episode string, ri *starwars.ReviewInput) *starwars.Review {
	r := &ri.Review
	s.reviewsMu.Lock()
	sr, ok := s.reviews[episode]
	if !ok {
		tmp := []*starwars.Review{}
		sr = tmp
		s.reviews[episode] = tmp
	}
	sr = append(sr, r)
	s.reviewsMu.Lock()
	return r
}
