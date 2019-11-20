package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetLeturDetailHandler get leture detail
func GetLeturDetailHandler(ctx iris.Context) {
	type rule struct {
		LetureID string `valid:"required"`
	}

	params := &rule{
		LetureID: ctx.Params().Get("letureID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.GetLetureDetail(params.LetureID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
