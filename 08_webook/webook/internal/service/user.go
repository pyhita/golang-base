package service

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/pyhita/webook/internal/domain"
	"github.com/pyhita/webook/internal/repository"
)

var (
	ErrUserEmailDuplicate    = repository.ErrUserEmailDuplicate
	ErrInvalidUserOrPassword = errors.New("用户不存在或密码不对")
	ErrUserNotExist          = repository.ErrUserNotExists
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Signup(ctx context.Context, user domain.User) error {
	// 将密码加密之后再存入DB
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)

	err = u.userRepo.Create(ctx, user)
	if errors.Is(err, ErrUserEmailDuplicate) {
		return ErrUserEmailDuplicate
	}
	return err
}

func (u *UserService) Login(ctx context.Context, email, password string) (domain.User, error) {
	user, err := u.userRepo.FindUserByEmail(ctx, email)
	if errors.Is(err, repository.ErrInvalidUserOrPassword) {
		return user, ErrInvalidUserOrPassword
	}

	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, ErrInvalidUserOrPassword
	}
	return user, nil
}

func (u *UserService) Edit(ctx context.Context, user domain.User) error {
	// Check user exist
	res, err := u.userRepo.FindUserByID(ctx, user.ID)
	if err != nil {
		return err
	}

	// 赋值操作 deepcopy
	if user.Email != "" {
		res.Email = user.Email
	}
	if user.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		res.Password = string(hash)
	}
	if user.NickName != "" {
		res.NickName = user.NickName
	}
	if user.Birthday != "" {
		res.Birthday = user.Birthday
	}
	if user.Intro != "" {
		res.Intro = user.Intro
	}

	return u.userRepo.UpdateUser(ctx, res)
}

func (u *UserService) Profile(ctx context.Context, userID int64) (domain.User, error) {
	user, err := u.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
