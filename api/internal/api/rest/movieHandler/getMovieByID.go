package movieHandler

import (
	"github.com/go-chi/chi/v5"
	"github.com/hxcuber/hollywood-api/api/pkg/httpserv"
	"net/http"
	"strconv"
)

func (h Handler) GetMovieById() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		idString := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			return err
		}

		ctx := r.Context()
		movie, err := h.movieCtrl.GetMovieByID(ctx, int64(id))
		if err != nil {
			return err
		}

		httpserv.RespondJSON(ctx, w, movie)

		return nil
	})
}
