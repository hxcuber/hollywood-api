package actorResponse

import "github.com/hxcuber/hollywood-api/api/internal/models"

type ActorResponse struct {
	*models.Actor
}

func New(a *models.Actor) *ActorResponse {
	return &ActorResponse{a}
}
