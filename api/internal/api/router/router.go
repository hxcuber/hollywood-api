package router

import (
	"github.com/gorilla/mux"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/object"
	"net/http"
)

func Handler() *mux.Router {
	r := mux.NewRouter()

	helloWorldSubRouter := r.PathPrefix("/helloWorld").Subrouter()
	helloWorldSubRouter.Handle("", object.GetHelloWorld()).Methods(http.MethodGet)
	helloWorldSubRouter.Handle("", object.PostHelloWorld()).Methods(http.MethodPost)

	actorSubRouter := r.PathPrefix("/actors").Subrouter()
	actorSubRouter.Handle("", object.GetActors()).Methods(http.MethodGet)

	return r
}
