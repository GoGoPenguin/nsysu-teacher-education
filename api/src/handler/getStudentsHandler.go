package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

// GetStudentsHandler get student list
func GetStudentsHandler(ctx iris.Context) {
	type rule struct {
		Index string `valid:"required"`
		Count string `valid:"required"`
	}

	params := &rule{
		Index: ctx.URLParamDefault("Index", "0"),
		Count: ctx.URLParamDefault("Count", "30"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.GetStudents(params.Index, params.Count)
	logger.Debug(result)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
