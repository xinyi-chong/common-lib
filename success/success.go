package success

import (
	locale "github.com/xinyi-chong/common-lib/i18n"
)

type Success struct {
	MessageKey   string
	HTTPStatus   int
	TemplateData locale.TemplateData
}

func New(msgKey string, httpStatus int) *Success {
	return &Success{MessageKey: msgKey, HTTPStatus: httpStatus}
}
