package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// UpdateCourseStatusHandler update student course status
func UpdateCourseStatusHandler(ctx iris.Context) {
	type rule struct {
		StudentCourseID string `valid:"required"`
		Status          string `valid:"required, in(pass|failed)"`
	}

	params := &rule{
		StudentCourseID: ctx.FormValue("StudentCourseID"),
		Status:          ctx.FormValue("Status"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateCourseStatus(params.StudentCourseID, params.Status)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
