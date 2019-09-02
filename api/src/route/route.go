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

	v1 := app.Party("/v1", hero.Handler(middleware.AuthMiddleware))
	{
		v1.Post("/login", hero.Handler(handler.LoginHandler)) // 登入

		users := v1.Party("/users")
		{
			users.Get("/", hero.Handler(handler.GetStudentsHandler))
			users.Post("/", hero.Handler(handler.CreateStudentsHandler))
		}
	}

	app.Run(iris.Addr(addr))
}
