package repository

import (
	"context"
	"errors"

	"github.com/pyhita/webook/internal/domain"
	"github.com/pyhita/webook/internal/repository/dao"
)

var (
	ErrUserEmailDuplicate    = dao.ErrUserEmailDuplicate
	ErrInvalidUserOrPassword = dao.ErrInvalidUserOrPassword
	ErrUserNotExists         = dao.ErrUserNotExists
)

type UserRepository struct {
	userDao *dao.UserDao
}

func NewUserRepository(userDao *dao.UserDao) *UserRepository {
	return &UserRepository{userDao: userDao}
}

func (repo *UserRepository) Create(ctx context.Context, user domain.User) error {
	err := repo.userDao.Insert(ctx, dao.User{
		Email:    user.Email,
		Password: user.Password,
	})

	if errors.Is(err, ErrUserEmailDuplicate) {
		return ErrUserEmailDuplicate
	}
	return err
}

func (repo *UserRepository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := repo.userDao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (repo *UserRepository) FindUserByID(ctx context.Context, userID int64) (domain.User, error) {
	user, err := repo.userDao.FindByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		NickName: user.NickName,
		Birthday: user.Birthday,
		Intro:    user.Intro,
	}, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user domain.User) error {
	return repo.userDao.UpdateUser(ctx, dao.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		NickName: user.NickName,
		Birthday: user.Birthday,
		Intro:    user.Intro,
	})
}
