package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
)

// LoginHandler user login
func LoginHandler(ctx iris.Context) {
	type rule struct {
		Account  string `valid:"required"`
		Password string `valid:"required"`
		Role     string `valid:"required, in(student|admin)"`
	}

	params := &rule{
		Account:  ctx.FormValue("Account"),
		Password: ctx.FormValue("Password"),
		Role:     ctx.FormValue("Role"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.Login(params.Account, params.Password, params.Role)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
