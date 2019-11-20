package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// UpdateStudentSubjectHandler update student subject
func UpdateStudentSubjectHandler(ctx iris.Context) {
	type rule struct {
		StudentLetureID string `valid:"required"`
		SubjectID       string `valid:"required"`
		Score           string `valid:"range(0|100), int"`
		Pass            string `valid:"required, in(true|false)"`
	}

	params := &rule{
		StudentLetureID: ctx.FormValue("StudentLetureID"),
		SubjectID:       ctx.FormValue("SubjectID"),
		Score:           ctx.FormValue("Score"),
		Pass:            ctx.FormValue("Pass"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	account := auth.Account(ctx)
	result, err := service.UpdateStudentSubject(account, params.StudentLetureID, params.SubjectID, params.Score, params.Pass)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
