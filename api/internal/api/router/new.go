package router

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/actorHandler"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/healthHandler"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/movieHandler"
	"github.com/hxcuber/hollywood-api/api/internal/controller/actorController"
	"github.com/hxcuber/hollywood-api/api/internal/controller/movieController"
	"github.com/hxcuber/hollywood-api/api/internal/controller/systemController"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	corsOrigin []string,
	isGQLIntrospectionOn bool,
	systemCtrl systemController.Controller,
	actorCtrl actorController.Controller,
	movieCtrl movieController.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          corsOrigin,
		isGQLIntrospectionOn: isGQLIntrospectionOn,
		healthRESTHandler:    healthHandler.New(systemCtrl),
		actorRESTHandler:     actorHandler.New(actorCtrl),
		movieRESTHandler:     movieHandler.New(movieCtrl),
	}
}
