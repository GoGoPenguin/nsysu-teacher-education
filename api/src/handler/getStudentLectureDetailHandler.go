package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetStudentLectureDetailHandler get student-lecture detail
func GetStudentLectureDetailHandler(ctx iris.Context) {
	type rule struct {
		StudentLectureID string `valid:"required"`
	}

	params := &rule{
		StudentLectureID: ctx.Params().Get("studentLectureID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.GetStudentLectureDetail(params.StudentLectureID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
