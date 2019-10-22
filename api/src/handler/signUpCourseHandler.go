package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// SignUpCourseHandler sign up course
func SignUpCourseHandler(ctx iris.Context) {
	type rule struct {
		Account  string `valid:"required"`
		CourseID string `valid:"required"`
		Meal     string `valid:"required, in(meat|vegetable)"`
	}

	params := &rule{
		Account:  ctx.FormValue("Account"),
		CourseID: ctx.FormValue("CourseID"),
		Meal:     ctx.FormValue("Meal"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.SingUpCourse(params.Account, params.CourseID, params.Meal)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
