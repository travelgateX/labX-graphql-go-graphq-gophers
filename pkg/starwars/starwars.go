package starwars

type Service interface {
	Hero(*Episode) *Character
	Reviews(*Episode) []*Review
	Search(text string) []*SearchResult
	Character(id string) *Character
	Droid(id string) *Droid
	Human(id string) *Human
	Starship(id string) *Starship
	CreateReview(*Episode, *ReviewInput) *Review
}

type ReviewInput struct{}

type Review struct{}

type SearchResult struct{}

type Character struct{}

type Episode struct{}

type Droids struct {
	DroidList []Droid `json:"droids"`
}

type Droid struct {
	AppearsIn       []string `json:"AppearsIn"`
	Friends         []string `json:"Friends"`
	ID              string   `json:"ID"`
	Name            string   `json:"Name"`
	PrimaryFunction string   `json:"PrimaryFunction"`
}
type Films struct {
	FilimList []Film `json:"films"`
}

type Film struct {
	Episode string `json:"Episide"`
	Hero    string `json:"Hero"`
}

type Humans struct {
	HumanList []Human `json:"humans"`
}

type Human struct {
	AppearsIn []string `json:"AppearsIn"`
	Friends   []string `json:"Friends"`
	Height    float64  `json:"Height"`
	ID        string   `json:"ID"`
	Mass      int      `json:"Mass"`
	Name      string   `json:"Name"`
	Starships []string `json:"Starships"`
}

type Starships struct {
	StarshipList []Starship `json:"starships"`
}

type Starship struct {
	ID     string  `json:"ID"`
	Length float64 `json:"Length"`
	Name   string  `json:"Name"`
}
