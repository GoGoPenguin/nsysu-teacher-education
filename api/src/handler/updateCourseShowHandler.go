package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// UpdateCourseShowHandler update course's state of show
func UpdateCourseShowHandler(ctx iris.Context) {
	type rule struct {
		CourseID string `valid:"required"`
		Show     string `valid:"required, boolean"`
	}

	params := &rule{
		CourseID: ctx.Params().Get("courseID"),
		Show:     ctx.FormValue("Show"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateStateOfShow(params.CourseID, typecast.StringToBool(params.Show))

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
