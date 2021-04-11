package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// SignUpServiceLearningHandler sign up service learning
func SignUpServiceLearningHandler(ctx iris.Context) {
	type rule struct {
		ServiceLearningID string `valid:"required"`
	}

	params := &rule{
		ServiceLearningID: ctx.FormValue("ServiceLearningID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	account := auth.Account(ctx)
	result, err := service.SingUpServiceLearning(nil, account, params.ServiceLearningID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
