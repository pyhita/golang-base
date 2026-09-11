package main

import (
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/pyhita/webook/config"
	redis2 "github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/pyhita/webook/internal/repository"
	"github.com/pyhita/webook/internal/repository/dao"
	"github.com/pyhita/webook/internal/service"
	"github.com/pyhita/webook/internal/web"
	"github.com/pyhita/webook/internal/web/middleware"
	"github.com/pyhita/webook/pkg/ginx/middleware/ratelimit"
)

func main() {
	db := initDB()

	server := initWebServer()
	initUserHdl(db, server)
	server.Run(":8080")
}

func initUserHdl(db *gorm.DB, server *gin.Engine) {
	ud := dao.NewUserDao(db)
	ur := repository.NewUserRepository(ud)
	us := service.NewUserService(ur)
	hdl := web.NewUserHandler(us)
	hdl.RegisterRouters(server)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/webook"))
	if err != nil {
		panic(err)
	}

	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initWebServer() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		//AllowAllOrigins: true,
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,

		AllowHeaders: []string{"Content-Type", "Authorization"},
		// 这个是允许前端访问你的后端响应中带的头部
		ExposeHeaders: []string{"x-jwt-token"},
		//AllowHeaders: []string{"content-type"},
		//AllowMethods: []string{"POST"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				//if strings.Contains(origin, "localhost") {
				return true
			}
			return strings.Contains(origin, "your_company.com")
		},
		MaxAge: 12 * time.Hour,
	}), func(ctx *gin.Context) {
		println("这是我的 Middleware")
	})

	//// 存储数据的，也就是你 userId 存哪里
	// 直接存 cookie 中，不安全
	//store := cookie.NewStore([]byte("secret"))

	// 切换成基于内存的实现，需要传两个key：authentication key 和 encryption key
	//store := memstore.NewStore([]byte("secret"), []byte("password"))

	//	[]byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwgK"),
	//	[]byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwgA"))

	// Use rate limit plugin
	redisClient := redis2.NewClient(&redis2.Options{
		Addr: config.Config.Redis.Addr,
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 1).Build())

	// 切换为redis 存储
	store, err := redis.NewStore(
		10,
		"tcp",
		"localhost:6379",
		"",
		"",
		[]byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwgK"),
		[]byte("k6CswdUm75WKcbM68UQUuxVsHSpTCwgA"),
	)
	store.Options(sessions.Options{
		MaxAge: 10,
	})
	if err != nil {
		panic(err)
	}

	// 初始化 session 缓存
	server.Use(sessions.Sessions("ssid", store))
	// 保持登录状态
	login := &middleware.LoginJWTMiddlewareBuilder{}
	server.Use(login.CheckLogin())
	return server
}
