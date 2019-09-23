package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/error"
	"github.com/nsysu/teacher-education/src/service"
)

// UpdateCourseReviewHandler update course review
func UpdateCourseReviewHandler(ctx iris.Context) {
	type rule struct {
		StudentCourseID string `valid:"required"`
		Review          string `valid:"required, length(0|150)"`
	}

	params := &rule{
		StudentCourseID: ctx.FormValue("StudentCourseID"),
		Review:          ctx.FormValue("Review"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, error.ValidateError(err.Error()))
		return
	}

	result, err := service.UpdateCourseReview(params.StudentCourseID, params.Review)

	if err != (*error.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
