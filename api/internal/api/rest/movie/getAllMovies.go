package movieHandler

import (
	"github.com/hxcuber/hollywood-api/api/pkg/httpserv"
	"net/http"
)

func (h Handler) GetAllMovies() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		// validate headers, in this case dont have to

		result, err := h.movieCtrl.GetAllMovies(r.Context())
		if err != nil {
			return err
		}

		httpserv.RespondJSON(r.Context(), w, result)
		return nil
	})
}
