package router

import (
	"github.com/gorilla/mux"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/helloWorld"
)

func Handler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld.GetHelloWorld)

	return r
}
