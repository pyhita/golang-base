package dao

import (
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserEmailDuplicate    = errors.New("email already exists")
	ErrInvalidUserOrPassword = gorm.ErrRecordNotFound
	ErrUserNotExists         = gorm.ErrRecordNotFound
)

// 封装操作数据库的逻辑
type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// 定义增删改查 User 方法
func (u *UserDao) Insert(ctx context.Context, user User) error {
	now := time.Now().UnixMilli()
	user.Ctime = now
	user.Utime = now

	err := u.db.WithContext(ctx).Create(&user).Error

	var me *mysql.MySQLError
	if errors.As(err, &me) {
		if me.Number == 1062 {
			return ErrUserEmailDuplicate
		}
	}
	return err
}

func (u *UserDao) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := u.db.WithContext(ctx).First(&user, "email = ?", email).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserDao) FindByID(ctx context.Context, userID int64) (User, error) {
	var user User
	err := u.db.WithContext(ctx).First(&user, "id = ?", userID).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *UserDao) UpdateUser(ctx context.Context, user User) error {
	ret := u.db.WithContext(ctx).Where("id = ?", user.ID).Updates(&user)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

// 完全对应数据库中的表 PO
type User struct {
	ID       int64  `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"type:varchar(100);unique"`
	Password string `gorm:"type:varchar(100)"`
	NickName string `gorm:"type:varchar(100)"`
	Birthday string `gorm:"type:varchar(100)"`
	Intro    string `gorm:"type:varchar(100)"`
	Ctime    int64
	Utime    int64
}
