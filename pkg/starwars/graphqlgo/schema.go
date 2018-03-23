package graphqlgo

import (
	"labX-graphql-go-graphq-gophers/pkg/starwars"

	"github.com/graphql-go/graphql"
)

var (
	humanType *graphql.Object
	droidType *graphql.Object
)

func NewSchema(service starwars.Service) graphql.Schema {
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
			if character, ok := p.Value.(*starwars.Human); ok {
				human := service.Human(character.ID)
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
					if human, ok := p.Source.(*starwars.Human); ok {
						return human.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the human.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(*starwars.Human); ok {
						return human.Name, nil
					}
					return nil, nil
				},
			},
			"friends": &graphql.Field{
				Type:        graphql.NewList(characterInterface),
				Description: "The friends of the human, or an empty list if they have none.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(*starwars.Human); ok {
						var list []interface{}
						for _, char := range *human.Friends {
							list = append(list, service.Character(char))
						}
						return list, nil
					}
					return []interface{}{}, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(*starwars.Human); ok {
						var list []interface{}
						for _, char := range human.AppearsIn {
							list = append(list, service.Character(char))
						}
						return list, nil
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
					if droid, ok := p.Source.(*starwars.Droid); ok {
						return droid.ID, nil
					}
					return nil, nil
				},
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(*starwars.Droid); ok {
						return droid.Name, nil
					}
					return nil, nil
				},
			},
			"appearsIn": &graphql.Field{
				Type:        graphql.NewList(episodeEnum),
				Description: "Which movies they appear in.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(*starwars.Droid); ok {
						return droid.AppearsIn, nil
					}
					return nil, nil
				},
			},
			"primaryFunction": &graphql.Field{
				Type:        graphql.String,
				Description: "The primary function of the droid.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if droid, ok := p.Source.(*starwars.Droid); ok {
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
					return service.Hero(p.Args["episode"].(string)), nil
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
					return service.Human(p.Args["id"].(string)), nil
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
					return service.Droid(p.Args["id"].(string)), nil
				},
			},
		},
	})
	ret, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	return ret
}
