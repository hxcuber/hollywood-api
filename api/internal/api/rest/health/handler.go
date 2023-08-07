package health

import "github.com/hxcuber/hollywood-api/api/internal/controller/systems"

// Handler is the web handler for this pkg
type Handler struct {
	systemCtrl systems.Controller
}

// New instantiates a new Handler and returns it
func New(systemCtrl systems.Controller) Handler {
	return Handler{systemCtrl: systemCtrl}
}
