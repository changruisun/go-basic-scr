package repository

import (
	"context"
	"go-basic-scr/geektime/week02/webook/internal/domain"
	"go-basic-scr/geektime/week02/webook/internal/repository/dao"
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
	}

}
