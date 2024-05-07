package user

import (
	"context"

	"github.com/isdzulqor/donation-hub/internal/driver/rest"
)

type Service interface {
	RegisterUser(ctx context.Context, rest rest.RegisterRequestBody) (err error)
	LoginUser(ctx context.Context, rest rest.LoginRequestBody) (err error)
	GetListUser(ctx context.Context, limit int, page int, role string) (err error)
}

type service struct {
	storage UserStorage
}

func NewService(storage UserStorage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) RegisterUser(ctx context.Context, rest rest.RegisterRequestBody) (err error) {
	// check if username is valid
	// check if password is valid
	// check if email is valid
	// check if user email has aready exists

	return
}

func (s *service) LoginUser(ctx context.Context, rest rest.LoginRequestBody) (err error) {
	// todo : implement repository for login user
	return
}

func (s *service) GetListUser(ctx context.Context, limit int, page int, role string) (err error) {
	// todo : implement repository for get list user
	return
}
