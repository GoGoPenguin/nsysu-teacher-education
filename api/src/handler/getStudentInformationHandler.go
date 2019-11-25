package handler

import (
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// GetStudentInformationHandler get student information
func GetStudentInformationHandler(ctx iris.Context) {
	account := auth.Account(ctx)
	result, err := service.GetStudentInformation(account)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
