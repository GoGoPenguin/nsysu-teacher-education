package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// UpdateServiceLearningShowHandler update service-learning's state of show
func UpdateServiceLearningShowHandler(ctx iris.Context) {
	type rule struct {
		ServiceLearningID string `valid:"required"`
		Show              string `valid:"required, boolean"`
	}

	params := &rule{
		ServiceLearningID: ctx.Params().Get("serviceLearnginID"),
		Show:              ctx.FormValue("Show"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateServiceLearningStateOfShow(params.ServiceLearningID, typecast.StringToBool(params.Show))

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
