package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) InitRouter(r *mux.Router) {

	r.HandleFunc("/register", h.UserHandler.RegisterUser).Methods(http.MethodPost)

}
