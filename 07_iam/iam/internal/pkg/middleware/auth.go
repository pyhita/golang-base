package middleware

import "github.com/gin-gonic/gin"

// AuthStrategy 策略接口
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

// AuthOperator 策略上下文
type AuthOperator struct {
	strategy AuthStrategy
}

// SetStrategy 切换策略
func (a *AuthOperator) SetStrategy(strategy AuthStrategy) {
	a.strategy = strategy
}

// AuthFunc 执行策略
func (a *AuthOperator) AuthFunc() gin.HandlerFunc {
	// 可以进行统一处理比如记录日志、metrics等
	return a.strategy.AuthFunc()
}
