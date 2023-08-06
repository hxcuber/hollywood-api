package actorHandler

import "github.com/hxcuber/hollywood-api/api/internal/controller/actorController"

type Handler struct {
	actorCtrl actorController.Controller
}

func New(actorCtrl actorController.Controller) Handler {
	return Handler{actorCtrl: actorCtrl}
}
