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
		v1.Post("/login", hero.Handler(handler.LoginHandler))            // 登入
		v1.Post("/logout", hero.Handler(handler.LogoutHandler))          // 登出
		v1.Post("/renew-token", hero.Handler(handler.RenewTokenHandler)) // 取得新的 access token

		users := v1.Party("/users")
		{
			users.Get("/", hero.Handler(handler.GetStudentsHandler))     // 取得學生列表
			users.Post("/", hero.Handler(handler.CreateStudentsHandler)) // 新增學生帳號
		}

		course := v1.Party("/course")
		{
			course.Post("/", hero.Handler(handler.CreateCourseHandler))                  // 新增講座
			course.Get("/", hero.Handler(handler.GetCourseHandler))                      // 取得講座列表
			course.Get("/{filename}", hero.Handler(handler.GetCourseInformationHandler)) // 取得講座資訊
		}
	}

	app.Run(iris.Addr(addr))
}
