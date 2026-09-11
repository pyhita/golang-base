package domain

// 业务对象，业务逻辑对应的对象
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Birthday string `json:"birthday"`
	Intro    string `json:"intro"`

	Ctime int64 `json:"ctime"`
}
