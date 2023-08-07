package router

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/actor"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/health"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/movie"
	"github.com/hxcuber/hollywood-api/api/internal/controller/actors"
	"github.com/hxcuber/hollywood-api/api/internal/controller/movies"
	"github.com/hxcuber/hollywood-api/api/internal/controller/systems"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	corsOrigin []string,
	isGQLIntrospectionOn bool,
	systemCtrl systems.Controller,
	actorCtrl actors.Controller,
	movieCtrl movies.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          corsOrigin,
		isGQLIntrospectionOn: isGQLIntrospectionOn,
		healthRESTHandler:    health.New(systemCtrl),
		actorRESTHandler:     actor.New(actorCtrl),
		movieRESTHandler:     movie.New(movieCtrl),
	}
}
