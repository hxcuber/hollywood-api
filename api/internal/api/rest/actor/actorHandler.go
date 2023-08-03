package actor

import (
	"fmt"
	"net/http"
)

func GetActors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Get actors")
	}
}
