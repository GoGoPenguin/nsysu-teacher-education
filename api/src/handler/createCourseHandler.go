package handler

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
	t "github.com/nsysu/teacher-education/src/utils/time"
)

// CreateCourseHandler create a course
func CreateCourseHandler(ctx iris.Context) {
	type rule struct {
		Topic string    `valid:"required"`
		Type  string    `valid:"required, in(A|B|C)"`
		Start time.Time `valid:"required"`
		End   time.Time `valid:"required"`
	}

	file, header, err := ctx.FormFile("Information")

	if err != nil {
		failed(ctx, error.ValidateError("Information: non zero value required"))
		return
	}
	defer file.Close()

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
		Topic: ctx.FormValue("Topic"),
		Type:  ctx.FormValue("Type"),
		Start: startTime,
		End:   endTime,
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.CreateCourse(
		params.Topic,
		params.Type,
		file,
		header,
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
