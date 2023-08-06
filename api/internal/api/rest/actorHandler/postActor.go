package actorHandler

import (
	"github.com/go-chi/render"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/requestModel/actorRequest"
	"github.com/hxcuber/hollywood-api/api/internal/api/rest/responseModel/actorResponse"
	"github.com/hxcuber/hollywood-api/api/pkg/httpserv"
	"net/http"
)

func (h Handler) PostActor() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		// validate headers, dont need in this case
		var data actorRequest.ActorRequest
		if err := render.Bind(r, &data); err != nil {
			return err
		}

		newActorPointer := data.Actor
		newActor, err := h.actorCtrl.CreateActor(r.Context(), *newActorPointer)
		if err != nil {
			return err
		}

		render.Render(w, r, actorResponse.New(&newActor))
		return nil
	})
}
