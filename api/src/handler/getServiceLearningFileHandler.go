package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// GetServiceLearningFileHandler get student-service-learning file
func GetServiceLearningFileHandler(ctx iris.Context) {
	type rule struct {
		StudentServiceLearningID string `valid:"required"`
		File                     string `valid:"required, in(reference|review)"`
	}

	params := &rule{
		StudentServiceLearningID: ctx.URLParamDefault("StudentServiceLearningID", ""),
		File:                     ctx.Params().Get("file"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	operator := auth.Account(ctx)
	result, err := service.GetStudentServiceLearningFile(operator, params.StudentServiceLearningID, params.File)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	file(ctx, result["Path"], result["Name"])
	return
}
