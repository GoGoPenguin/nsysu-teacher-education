package middleware

import (
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/utils/token"
)

// AuthMiddleware validate token before handler
func AuthMiddleware(ctx iris.Context) {
	tokenString, err := jwtmiddleware.FromAuthHeader(ctx)
	if err != nil {
		return
	}

	if err := token.Validate(tokenString); err != nil {
		return
	}

	ctx.Next()
}
