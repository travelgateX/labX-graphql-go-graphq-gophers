package starwars

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
