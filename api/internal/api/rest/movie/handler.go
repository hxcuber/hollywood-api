package movie

import "github.com/hxcuber/hollywood-api/api/internal/controller/movies"

type Handler struct {
	movieCtrl movies.Controller
}

func New(movieCtrl movies.Controller) Handler {
	return Handler{movieCtrl: movieCtrl}
}
