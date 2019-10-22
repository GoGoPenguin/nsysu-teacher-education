package handler

import (
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// CreateStudentsHandler create students
func CreateStudentsHandler(ctx iris.Context) {
	file, _, err := ctx.FormFile("CSV")

	if err != nil {
		json(ctx, map[string]interface{}{
			"error": "CSV: non zero value required",
		})
		return
	}
	defer file.Close()

	_, err = service.CreateStudents(file)

	if err != (*errors.Error)(nil) {
		json(ctx, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	json(ctx, map[string]interface{}{})
	return
}
