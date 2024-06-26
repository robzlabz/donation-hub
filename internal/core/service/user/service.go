package user

import (
	"context"
	"errors"
	errors2 "github.com/isdzulqor/donation-hub/internal/common/errors"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"time"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
)

type Service interface {
	RegisterUser(ctx context.Context, req request.RegisterRequestBody) (user entity.User, err error)
	LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, err error)
	GetListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error)
}

type service struct {
	storage UserStorage
}

func NewService(storage UserStorage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) RegisterUser(ctx context.Context, req request.RegisterRequestBody) (user entity.User, err error) {
	user = entity.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	}

	// check if username, email and password are valid
	err = user.Validate()
	if err != nil {
		return
	}

	// check if user has already used
	existEmail, err := s.storage.IsExistEmail(ctx, user.Email)
	if err != nil {
		return
	}
	if existEmail {
		err = errors.New("email already used")
		return
	}

	existUsername, err := s.storage.IsExistUsername(ctx, user.Username)
	if err != nil {
		return
	}

	if existUsername {
		err = errors.New("username already exist")
		return
	}

	user.CreatedAt = time.Now().Unix()

	err = s.storage.RegisterNewUser(ctx, &user)

	return
}

func (s *service) LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, err error) {
	exist, err := s.storage.IsExistEmail(ctx, req.Email)
	if err != nil {
		return
	}
	if !exist {
		err = errors2.ErrInvalidUsernameOrPassword
		return
	}

	user, err = s.storage.LoginUser(ctx, req)
	if err != nil {
		return
	}

	return user, nil
}

func (s *service) GetListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error) {
	return s.storage.ListUser(ctx, limit, page, role)
}
