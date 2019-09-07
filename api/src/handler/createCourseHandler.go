package handler

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
	t "github.com/nsysu/teacher-education/src/utils/time"
)

// CreateCourseHandler user login
func CreateCourseHandler(ctx iris.Context) {
	type rule struct {
		Topic       string    `valid:"required"`
		Information string    `valid:"required"`
		Type        string    `valid:"required, in(A|B|C)"`
		Start       time.Time `valid:"required"`
		End         time.Time `valid:"required"`
	}

	startTime, err := time.Parse(t.DateTime, ctx.FormValue("Start"))
	if err != nil {
		failed(ctx, error.ValidateError("Start: "+ctx.FormValue("Start")+" does not validate as timestamp"))
		return
	}
	endTime, err := time.Parse(t.DateTime, ctx.FormValue("End"))
	if err != nil {
		failed(ctx, error.ValidateError("End: "+ctx.FormValue("Start")+" does not validate as timestamp"))
		return
	}
	if !startTime.Before(endTime) {
		failed(ctx, error.ValidateError("Start: "+ctx.FormValue("Start")+" does not after "+ctx.FormValue("End")))
		return
	}

	params := &rule{
		Topic:       ctx.FormValue("Topic"),
		Information: ctx.FormValue("Information"),
		Type:        ctx.FormValue("Type"),
		Start:       startTime,
		End:         endTime,
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
