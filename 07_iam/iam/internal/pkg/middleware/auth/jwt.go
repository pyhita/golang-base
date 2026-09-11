package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/pyhita/iam/internal/pkg/middleware"
)

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ middleware.AuthStrategy = (*JWTStrategy)(nil)

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) *JWTStrategy {
	return &JWTStrategy{gjwt}
}

func (j *JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}
