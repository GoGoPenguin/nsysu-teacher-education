package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// UpdateStudentLetureStatusHandler update studetn leture status
func UpdateStudentLetureStatusHandler(ctx iris.Context) {
	type rule struct {
		LetureID string `valid:"required"`
		Pass     string `valid:"required, in(true|false)"`
	}

	params := &rule{
		LetureID: ctx.FormValue("LetureID"),
		Pass:     ctx.FormValue("Pass"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	account := auth.Account(ctx)
	result, err := service.UpdateStudentLetuerPass(account, params.LetureID, params.Pass)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
