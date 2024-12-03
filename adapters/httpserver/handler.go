package httpserver

import (
	"fmt"
	"net/http"

	go_specs_greet "github.com/lobiritus/go-specs-greet/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, go_specs_greet.Greet(name))
}