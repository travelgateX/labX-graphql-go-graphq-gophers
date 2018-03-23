package starwars

import (
	"encoding/json"
)

type Service interface {
	Hero(episode string) interface{}
	Reviews(episode string) []*Review
	Search(text string) []SearchResult
	Character(id string) interface{}
	Droid(id string) *Droid
	Human(id string) *Human
	Starship(id string) *Starship
	CreateReview(episode string, ri *ReviewInput) *Review
}

type Review struct {
	Stars      int
	Commentary *string
}

type ReviewInput struct {
	Review
}

// SearchResult is an union of Human | Droid | Starship
type SearchResult interface{}

type Character struct {
	ID      string    `json:"ID"`
	Name    string    `json:"Name"`
	Friends *[]string `json:"Friends"`
	// set of episodes
	AppearsIn []string `json:"AppearsIn"`
}

type Droids struct {
	List []*Droid `json:"droids"`
}

type DroidMap struct {
	M map[string]*Droid
}

func (d *DroidMap) UnmarshalJSON(data []byte) error {
	var droids Droids
	err := json.Unmarshal(data, &droids)
	if err != nil {
		return err
	}

	d.M = make(map[string]*Droid, len(droids.List))
	for i := range droids.List {
		droid := droids.List[i]
		d.M[droid.ID] = droid
	}
	return nil
}

type Droid struct {
	Character
	PrimaryFunction string `json:"PrimaryFunction"`
}
type Films struct {
	FilimList []Film `json:"films"`
}

type Film struct {
	Episode string `json:"Episide"`
	Hero    string `json:"Hero"`
}

type Humans struct {
	HumanList []*Human `json:"humans"`
}

type HumanMap struct {
	M map[string]*Human
}

func (h *HumanMap) UnmarshalJSON(data []byte) error {
	var s Humans
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	h.M = make(map[string]*Human, len(s.HumanList))
	for i := range s.HumanList {
		human := s.HumanList[i]
		h.M[human.ID] = human
	}
	return nil
}

type Human struct {
	Character
	Height    float64  `json:"Height"`
	Mass      float64  `json:"Mass"`
	Starships []string `json:"Starships"`
}

type Starships struct {
	StarshipList []*Starship `json:"starships"`
}

type StarshipMap struct {
	M map[string]*Starship
}

func (ss *StarshipMap) UnmarshalJSON(data []byte) error {
	var s Starships
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	ss.M = make(map[string]*Starship, len(s.StarshipList))
	for i := range s.StarshipList {
		starship := s.StarshipList[i]
		ss.M[starship.ID] = starship
	}
	return nil
}

type Starship struct {
	ID     string  `json:"ID"`
	Length float64 `json:"Length"`
	Name   string  `json:"Name"`
}
