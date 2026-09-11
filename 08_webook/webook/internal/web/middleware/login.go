package middleware

import (
	"encoding/gob"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginMiddlewareBuilder struct {
}

func (b *LoginMiddlewareBuilder) CheckLogin() gin.HandlerFunc {
	// 注册一下这个类型
	gob.Register(time.Now())
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if path == "/users/login" ||
			path == "/users/signup" {
			return
		}

		sess := sessions.Default(c)
		userID := sess.Get("userID")
		if userID == nil {
			// 用户未登录进行拦截
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()
		// 刷新 session 信息
		const updateTimeKey = "updateTime"
		// 拿到上一次的刷新时间
		val := sess.Get(updateTimeKey)
		lastUpdateTime, ok := val.(time.Time)
		if val == nil || (ok && now.Sub(lastUpdateTime) > time.Duration(10)*time.Second) {
			sess.Set(updateTimeKey, now)
			sess.Set("userID", userID)
			if err := sess.Save(); err != nil {
				log.Println("session save error:", err)
			}
		}
	}
}
