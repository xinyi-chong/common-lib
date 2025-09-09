package apperrors

import (
	"errors"
	"github.com/xinyi-chong/common-lib/consts"
	locale "github.com/xinyi-chong/common-lib/i18n"
)

type Error struct {
	MessageKey   string              // i18n key
	HTTPStatus   int                 // HTTP status code
	Err          error               // Wrapped error, if any
	TemplateData locale.TemplateData // For dynamic i18n translation
	Op           string              // Operation name for logging or debugging
}

func New(msgKey string, status int) *Error {
	return &Error{MessageKey: msgKey, HTTPStatus: status}
}

func NewWithDefaultField(msgKey string, status int) *Error {
	return New(msgKey, status).WithField(consts.DefaultField).WithValue("0")
}

func Is(err error, target *Error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.MessageKey == target.MessageKey
	}
	return false
}

func (e *Error) Error() string {
	var opPrefix string
	if e.Op != "" {
		opPrefix = "[" + e.Op + "] "
	}

	errStr := e.MessageKey

	if e.Err != nil {
		var appErr *Error
		if !errors.As(e.Err, &appErr) {
			errStr = e.Err.Error()
		}
	}

	return opPrefix + errStr
}

func (e *Error) WithOp(op string) *Error {
	e.Op = op
	return e
}

func (e *Error) Wrap(err error) *Error {
	e.Err = err
	return e
}

func (e *Error) WithField(field consts.Field) *Error {
	if e.TemplateData == nil {
		e.TemplateData = locale.TemplateData{}
	}
	e.TemplateData["Field"] = field
	return e
}

func (e *Error) WithValue(value string) *Error {
	if e.TemplateData == nil {
		e.TemplateData = locale.TemplateData{}
	}
	e.TemplateData["Value"] = value
	return e
}

func (e *Error) WithTemplateData(data locale.TemplateData) *Error {
	if e.TemplateData == nil {
		e.TemplateData = locale.TemplateData{}
	}
	for k, v := range data {
		e.TemplateData[k] = v
	}
	return e
}
