package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/actorHandler"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/movieHandler"
	"net/http"

	"github.com/hxcuber/hollywood-api/api/internal/api/rest/healthHandler"
)

// Router defines the routes & handlers of the app
type Router struct {
	ctx                  context.Context
	corsOrigins          []string
	isGQLIntrospectionOn bool
	healthRESTHandler    healthHandler.Handler
	actorRESTHandler     actorHandler.Handler
	movieRESTHandler     movieHandler.Handler
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	r := chi.NewRouter()
	// TODO: add middleware here
	r.Use(
		// render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger, // log relationship request calls
		// middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes, // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,    // recover from panics without crashing server
	)

	r.Get("/actors", rtr.actorRESTHandler.GetAllActors())
	r.Get("/actor/{id}", rtr.actorRESTHandler.GetActorById())
	r.Post("/actor/create", rtr.actorRESTHandler.PostActor())

	r.Get("/movies", rtr.movieRESTHandler.GetAllMovies())
	r.Get("/movie/{id}", rtr.movieRESTHandler.GetMovieById())

	r.Get("/_/ready", rtr.healthRESTHandler.CheckReadiness())
	return r
}
