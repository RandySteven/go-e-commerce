package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/RandySteven/go-e-commerce.git/enums/content_type"
	userrole "github.com/RandySteven/go-e-commerce.git/enums/user_role"
	"github.com/RandySteven/go-e-commerce.git/pkg/auth"
	"github.com/RandySteven/go-e-commerce.git/utils"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.ContentType(w, content_type.ApplicationJson)

		authToken := r.Header.Get("Authorization")
		log.Println(authToken)
		tokenStr := authToken[len("Bearer "):]
		claims := &auth.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return auth.JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			return
		}

		ctx := context.WithValue(r.Context(), "id", claims.UserID)
		ctx2 := context.WithValue(ctx, "email", claims.Email)
		if claims.UserID != nil {
			claims.RoleID = userrole.UserRole
		} else if claims.ShopID != nil {
			claims.RoleID = userrole.ShopRole
		}
		ctx3 := context.WithValue(ctx2, "role_id", claims.RoleID)
		log.Println(ctx3)
		next.ServeHTTP(w, r.WithContext(ctx3))
	})
}

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		whiteListedEndpoint := make(map[uint][]string)
		whiteListedEndpoint[userrole.UserRole] = []string{
			"/users",
		}
		whiteListedEndpoint[userrole.ShopRole] = []string{
			"/shops",
			"/products",
		}
		utils.ContentType(w, content_type.ApplicationJson)

		roleId, ok := r.Context().Value("role_id").(uint)

		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		requestedEndpoint := r.URL.Path
		if !contains(whiteListedEndpoint[roleId], requestedEndpoint) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
