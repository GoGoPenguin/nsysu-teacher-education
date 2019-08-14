package route

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/hero"
	"github.com/nsysu/teacher-education/src/handler"
	"github.com/nsysu/teacher-education/src/middleware"
	"github.com/nsysu/teacher-education/src/utils/config"
)

var addr = fmt.Sprintf("%v:%v", config.Get("server.host"), config.Get("server.port"))

// Run maps the routing path and keeps listening for request
func Run() {
	app := iris.New()

	// CORS
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CorsMiddleware)

	app.Get("/", hero.Handler(handler.HelloHandler))

	app.Run(iris.Addr(addr))
}
