package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
)

// SignUpServiceLearningHandler sign up service learning
func SignUpServiceLearningHandler(ctx iris.Context) {
	type rule struct {
		Account           string `valid:"required"`
		ServiceLearningID string `valid:"required"`
	}

	params := &rule{
		Account:           ctx.FormValue("Account"),
		ServiceLearningID: ctx.FormValue("ServiceLearningID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.SingUpServiceLearning(params.Account, params.ServiceLearningID)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
