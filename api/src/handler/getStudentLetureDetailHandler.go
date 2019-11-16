package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetStudentLetureDetailHandler get student-course list
func GetStudentLetureDetailHandler(ctx iris.Context) {
	type rule struct {
		StudentLetureID string `valid:"required"`
	}

	params := &rule{
		StudentLetureID: ctx.Params().Get("studentLetureID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.GetStudentLetureDetail(params.StudentLetureID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
