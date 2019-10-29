package handler

import (
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
)

func success(ctx iris.Context, data interface{}) {
	ctx.JSON(iris.Map{
		"code":    0,
		"message": "success",
		"data":    data,
	})
}

func failed(ctx iris.Context, err interface{}) {
	ctx.JSON(iris.Map{
		"code":    err.(*errors.Error).Code(),
		"message": err.(*errors.Error).Error(),
		"data":    []string{},
	})
}

func json(ctx iris.Context, data map[string]interface{}) {
	ctx.JSON(data)
}

func file(ctx iris.Context, filename, destinationName string) {
	ctx.SendFile(filename, destinationName)
}
