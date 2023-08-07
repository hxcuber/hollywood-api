package movieHandler

import "github.com/hxcuber/hollywood-api/api/internal/controller/movieController"

type Handler struct {
	movieCtrl movieController.Controller
}

func New(movieCtrl movieController.Controller) Handler {
	return Handler{movieCtrl: movieCtrl}
}
