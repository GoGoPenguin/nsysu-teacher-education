package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// GetStudentCourseHandler get student-course list
func GetStudentCourseHandler(ctx iris.Context) {
	type rule struct {
		Start  string `valid:"required"`
		Length string `valid:"required"`
		Draw   string `valid:"-"`
		Search string `valid:"-"`
	}

	params := &rule{
		Start:  ctx.URLParamDefault("start", "0"),
		Length: ctx.URLParamDefault("length", "30"),
		Draw:   ctx.URLParam("draw"),
		Search: ctx.URLParam("search[value]"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		json(ctx, map[string]interface{}{
			"error": errors.ValidateError(err.Error()).Error(),
		})
		return
	}

	operator := auth.Account(ctx)
	result, err := service.GetStudentCourseList(operator, params.Start, params.Length, params.Search)

	if err != (*errors.Error)(nil) {
		json(ctx, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	result["draw"] = params.Draw
	json(ctx, result)
	return
}
