package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/pyhita/iam/internal/pkg/code"

	"github.com/pyhita/iam/internal/pkg/middleware"
)

// BasicStrategy 策略实现类，实现 Basic 认证
type BasicStrategy struct {
	compare func(username, password string) bool
}

var _ middleware.AuthStrategy = &BasicStrategy{}

func NewBasicStrategy(compare func(username, password string) bool) *BasicStrategy {
	return &BasicStrategy{
		compare: compare,
	}
}

func (s *BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong."),
				nil,
			)
			c.Abort()
			return
		}

		if !s.compare(username, password) {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrSignatureInvalid, "Authorization header format is wrong."),
				nil,
			)
			c.Abort()
			return
		}

		c.Set(middleware.UsernameKey, username)
		c.Next()
	}
}
