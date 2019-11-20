package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetCourseInformationHandler get course information file
func GetCourseInformationHandler(ctx iris.Context) {
	type rule struct {
		CourseID string `valid:"required"`
	}

	params := &rule{
		CourseID: ctx.Params().Get("courseID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.GetInformation(params.CourseID)

	if err != (*errors.Error)(nil) {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	file(ctx, result["Path"], result["Name"])
	return
}
