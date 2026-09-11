package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/pyhita/iam/internal/pkg/code"
	"github.com/pyhita/iam/internal/pkg/middleware"
)

const authHeaderCount = 2

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
// according `Authorization` header.
type AutoStrategy struct {
	basic middleware.AuthStrategy
	jwt   middleware.AuthStrategy
}

var _ middleware.AuthStrategy = &AutoStrategy{}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(basic, jwt middleware.AuthStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

func (a *AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := &middleware.AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeaderCount {
			core.WriteResponse(
				c,
				errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."),
				nil,
			)
			c.Abort()
			return
		}

		// Set auth strategy by request header.
		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "Bearer":
			operator.SetStrategy(a.jwt)
		default:
			core.WriteResponse(c, errors.WithCode(code.ErrInvalidAuthHeader, "Authorization header format is wrong."), nil)
			c.Abort()
			return
		}

		// Execute auth strategy.
		operator.AuthFunc()(c)
		c.Next()
	}
}
