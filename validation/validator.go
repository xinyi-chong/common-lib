package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
	apperrors "github.com/xinyi-chong/common-lib/errors"
	"reflect"
	"strings"
)

// Check validates the struct against `validator` tags.
func Check(obj interface{}) error {
	v := validator.New()

	// Use json tag instead of struct field name.
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := v.Struct(obj); err != nil {
		var verrs validator.ValidationErrors
		if errors.As(err, &verrs) {
			return getErrorMessage(verrs[0])
		}
		return err
	}

	return nil
}

// getErrorMessage converts a validator error to an app error.
func getErrorMessage(fe validator.FieldError) error {
	fieldName := fe.Field()

	switch fe.Tag() {
	case "required":
		return apperrors.ErrXIsRequired.WithField(fieldName)
	case "email":
		return apperrors.ErrInvalidX.WithField(fieldName)
	case "min":
		return apperrors.ErrXMin.WithField(fieldName).WithValue(fe.Param())
	case "max":
		return apperrors.ErrXMax.WithField(fieldName).WithValue(fe.Param())
	default:
		return apperrors.ErrInvalidX.WithField(fieldName)
	}
}
