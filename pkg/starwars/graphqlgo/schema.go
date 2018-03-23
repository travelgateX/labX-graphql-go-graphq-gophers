package graphqlgo

import (
	"labX-graphql-go-graphq-gophers/pkg/starwars"
	"strconv"

	"github.com/graphql-go/graphql"
)

var (
	Luke           starwars.Character
	Vader          starwars.Character
	Han            starwars.Character
	Leia           starwars.Character
	Tarkin         starwars.Character
	Threepio       starwars.Character
	Artoo          starwars.Character
	HumanData      map[int]starwars.Character
	DroidData      map[int]starwars.Character
	StarWarsSchema graphql.Schema

	humanType *graphql.Object
	droidType *graphql.Object
)

func NewSchema(service starwars.Service) graphql.Schema {
	Luke = starwars.Character{
		ID:        "1000",
		Name:      "Luke Skywalker",
		AppearsIn: []string{"4", "5", "6"},
	}
	Vader = starwars.Character{
		ID:        "1001",
		Name:      "Darth Vader",
		AppearsIn: []string{"4", "5", "6"},
	}
	Han = starwars.Character{
		ID:        "1002",
		Name:      "Han Solo",
		AppearsIn: []string{"4", "5", "6"},
	}
	Leia = starwars.Character{
		ID:        "1003",
		Name:      "Leia Organa",
		AppearsIn: []string{"4", "5", "6"},
	}
	Tarkin = starwars.Character{
		ID:        "1004",
		Name:      "Wilhuff Tarkin",
		AppearsIn: []string{"4"},
	}
	Threepio = starwars.Character{
		ID:        "2000",
		Name:      "C-3PO",
		AppearsIn: []string{"4", "5", "6"},
	}
	Artoo = starwars.Character{
		ID:        "2001",
		Name:      "R2-D2",
		AppearsIn: []string{"4", "5", "6"},
	}

	// Luke.Friends = append(Luke.Friends, []starwars.Character{Han, Leia, Threepio, Artoo}...)
	// Vader.Friends = append(Luke.Friends, []starwars.Character{Tarkin}...)
	// Han.Friends = append(Han.Friends, []starwars.Character{Luke, Leia, Artoo}...)
	// Leia.Friends = append(Leia.Friends, []starwars.Character{Luke, Han, Threepio, Artoo}...)
	// Tarkin.Friends = append(Tarkin.Friends, []starwars.Character{Vader}...)
	// Threepio.Friends = append(Threepio.Friends, []starwars.Character{Luke, Han, Leia, Artoo}...)
	// Artoo.Friends = append(Artoo.Friends, []starwars.Character{Luke, Han, Leia}...)
	HumanData = map[int]starwars.Character{
		1000: Luke,
		1001: Vader,
		1002: Han,
		1003: Leia,
		1004: Tarkin,
	}
	DroidData = map[int]starwars.Character{
		2000: Threepio,
		2001: Artoo,
	}

	episodeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name:        "Episode",
		Description: "One of the films in the Star Wars Trilogy",
		Values: graphql.EnumValueConfigMap{
			"NEWHOPE": &graphql.EnumValueConfig{
				Value:       4,
				Description: "Released in 1977.",
			},
			"EMPIRE": &graphql.EnumValueConfig{
				Value:       5,
				Description: "Released in 1980.",
			},
			"JEDI": &graphql.EnumValueConfig{
				Value:       6,
				Description: "Released in 1983.",
			},
		},
	})

	characterInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name:        "Character",
		Description: "A character in the Star Wars Trilogy",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the character.",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the character.",
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
			},
		},
		ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
			if character, ok := p.Value.(starwars.Character); ok {
				id, _ := strconv.Atoi(character.ID)
				human := GetHuman(id)
				if human.ID != "" {
					return humanType
				}
			}
			return droidType
		},
	})
	characterInterface.AddFieldConfig("friends", &graphql.Field{
		Type:        graphql.NewList(characterInterface),
		Description: "The friends of the character, or an empty list if they have none.",
	})

	humanType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Human",
		Description: "A humanoid creature in the Star Wars universe.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(starwars.Human); ok {
						return human.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(starwars.Human); ok {
						return human.Name, nil
					}
					return nil, nil
				},
			},
			"friends": &graphql.Field{
				Type:        graphql.NewList(characterInterface),
				Description: "The friends of the human, or an empty list if they have none.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(starwars.Human); ok {
						return human.Friends, nil
					}
					return []interface{}{}, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(starwars.Human); ok {
						return human.AppearsIn, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})
	droidType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Droid",
		Description: "A mechanical creature in the Star Wars universe.",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The id of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(starwars.Character); ok {
						return droid.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(starwars.Character); ok {
						return droid.Name, nil
					}
					return nil, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(starwars.Character); ok {
						return droid.AppearsIn, nil
					}
					return nil, nil
				},
			},
			"primaryFunction": &graphql.Field{
				Type:        graphql.String,
				Description: "The primary function of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(starwars.Droid); ok {
						return droid.PrimaryFunction, nil
					}
					return nil, nil
				},
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hero": &graphql.Field{
				Type: characterInterface,
				Args: graphql.FieldConfigArgument{
					"episode": &graphql.ArgumentConfig{
						Description: "If omitted, returns the hero of the whole saga. If " +
							"provided, returns the hero of that particular episode.",
						Type: episodeEnum,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetHero(p.Args["episode"]), nil
				},
			},
			"human": &graphql.Field{
				Type: humanType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the human",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.Atoi(p.Args["id"].(string))
					if err != nil {
						return nil, err
					}
					return GetHuman(id), nil
				},
			},
			"droid": &graphql.Field{
				Type: droidType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "id of the droid",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return GetDroid(p.Args["id"].(int)), nil
				},
			},
		},
	})
	ret, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	return ret
}

func GetHuman(id int) starwars.Character {
	if human, ok := HumanData[id]; ok {
		return human
	}
	return starwars.Character{}
}
func GetDroid(id int) starwars.Character {
	if droid, ok := DroidData[id]; ok {
		return droid
	}
	return starwars.Character{}
}
func GetHero(episode interface{}) interface{} {
	if episode == 5 {
		return Luke
	}
	return Artoo
}
