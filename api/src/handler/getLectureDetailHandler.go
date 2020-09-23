package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// GetLectureDetailHandler get lecture detail
func GetLectureDetailHandler(ctx iris.Context) {
	type rule struct {
		LectureID string `valid:"required"`
	}

	params := &rule{
		LectureID: ctx.Params().Get("lectureID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.GetLectureDetail(params.LectureID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
