package middlewares

import (
	"errors"
	"log"
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/apperror"
	"github.com/RandySteven/go-e-commerce.git/entity/payload/responses"
	"github.com/RandySteven/go-e-commerce.git/utils"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		var (
			errAgeValidation            *apperror.ErrAgeValidation
			errEmailAndPasswordNotMatch *apperror.ErrEmailAndPasswordNotMatch
			errBadRequest               *apperror.ErrBadRequest
		)
		resp := responses.Response{}
		err := r.Context().Value("error")
		log.Println(err)
		if err != nil {
			err2 := err.(error)
			resp.Error = err2.Error()
			switch {
			case errors.Is(err2, errAgeValidation):
				utils.ResponseHandler(w, http.StatusBadRequest, resp)
			case errors.Is(err2, errEmailAndPasswordNotMatch):
				utils.ResponseHandler(w, http.StatusBadRequest, resp)
			case errors.Is(err2, errBadRequest):
				utils.ResponseHandler(w, http.StatusBadRequest, resp)
			default:
				utils.ResponseHandler(w, http.StatusInternalServerError, resp)
			}
		}
	})
}
