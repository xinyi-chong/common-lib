package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	apperrors "github.com/xinyi-chong/common-lib/errors"
	locale "github.com/xinyi-chong/common-lib/i18n"
	"github.com/xinyi-chong/common-lib/success"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, success *success.Success, data interface{}) {
	message := locale.Translate(c, locale.CategorySuccess, success.MessageKey, success.TemplateData)

	c.JSON(success.HTTPStatus, Response{
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, err error, overrideErrs ...*apperrors.Error) {
	var overrideErr *apperrors.Error
	if len(overrideErrs) > 0 {
		overrideErr = overrideErrs[0]
	}

	appErr := resolveAppError(err, overrideErr)
	message := locale.Translate(c, locale.CategoryError, appErr.MessageKey, appErr.TemplateData)

	c.AbortWithStatusJSON(appErr.HTTPStatus, Response{Message: message})
}

func resolveAppError(err error, overrideErr *apperrors.Error) *apperrors.Error {
	var appErr *apperrors.Error

	if overrideErr != nil && (!errors.As(err, &appErr) || errors.Is(err, apperrors.ErrInternalServerError)) {
		return overrideErr.Wrap(err)
	}

	if errors.As(err, &appErr) {
		return appErr
	}

	return apperrors.ErrInternalServerError.Wrap(err)
}
