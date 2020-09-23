package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// SignUpLectureHandler sign up lecture
func SignUpLectureHandler(ctx iris.Context) {
	type rule struct {
		LectureID string `valid:"required"`
	}

	params := &rule{
		LectureID: ctx.FormValue("LectureID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	account := auth.Account(ctx)
	result, err := service.SingUpLecture(account, params.LectureID)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
