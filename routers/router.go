package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handlers) InitRouter(r *mux.Router) {

	userGroup := r.PathPrefix("/users").Subrouter()
	userGroup.HandleFunc("/register", h.UserHandler.RegisterUser).Methods(http.MethodPost)
	userGroup.HandleFunc("/login", h.UserHandler.LoginUser).Methods(http.MethodPost)
	// userGroup.Use(middlewares.AuthenticationMiddleware)
	// userGroup.Use(middlewares.AuthorizationMiddleware)
	userGroup.HandleFunc("", h.UserHandler.LoginUser).Methods(http.MethodGet)

}
