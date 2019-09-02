package middleware

import (
	"regexp"

	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/utils/token"
)

var whiteList = []string{
	`/v1/login`,
	`/v1/renew-token`,
}

// AuthMiddleware validate token before handler
func AuthMiddleware(ctx iris.Context) {
	for _, path := range whiteList {
		pathRegexp := regexp.MustCompile(path)
		if pathRegexp.MatchString(ctx.Path()) {
			ctx.Next()
			return
		}
	}

	tokenString, err := jwtmiddleware.FromAuthHeader(ctx)
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	if err := token.Validate(tokenString); err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		return
	}

	ctx.Next()
}
