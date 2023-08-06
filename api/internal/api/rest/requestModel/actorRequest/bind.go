package actorRequest

import (
	"errors"
	"net/http"
)

func (a *ActorRequest) Bind(r *http.Request) error {
	if a.Actor == nil {
		return errors.New("cannot parse fields, or fields are empty")
	}
	return nil
}
