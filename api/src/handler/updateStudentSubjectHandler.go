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
		Name            string `valid:"-"`
		Year            string `valid:"-"`
		Semester        string `valid:"-"`
		Credit          string `valid:"-"`
		Score           string `valid:"-"`
	}

	params := &rule{
		StudentLetureID: ctx.FormValue("StudentLetureID"),
		SubjectID:       ctx.FormValue("SubjectID"),
		Name:            ctx.FormValue("Name"),
		Year:            ctx.FormValue("Year"),
		Semester:        ctx.FormValue("Semester"),
		Credit:          ctx.FormValue("Credit"),
		Score:           ctx.FormValue("Score"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	account := auth.Account(ctx)
	result, err := service.UpdateStudentSubject(
		account,
		params.StudentLetureID,
		params.SubjectID,
		params.Name,
		params.Year,
		params.Semester,
		params.Credit,
		params.Score,
	)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
