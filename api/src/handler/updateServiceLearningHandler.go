package handler

import (
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris"
	"github.com/nsysu/teacher-education/src/errors"
	"github.com/nsysu/teacher-education/src/service"
	"github.com/nsysu/teacher-education/src/utils/auth"
)

// UpdateServiceLearningHandler update course review
func UpdateServiceLearningHandler(ctx iris.Context) {

	type rule struct {
		StudentServiceLearningID string `valid:"required"`
	}

	params := &rule{
		StudentServiceLearningID: ctx.FormValue("StudentServiceLearningID"),
	}

	if _, err := govalidator.ValidateStruct(params); err != nil {
		failed(ctx, errors.ValidateError(err.Error()))
		return
	}

	var referenceFileName, reviewFileName string

	reference, referenceHeader, err := ctx.FormFile("Reference")
	if err == nil {
		referenceFileName = referenceHeader.Filename
		defer reference.Close()
	}

	review, reviewHeader, err := ctx.FormFile("Review")
	if err == nil {
		reviewFileName = reviewHeader.Filename
		defer review.Close()
	}

	if reference == nil && review == nil {
		failed(ctx, errors.ValidateError("Upload Reference or Review"))
		return
	}

	operator := auth.Account(ctx)
	result, err := service.UpdateServiceLearning(
		reference,
		review,
		operator,
		params.StudentServiceLearningID,
		referenceFileName,
		reviewFileName,
	)

	if err != (*errors.Error)(nil) {
		failed(ctx, err)
		return
	}

	success(ctx, result)
	return
}
