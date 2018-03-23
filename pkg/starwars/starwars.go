package starwars

import (
	"encoding/json"
)

type Service interface {
	Hero(episode string) *Character
	Reviews(episode string) []*Review
	Search(text string) []SearchResult
	Character(id string) *Character
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
	ID      string        `json:"ID"`
	Name    string        `json:"Name"`
	Friends *[]*Character `json:"Friends"`
	// set of episodes
	AppearsIn []string `json:"AppearsIn"`
}

type Droids struct {
	DroidList map[string]*Droid `json:"droids"`
}

func (d *Droids) UnmarshalJSON(data []byte) error {
	var s []*Droid
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	d.DroidList = make(map[string]*Droid, len(s))
	for i := range s {
		droid := s[i]
		d.DroidList[droid.ID] = droid
	}
	return nil
}

type Droid struct {
<<<<<<< HEAD
	AppearsIn       []string      `json:"AppearsIn"`
	Friends         *[]*Character `json:"Friends"`
	ID              string        `json:"ID"`
	Name            string        `json:"Name"`
	PrimaryFunction string        `json:"PrimaryFunction"`
=======
	Character
	PrimaryFunction string `json:"PrimaryFunction"`
>>>>>>> 6329494fb09c8529652197ec9bc49370b27d21c4
}
type Films struct {
	FilimList []Film `json:"films"`
}

type Film struct {
	Episode string `json:"Episide"`
	Hero    string `json:"Hero"`
}

type Humans struct {
	HumanList map[string]*Human `json:"humans"`
}

func (h *Humans) UnmarshalJSON(data []byte) error {
	var s []*Human
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	h.HumanList = make(map[string]*Human, len(s))
	for i := range s {
		human := s[i]
		h.HumanList[human.ID] = human
	}
	return nil
}

type Human struct {
<<<<<<< HEAD
	AppearsIn []string      `json:"AppearsIn"`
	Friends   *[]*Character `json:"Friends"`
	Height    float64       `json:"Height"`
	ID        string        `json:"ID"`
	Mass      int           `json:"Mass"`
	Name      string        `json:"Name"`
	Starships []string      `json:"Starships"`
=======
	Character
	Height    float64  `json:"Height"`
	Mass      int      `json:"Mass"`
	Starships []string `json:"Starships"`
>>>>>>> 6329494fb09c8529652197ec9bc49370b27d21c4
}

type Starships struct {
	StarshipList map[string]*Starship `json:"starships"`
}

func (ss *Starships) UnmarshalJSON(data []byte) error {
	var s []*Starship
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	ss.StarshipList = make(map[string]*Starship, len(s))
	for i := range s {
		starship := s[i]
		ss.StarshipList[starship.ID] = starship
	}
	return nil
}

type Starship struct {
	ID     string  `json:"ID"`
	Length float64 `json:"Length"`
	Name   string  `json:"Name"`
}
