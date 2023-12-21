package interfaces

import "net/http"

type (
	UserHandler interface {
		RegisterUser(res http.ResponseWriter, req *http.Request)
		LoginUser(res http.ResponseWriter, req *http.Request)
	}

	ShopHandler interface {
		RegisterShop(res http.ResponseWriter, req *http.Request)
		LoginShop(res http.ResponseWriter, req *http.Request)
	}
)
