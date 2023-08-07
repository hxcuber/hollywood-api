package actor

import (
	"errors"
	"net/http"
)

func (a *Request) Bind(r *http.Request) error {
	if a.Actor == nil {
		return errors.New("cannot parse fields, or fields are empty")
	}
	return nil
}
