package validation

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xinyi-chong/common-lib/consts"
	apperrors "github.com/xinyi-chong/common-lib/errors"
	"reflect"
	"strings"
	"sync"
)

var (
	validatorInstance *validator.Validate
	once              sync.Once
)

func getValidator() *validator.Validate {
	once.Do(func() {
		validatorInstance = validator.New()

		// Use json tag instead of struct field name
		validatorInstance.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	})

	return validatorInstance
}

// GinBindAndValidate binds JSON request body to a struct and validates it
func GinBindAndValidate[T any](c *gin.Context) (T, error) {
	var param T
	if err := c.ShouldBindJSON(&param); err != nil {
		return param, apperrors.ErrBadRequest.Wrap(err)
	}
	if err := ValidateStruct(&param); err != nil {
		return param, err
	}
	return param, nil
}

// ValidateStruct validates the struct against validation rules
func ValidateStruct(obj interface{}) error {
	v := getValidator()

	if err := v.Struct(obj); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			return convertValidationError(validationErrors[0])
		}
		return err
	}

	return nil
}

// convertValidationError converts a validator FieldError to an application error
func convertValidationError(fe validator.FieldError) error {
	fieldName := consts.Field(fe.Field())

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
