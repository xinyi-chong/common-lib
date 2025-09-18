package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xinyi-chong/common-lib/consts"
	"strings"
)

func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToCtx := map[string]string{
			consts.HeaderUserID:    consts.CtxUserID,
			consts.HeaderUsername:  consts.CtxUsername,
			consts.HeaderUserEmail: consts.CtxUserEmail,
		}

		for header, ctxKey := range headerToCtx {
			if val := strings.TrimSpace(c.GetHeader(header)); val != "" {
				c.Set(ctxKey, val)
			}
		}

		c.Next()
	}
}
