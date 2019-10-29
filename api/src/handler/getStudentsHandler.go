package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetStudentsHandler get student list
func GetStudentsHandler(ctx iris.Context) {
	type rule struct {
		Start  string `valid:"required"`
		Length string `valid:"required"`
		Draw   string `valid:"required"`
	}

	params := &rule{
		Start:  ctx.URLParamDefault("start", "0"),
		Length: ctx.URLParamDefault("length", "30"),
		Draw:   ctx.URLParam("draw"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		json(ctx, map[string]interface{}{
			"error": errors.ValidateError(err.Error()).Error(),
		})
		return
	}

	result, err := service.GetStudents(params.Start, params.Length)

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
