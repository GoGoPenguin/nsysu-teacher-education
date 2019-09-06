package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
)

// CreateCourseHandler user login
func CreateCourseHandler(ctx iris.Context) {
	type rule struct {
		Topic       string `valid:"required"`
		Information string `valid:"required"`
		Type        string `valid:"required, in(A|B|C)"`
		Start       string `valid:"required"`
		End         string `valid:"required"`
	}

	params := &rule{
		Topic:       ctx.FormValue("Topic"),
		Information: ctx.FormValue("Information"),
		Type:        ctx.FormValue("Type"),
		Start:       ctx.FormValue("Start"),
		End:         ctx.FormValue("End"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.CreateCourse(
		params.Topic,
		params.Information,
		params.Type,
		params.Start,
		params.End,
	)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
