package schemas

import (
	"github.com/golang-jwt/jwt"
)

type JWtMetaRequest struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
