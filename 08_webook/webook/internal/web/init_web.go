package web

import "github.com/gin-gonic/gin"

func RegisterRouters(server *gin.Engine) {
	u := NewUserHandler()
	u.registerRouters(server)
}
