package routers

import (
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/middlewares"
	"github.com/gorilla/mux"
)

func (h *Handlers) InitRouter(r *mux.Router) {

	userAuth := r.PathPrefix("/users").Subrouter()
	userAuth.HandleFunc("/register", h.UserHandler.RegisterUser).Methods(http.MethodPost)
	userAuth.HandleFunc("/login", h.UserHandler.LoginUser).Methods(http.MethodPost)

	userGroup := r.PathPrefix("/users").Subrouter()
	userGroup.Use(middlewares.AuthenticationMiddleware, middlewares.AuthorizationMiddleware)
	userGroup.HandleFunc("", h.UserHandler.TestGetUser).Methods(http.MethodGet)

	shopAuth := r.PathPrefix("/shops").Subrouter()
	shopAuth.Use(middlewares.TimeMiddleware)
	shopAuth.HandleFunc("/register", h.ShopHandler.RegisterShop).Methods(http.MethodPost)
	shopAuth.HandleFunc("/login", h.ShopHandler.LoginShop).Methods(http.MethodPost)

	shopGroup := r.PathPrefix("/shops").Subrouter()
	shopGroup.Use(middlewares.AuthenticationMiddleware)
	shopGroup.HandleFunc("", h.ShopHandler.ShopDetail).Methods(http.MethodGet)

	productGroup := r.PathPrefix("/products").Subrouter()
	productGroup.Use(middlewares.AuthenticationMiddleware, middlewares.AuthorizationMiddleware)
	productGroup.HandleFunc("", h.ProductHandler.GetAllProducts).Methods(http.MethodGet)
	productGroup.HandleFunc("", h.ProductHandler.CreateProduct).Methods(http.MethodPost)
}
