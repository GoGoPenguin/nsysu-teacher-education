package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// UpdateServiceLearningStatusHandler update student-service-learning status
func UpdateServiceLearningStatusHandler(ctx iris.Context) {
	type rule struct {
		StudentServiceLearningID string `valid:"required"`
		Status                   string `valid:"required, in(pass|failed)"`
		Comment                  string `valid:"required, length(0|150)"`
	}

	params := &rule{
		StudentServiceLearningID: ctx.FormValue("StudentServiceLearningID"),
		Status:                   ctx.FormValue("Status"),
		Comment:                  ctx.FormValue("Comment"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateStudentServiceLearningStatus(params.StudentServiceLearningID, params.Status, params.Comment)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
