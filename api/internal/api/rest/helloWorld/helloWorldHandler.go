package helloWorld

import (
	"fmt"
	"net/http"
)

func GetHelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "get hello world!\n")
	}
}
func PostHelloWorld() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "post hello world!\n")
	}
}
