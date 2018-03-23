package gopher

import (
	"encoding/json"
	"github.com/graph-gophers/graphql-go"
	"net/http"
)

func SchemaHandler(schema *graphql.Schema) http.HandlerFunc {
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

		response := schema.Exec(r.Context(), params.Query, params.OperationName, params.Variables)
		json.NewEncoder(w).Encode(response)
	}
}
