package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// DeleteServiceLearningHandler delete service-learning
func DeleteServiceLearningHandler(ctx iris.Context) {
	type rule struct {
		ServiceLearnginID string `valid:"required"`
	}

	params := &rule{
		ServiceLearnginID: ctx.Params().Get("serviceLearnginID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.DeleteServiceLearning(params.ServiceLearnginID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
