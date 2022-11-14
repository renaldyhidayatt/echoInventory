package pkg

import (
	"echoinventory/schemas"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func Sign(schemajwt *schemas.SchemaJWT) (string, error) {

	expTimeMs, _ := strconv.Atoi("86400000")
	exp := time.Now().Add(time.Millisecond * time.Duration(expTimeMs)).Unix()

	claims := &schemas.JWtMetaRequest{
		ID:    schemajwt.ID,
		Email: schemajwt.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwt, err := token.SignedString([]byte("secret"))

	return jwt, err
}
