package web

import (
	"net/http"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// 处理user相关的请求
type UserHandler struct {
	emailRegex *regexp.Regexp
	passRegex  *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		emailRegex: regexp.MustCompile(emailRegexPattern, regexp.None),
		passRegex:  regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	userGroup := server.Group("/user")

	// register user routers
	userGroup.POST("/login", u.Login)
	userGroup.POST("/edit", u.Edit)
	userGroup.POST("/signup", u.SignUp)
	userGroup.GET("/profile", u.Profile)
}

// 通用逻辑：
// 1. 接受请求，参数解析
// 2. 参数校验
// 3. 业务逻辑
// 4. 响应结果

func (u *UserHandler) SignUp(c *gin.Context) {
	type SingUpReq struct {
		email           string `json:"email"`
		password        string `json:"password"`
		confirmPassword string `json:"confirmPassword"`
	}

	var req SingUpReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Invalid req")
		return
	}

	if req.password != req.confirmPassword {
		c.JSON(http.StatusBadRequest, "两次输入密码不一致")
		return
	}

	ok, err := u.passRegex.MatchString(req.password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, "密码格式不正确")
		return
	}

	ok, err = u.emailRegex.MatchString(req.email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !ok {
		c.JSON(http.StatusBadRequest, "邮箱格式不正确")
		return
	}

	c.JSON(200, "注册成功")
}

func (u *UserHandler) Login(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {

}

func (u *UserHandler) Edit(c *gin.Context) {

}
