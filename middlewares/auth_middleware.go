package middlewares

import (
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/enums/content_type"
	"github.com/RandySteven/go-e-commerce.git/pkg/auth"
	"github.com/RandySteven/go-e-commerce.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.ContentType(w, content_type.ApplicationJson)

		authToken := w.Header().Get("Authorization")

		tokenStr := authToken[len("Bearer "):]
		claims := &auth.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return auth.JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			return
		}

	})
}
