package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// SignUpCourseHandler sign up course
func SignUpCourseHandler(ctx iris.Context) {
	type rule struct {
		CourseID string `valid:"required"`
		Meal     string `valid:"required, in(meat|vegetable)"`
	}

	params := &rule{
		CourseID: ctx.FormValue("CourseID"),
		Meal:     ctx.FormValue("Meal"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	account := auth.Account(ctx)
	result, err := service.SingUpCourse(account, params.CourseID, params.Meal)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
