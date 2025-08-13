package main

import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pyhita/webook/internal/web"
	"strings"
	"time"
)

func main() {
	server := gin.Default()

	// 配置跨域
	server.Use(cors.New(cors.Config{
		AllowedOrigins: []string{"https://localhost:3000"},
		//AllowedMethods:     []string{"PUT", "PATCH"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		//ExposedHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return strings.Contains(origin, "http://localhost:3000")
		},
		MaxAge: 12 * time.Hour,
	}))

	web.RegisterRouters(server)
	server.Run(":8080")
}
