package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetCourseInformationHandler get list of course
func GetCourseInformationHandler(ctx iris.Context) {
	type rule struct {
		FileName string `valid:"required"`
	}

	params := &rule{
		FileName: ctx.Params().Get("filename"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.GetInformation(params.FileName)

	if err != (*errors.Error)(nil) {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	file(ctx, result, params.FileName)
	return
}
