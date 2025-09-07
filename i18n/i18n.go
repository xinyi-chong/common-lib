package locale

import (
	"context"
	"embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xinyi-chong/common-lib/consts"
	"github.com/xinyi-chong/common-lib/logger"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"path"
	"sync"
)

type Category string
type TemplateData map[string]interface{}

const (
	CategorySuccess Category = "success"
	CategoryError   Category = "errors"
)

//go:embed *.toml
var fs embed.FS

var (
	bundle     *i18n.Bundle
	bundleOnce sync.Once
)

func Init() error {
	var initErr error
	bundleOnce.Do(func() {
		bundle = i18n.NewBundle(language.English)
		bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

		entries, err := fs.ReadDir(".")
		if err != nil {
			initErr = err
			return
		}

		for _, e := range entries {
			if path.Ext(e.Name()) == ".toml" {
				logger.Info("load locale file", zap.String("file", e.Name()))
				if _, err := bundle.LoadMessageFileFS(fs, e.Name()); err != nil {
					initErr = err
					return
				}
			}
		}
	})
	return initErr
}

func GetLocalizer(langs ...string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, langs...)
}

func Translate(c context.Context, category Category, messageID string, templateData TemplateData) string {
	localizer, ok := c.Value(consts.Localizer).(*i18n.Localizer)
	if !ok || localizer == nil {
		logger.Error("Translate: no localizer found in context")
		return messageID
	}

	for key, val := range templateData {
		if fieldName, ok := val.(string); ok {
			fieldID := "field." + fieldName
			translatedField, err := localizer.Localize(&i18n.LocalizeConfig{
				MessageID: fieldID,
			})
			if err != nil {
				logger.Warn("Translate: field translation missing", zap.String("fieldID", fieldID), zap.Error(err))
				translatedField = fieldName
			}
			templateData[key] = translatedField
		}
	}

	fullMessageID := fmt.Sprintf("%s.%s", category, messageID)
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    fullMessageID,
		TemplateData: templateData,
	})
	if err != nil {
		logger.Warn("Translate: translation missing", zap.String("fullMessageID", fullMessageID), zap.Error(err))
		return messageID
	}
	return message
}
