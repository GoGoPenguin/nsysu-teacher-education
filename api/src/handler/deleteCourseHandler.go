package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// DeleteCourseHandler delete course
func DeleteCourseHandler(ctx iris.Context) {
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

	result, err := service.DeleteCourse(params.CourseID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
