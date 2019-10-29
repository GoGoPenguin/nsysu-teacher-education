package handler

import (
	"github.com/kataras/iris"
)

// HelloHandler return hello message
func HelloHandler(ctx iris.Context) {
	success(ctx, "hello")
}
