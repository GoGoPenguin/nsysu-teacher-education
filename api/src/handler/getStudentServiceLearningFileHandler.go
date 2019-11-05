package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetStudentServiceLearningFileHandler get student-service-learning file
func GetStudentServiceLearningFileHandler(ctx iris.Context) {
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

	result, err := service.GetStudentServiceLearningFile(params.StudentServiceLearningID, params.File)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	file(ctx, result["Path"], result["Name"])
	return
}
