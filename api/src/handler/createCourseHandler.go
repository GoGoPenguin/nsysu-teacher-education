package handler

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
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
		json(ctx, map[string]interface{}{
			"error": "Information: non zero value required",
		})
		return
	}
	defer file.Close()

	loc, _ := time.LoadLocation("Asia/Taipei")
	startTime, err := time.ParseInLocation(t.DateTime, ctx.FormValue("Start"), loc)
	if err != nil {
		json(ctx, map[string]interface{}{
			"error": "Start: " + ctx.FormValue("Start") + " does not validate as timestamp",
		})
		return
	}
	endTime, err := time.ParseInLocation(t.DateTime, ctx.FormValue("End"), loc)
	if err != nil {
		json(ctx, map[string]interface{}{
			"error": "End: " + ctx.FormValue("Start") + " does not validate as timestamp",
		})
		return
	}
	if !startTime.Before(endTime) {
		json(ctx, map[string]interface{}{
			"error": "Start: " + ctx.FormValue("Start") + " does not before " + ctx.FormValue("End"),
		})
		return
	}

	params := &rule{
		Topic: ctx.FormValue("Topic"),
		Type:  ctx.FormValue("Type"),
		Start: startTime,
		End:   endTime,
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		json(ctx, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	_, err = service.CreateCourse(
		params.Topic,
		params.Type,
		file,
		header,
		params.Start,
		params.End,
	)

	if err != (*errors.Error)(nil) {
		json(ctx, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json(ctx, map[string]interface{}{})
	return
}
