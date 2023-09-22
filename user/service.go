package user

import (
	"context"

	db "github.com/dzemildupljak/web_app_with_unitest/db/pg"
)

type IUserService interface {
	ListUsers(ctx context.Context) ([]User, error)
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) ListUsers(ctx context.Context) ([]User, error) {
	repo := NewUserRepo[User](ctx)
	qo := db.QueryOptions{}

	res, err := repo.GetUsers(qo)
	if err != nil {
		return nil, err
	}

	return res, err
}
