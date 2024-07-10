package repository

import (
	"context"
	"go-basic-scr/geektime/week02/webook/internal/domain"
	"go-basic-scr/geektime/week02/webook/internal/repository/dao"
	"time"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrorUserNotFound = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{dao: dao}
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {
	return repo.dao.Insert(ctx, dao.User{Email: u.Email, Password: u.Password})
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := repo.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(u), err
}

func (repo *UserRepository) toDomain(u dao.User) domain.User {

	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
		Nickname: u.Nickname,
		AboutMe:  u.AboutMe,
		Birthday: time.UnixMilli(u.Birthday),
	}

}

func (repo *UserRepository) UpdateNonZeroFields(ctx context.Context, user domain.User) error {
	userDao := repo.toEntity(user)
	err := repo.dao.UpdateById(ctx, userDao)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) toEntity(user domain.User) dao.User {
	return dao.User{
		Id:       user.Id,
		Nickname: user.Nickname,
		Birthday: user.Birthday.UnixMilli(),
		AboutMe:  user.AboutMe,
	}
}

func (repo *UserRepository) FindById(ctx context.Context, uid int64) (domain.User, error) {
	user, err := repo.dao.FindById(ctx, uid)
	if err != nil {
		return domain.User{}, err
	}
	return repo.toDomain(user), nil
}
