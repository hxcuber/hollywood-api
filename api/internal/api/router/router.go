package router

import (
	"github.com/gorilla/mux"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/actor"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/helloWorld"
	"net/http"
)

func Handler() *mux.Router {
	r := mux.NewRouter()

	helloWorldSubRouter := r.PathPrefix("/helloWorld").Subrouter()
	helloWorldSubRouter.Handle("", helloWorld.GetHelloWorld()).Methods(http.MethodGet)
	helloWorldSubRouter.Handle("", helloWorld.PostHelloWorld()).Methods(http.MethodPost)

	actorSubRouter := r.PathPrefix("/actors").Subrouter()
	actorSubRouter.Handle("", actor.GetActors()).Methods(http.MethodGet)

	return r
}
