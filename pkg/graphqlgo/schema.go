package graphqlgo

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
)

func SchemaHandler(schema graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := graphql.Do(
			graphql.Params{
				Schema:        schema,
				RequestString: params.Query,
			})
		json.NewEncoder(w).Encode(response)
	}
}
