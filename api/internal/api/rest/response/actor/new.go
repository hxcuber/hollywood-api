package actor

import "github.com/hxcuber/hollywood-api/api/internal/models"

type Response struct {
	*models.Actor
}

func New(a *models.Actor) *Response {
	return &Response{a}
}
