package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name          string
	Method        string
	Pattern       string
	GzipMandatory bool
	HandlerFunc   http.HandlerFunc
}

func NewRouter(routes []Route) http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var h http.Handler = route.HandlerFunc
		// h = gzipValidate(h, route.GzipMandatory)
		h = handlers.CompressHandler(h)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(h)
	}

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"accept", "content-type", "origin", "x-custom-header", "authorization"}),
	)(router)
	return handler
}

// gzipValidate checks if this route requires gzip or not.
// If is mandatory, checks if there is gzip in the accept-encoding header.
func gzipValidate(inner http.Handler, gzipMandatory bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if gzipMandatory {
			acceptGzip := false
			acceptEncoding := strings.Split(r.Header.Get("Accept-Encoding"), ",")
			for _, item := range acceptEncoding {
				if strings.TrimSpace(item) == "gzip" {
					acceptGzip = true
					break
				}
			}
			if !acceptGzip {
				fmt.Fprintln(w, "Add the header Accept-Encoding with the value gzip")
			} else {
				inner.ServeHTTP(w, r)
			}
		} else {
			inner.ServeHTTP(w, r)
		}
	})
}
