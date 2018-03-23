package graphqlgoSchemas

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/parser"
)

var (
	Luke           StarWarsChar
	Vader          StarWarsChar
	Han            StarWarsChar
	Leia           StarWarsChar
	Tarkin         StarWarsChar
	Threepio       StarWarsChar
	Artoo          StarWarsChar
	HumanData      map[int]StarWarsChar
	DroidData      map[int]StarWarsChar
	StarWarsSchema graphql.Schema

	humanType *graphql.Object
	droidType *graphql.Object
)

type StarWarsChar struct {
	ID              string
	Name            string
	Friends         []StarWarsChar
	AppearsIn       []int
	HomePlanet      string
	PrimaryFunction string
}

func init() {
	episodeEnum := graphql.NewEnum(graphql.EnumConfig{
		Name: "Episode",
		Values: graphql.EnumValueConfigMap{
			"NEWHOPE": &graphql.EnumValueConfig{
				Value: 4,
			},
			"EMPIRE": &graphql.EnumValueConfig{
				Value: 5,
			},
			"JEDI": &graphql.EnumValueConfig{
				Value: 6,
			},
		},
	})

	lengthUnitEnum := graphql.NewEnum(graphql.EnumConfig{
		Name: "LengthUnit",
		Values: graphql.EnumValueConfigMap{
			"METER": &graphql.EnumValueConfig{
				Value: "METER",
			},
			"FOOT": &graphql.EnumValueConfig{
				Value: "FOOT",
			},
		},
	})

	reviewInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Review",
		Fields: graphql.Fields{
			"stars": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"commentary": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	characterInterface := graphql.NewInterface(graphql.InterfaceConfig{
		Name: "Character",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"appearsIn": &graphql.Field{
				Type: graphql.NewList(episodeEnum),
			},
		},
		// ResolveType: func(p graphql.ResolveTypeParams) *graphql.Object {
		// 	if character, ok := p.Value.(StarWarsChar); ok {
		// 		id, _ := strconv.Atoi(character.ID)
		// 		human := GetHuman(id)
		// 		if human.ID != "" {
		// 			return humanType
		// 		}
		// 	}
		// 	return droidType
		// },
	})
	characterInterface.AddFieldConfig("friends", &graphql.Field{
		Type: graphql.NewList(characterInterface),
	})

	starshipType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Character",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"length": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"unit": &graphql.ArgumentConfig{
						Type: lengthUnitEnum,
					},
				},
				Type: graphql.Float,
			},
		},
	})
	droidType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Droid",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"friends": &graphql.Field{
				Type: graphql.NewList(characterInterface),
			},
			"appearsIn": &graphql.Field{
				Type: graphql.NewList(episodeEnum),
			},
			"primaryFunction": &graphql.Field{
				Type: graphql.String,
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})
	humanType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Human",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if human, ok := p.Source.(StarWarsChar); ok {
						return human.Name, nil
					}
					return nil, nil
				},
			},
			"height": &graphql.Field{
				Args: graphql.FieldConfigArgument{
					"unit": &graphql.ArgumentConfig{
						Type: lengthUnitEnum,
					},
				},
				Type: graphql.Float,
			},
			"mass": &graphql.Field{
				Type: graphql.Float,
			},
			"friends": &graphql.Field{
				Type: graphql.NewList(characterInterface),
			},
			"appearsIn": &graphql.Field{
				Type: graphql.NewList(episodeEnum),
			},
			"starships": &graphql.Field{
				Type: graphql.NewList(starshipType),
			},
		},
		Interfaces: []*graphql.Interface{
			characterInterface,
		},
	})

	searchResultUnion := graphql.NewUnion(graphql.UnionConfig{
		Name:  "SearchResult",
		Types: []*graphql.Object{droidType, humanType, starshipType},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hero": &graphql.Field{
				Type: characterInterface,
				Args: graphql.FieldConfigArgument{
					"episode": &graphql.ArgumentConfig{
						Type: episodeEnum,
					},
				},
			},
			"reviews": &graphql.Field{
				Type: reviewInterface,
				Args: graphql.FieldConfigArgument{
					"episode": &graphql.ArgumentConfig{
						Type: episodeEnum,
					},
				},
			},
			"search": &graphql.Field{
				Type: searchResultUnion,
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
			},
			"character": &graphql.Field{
				Type: characterInterface,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
			"human": &graphql.Field{
				Type: humanType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
			"droid": &graphql.Field{
				Type: droidType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
			"starship": &graphql.Field{
				Type: starshipType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createReview": &graphql.Field{
				Type: reviewInterface,
				Args: graphql.FieldConfigArgument{
					"episode": &graphql.ArgumentConfig{
						Type: episodeEnum,
					},
					"review": &graphql.ArgumentConfig{
						Type: reviewInterface,
					},
				},
			},
		},
	})
	StarWarsSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

func GetHuman(id int) StarWarsChar {
	if human, ok := HumanData[id]; ok {
		return human
	}
	return StarWarsChar{}
}
func GetDroid(id int) StarWarsChar {
	if droid, ok := DroidData[id]; ok {
		return droid
	}
	return StarWarsChar{}
}
func GetHero(episode interface{}) interface{} {
	if episode == 5 {
		return Luke
	}
	return Artoo
}

// Test helper functions
func TestParse(t *testing.T, query string) *ast.Document {
	astDoc, err := parser.Parse(parser.ParseParams{
		Source: query,
		Options: parser.ParseOptions{
			// include source, for error reporting
			NoSource: false,
		},
	})
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	return astDoc
}
func TestExecute(t *testing.T, ep graphql.ExecuteParams) *graphql.Result {
	return graphql.Execute(ep)
}

func Diff(want, got interface{}) []string {
	return []string{fmt.Sprintf("\ngot: %v", got), fmt.Sprintf("\nwant: %v\n", want)}
}

func ASTToJSON(t *testing.T, a ast.Node) interface{} {
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("Failed to marshal Node %v", err)
	}
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		t.Fatalf("Failed to unmarshal Node %v", err)
	}
	return f
}

func ContainSubsetSlice(super []interface{}, sub []interface{}) bool {
	if len(sub) == 0 {
		return true
	}
subLoop:
	for _, subVal := range sub {
		found := false
	innerLoop:
		for _, superVal := range super {
			if subVal, ok := subVal.(map[string]interface{}); ok {
				if superVal, ok := superVal.(map[string]interface{}); ok {
					if ContainSubset(superVal, subVal) {
						found = true
						break innerLoop
					} else {
						continue
					}
				} else {
					return false
				}

			}
			if subVal, ok := subVal.([]interface{}); ok {
				if superVal, ok := superVal.([]interface{}); ok {
					if ContainSubsetSlice(superVal, subVal) {
						found = true
						break innerLoop
					} else {
						continue
					}
				} else {
					return false
				}
			}
			if reflect.DeepEqual(superVal, subVal) {
				found = true
				break innerLoop
			}
		}
		if !found {
			return false
		}
		continue subLoop
	}
	return true
}

func ContainSubset(super map[string]interface{}, sub map[string]interface{}) bool {
	if len(sub) == 0 {
		return true
	}
	for subKey, subVal := range sub {
		if superVal, ok := super[subKey]; ok {
			switch superVal := superVal.(type) {
			case []interface{}:
				if subVal, ok := subVal.([]interface{}); ok {
					if !ContainSubsetSlice(superVal, subVal) {
						return false
					}
				} else {
					return false
				}
			case map[string]interface{}:
				if subVal, ok := subVal.(map[string]interface{}); ok {
					if !ContainSubset(superVal, subVal) {
						return false
					}
				} else {
					return false
				}
			default:
				if !reflect.DeepEqual(superVal, subVal) {
					return false
				}
			}
		} else {
			return false
		}
	}
	return true
}

func EqualErrorMessage(expected, result *graphql.Result, i int) bool {
	return expected.Errors[i].Message == result.Errors[i].Message
}
