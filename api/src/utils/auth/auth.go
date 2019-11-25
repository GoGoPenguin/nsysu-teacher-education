package auth

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
)

// Account get account field of token
func Account(ctx iris.Context) string {
	tokenString, _ := jwtmiddleware.FromAuthHeader(ctx)
	token, _ := jwt.Parse(tokenString, nil)
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims["account"].(string)
}
