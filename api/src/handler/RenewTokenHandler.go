package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// RenewTokenHandler get new access token
func RenewTokenHandler(ctx iris.Context) {
	type rule struct {
		Account      string `valid:"required"`
		RefreshToken string `valid:"required"`
	}

	params := &rule{
		Account:      ctx.FormValue("Account"),
		RefreshToken: ctx.FormValue("RefreshToken"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.RenewToken(params.Account, params.RefreshToken)
	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
