package handler

import (
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
)

// CreateStudentsHandler create students
func CreateStudentsHandler(ctx iris.Context) {
	file, _, err := ctx.FormFile("CSV")

	if err != nil {
		failed(ctx, error.ValidateError("CSV: non zero value required"))
		return
	}
	defer file.Close()

	result, err := service.CreateStudents(file)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
