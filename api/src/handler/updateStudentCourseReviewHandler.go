package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
)

// UpdateStudentCourseReviewHandler update course review
func UpdateStudentCourseReviewHandler(ctx iris.Context) {
	type rule struct {
		StudentCourseID string `valid:"required"`
		Review          string `valid:"required, length(0|300)"`
	}

	params := &rule{
		StudentCourseID: ctx.FormValue("StudentCourseID"),
		Review:          ctx.FormValue("Review"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateStudentCourseReview(params.StudentCourseID, params.Review)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
