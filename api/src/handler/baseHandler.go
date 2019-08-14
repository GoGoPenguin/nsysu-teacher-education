package handler

import (
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
)

func success(ctx iris.Context, data interface{}) {
	ctx.JSON(iris.Map{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func failed(ctx iris.Context, err *error.Error) {
	ctx.JSON(iris.Map{
		"code":    err.Code(),
		"message": err.Error(),
		"data":    []string{},
	})
}
