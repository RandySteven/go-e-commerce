package auth

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("secret")

type JWTClaim struct {
	UserID *uint
	ShopID *uint
	Email  string
	RoleID uint
	jwt.RegisteredClaims
}
