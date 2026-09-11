package web

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"time"

	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/pyhita/webook/internal/domain"
	"github.com/pyhita/webook/internal/service"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
)

// 处理user相关的请求
type UserHandler struct {
	svc        *service.UserService
	emailRegex *regexp.Regexp
	passRegex  *regexp.Regexp
}

func NewUserHandler(usvc *service.UserService) *UserHandler {
	return &UserHandler{
		svc:        usvc,
		emailRegex: regexp.MustCompile(emailRegexPattern, regexp.None),
		passRegex:  regexp.MustCompile(passwordRegexPattern, regexp.None),
	}
}

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	userGroup := server.Group("/users")

	// register user routers
	userGroup.POST("/login", u.LoginJWT)
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
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	if err := c.Bind(&req); err != nil {
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, "两次输入密码不一致")
		return
	}

	//ok, err := u.passRegex.MatchString(req.password)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//if !ok {
	//	c.JSON(http.StatusBadRequest, "密码格式不正确")
	//	return
	//}

	//ok, err := u.emailRegex.MatchString(req.email)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//if !ok {
	//	c.JSON(http.StatusBadRequest, "邮箱格式不正确")
	//	return
	//}

	err := u.svc.Signup(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})

	switch err {
	case nil:
		c.JSON(http.StatusOK, "注册成功")
	case service.ErrUserEmailDuplicate:
		c.JSON(http.StatusBadRequest, "邮箱已存在")
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return
	}

	user, err := u.svc.Login(c, req.Email, req.Password)

	switch err {
	case nil:
		// 登录成功，将登录信息保存到session中
		sess := sessions.Default(c)
		sess.Options(sessions.Options{
			// 有效时间 15min
			MaxAge: 30,
		})
		sess.Set("userID", user.ID)
		err = sess.Save()
		if err != nil {
			log.Println(err)
			c.String(http.StatusOK, "系统错误")
			return
		}
		c.JSON(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		c.JSON(http.StatusUnauthorized, "用户名或者密码不对")
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (h *UserHandler) LoginJWT(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc := UserClaims{
			Uid:       u.ID,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				// 1 分钟过期
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
		tokenStr, err := token.SignedString(JWTKey)
		if err != nil {
			ctx.String(http.StatusOK, "系统错误")
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.String(http.StatusOK, "登录成功")
	case service.ErrInvalidUserOrPassword:
		ctx.String(http.StatusOK, "用户名或者密码不对")
	default:
		ctx.String(http.StatusOK, "系统错误")
	}
}

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}

func (u *UserHandler) Edit(c *gin.Context) {
	type EditUserReq struct {
		UserID   *int64  `json:"user_id"`
		Email    *string `json:"email"`
		Password *string `json:"password"`
		NickName *string `json:"nickname"`
		Birthday *string `json:"birthday"`
		Intro    *string `json:"intro"`
	}

	var req EditUserReq
	if err := c.Bind(&req); err != nil {
		return
	}

	// 校验
	if len(*req.NickName) > 100 ||
		len(*req.Intro) > 100 {
		c.JSON(http.StatusBadRequest, "请求参数不合法")
		return
	}

	if req.UserID == nil {
		c.JSON(http.StatusBadRequest, "用户ID不能为空")
		return
	}

	err := u.svc.Edit(c, domain.User{
		ID:       *req.UserID,
		Email:    *req.Email,
		Password: *req.Password,
		NickName: *req.NickName,
		Birthday: *req.Birthday,
		Intro:    *req.Intro,
	})

	switch err {
	case nil:
		c.JSON(http.StatusOK, "修改成功")
	case err, service.ErrUserNotExist:
		c.JSON(http.StatusBadRequest, "用户不存在")
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (u *UserHandler) Profile(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, "用户ID不能为空")
		return
	}

	uID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user, err := u.svc.Profile(c, int64(uID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
