package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/typecast"
)

// UpdateStudentServiceLearningStatusHandler update student-service-learning status
func UpdateStudentServiceLearningStatusHandler(ctx iris.Context) {
	type rule struct {
		StudentServiceLearningID string `valid:"required"`
		Status                   string `valid:"required, in(pass|failed)"`
		Hours                    uint   `valid:"required, int"`
		Comment                  string `valid:"length(0|300)"`
	}

	params := &rule{
		StudentServiceLearningID: ctx.FormValue("StudentServiceLearningID"),
		Status:                   ctx.FormValue("Status"),
		Hours:                    typecast.StringToUint(ctx.FormValue("Hours")),
		Comment:                  ctx.FormValue("Comment"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateStudentServiceLearningStatus(
		params.StudentServiceLearningID,
		params.Status,
		params.Comment,
		params.Hours,
	)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
