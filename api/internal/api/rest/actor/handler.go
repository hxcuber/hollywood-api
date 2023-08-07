package actor

import "github.com/hxcuber/hollywood-api/api/internal/controller/actors"

type Handler struct {
	actorCtrl actors.Controller
}

func New(actorCtrl actors.Controller) Handler {
	return Handler{actorCtrl: actorCtrl}
}
