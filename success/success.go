package success

import (
	"github.com/xinyi-chong/common-lib/consts"
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

func NewWithDefaultField(msgKey string, status int) *Success {
	return New(msgKey, status).WithField(consts.DefaultField).WithValue("0")
}

func (s *Success) WithField(field consts.Field) *Success {
	if s.TemplateData == nil {
		s.TemplateData = locale.TemplateData{}
	}
	s.TemplateData["Field"] = field
	return s
}

func (s *Success) WithValue(value string) *Success {
	if s.TemplateData == nil {
		s.TemplateData = locale.TemplateData{}
	}
	s.TemplateData["Value"] = value
	return s
}
