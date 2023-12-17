package interfaces

import "net/http"

type (
	UserHandler interface {
		RegisterUser(res http.ResponseWriter, req *http.Request)
		LoginUser(res http.ResponseWriter, req *http.Request)
	}
)
