package user

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	errors2 "github.com/isdzulqor/donation-hub/internal/common/errors"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/middleware/jwt"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"slices"
)

type InputRegister struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string `validate:"required,oneof=donor requester"`
}

func (r InputRegister) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(r)
	if err != nil {
		return err
	}

	return nil
}

type OutputRegister struct {
	ID       int64
	Username string
	Email    string
}

type Service interface {
	RegisterUser(context.Context, InputRegister) (*OutputRegister, error)
	LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, accessToken string, err error)
	GetListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error)
}

type service struct {
	storage    UserStorage
	encryption encryption.JWTService
}

func (s *service) RegisterUser(ctx context.Context, register InputRegister) (*OutputRegister, error) {

	var userId int64

	// validate user input
	err := register.Validate()
	if err != nil {
		return nil, err
	}

	// check if user with email and role is existed
	emailExist, err := s.storage.IsExistEmail(ctx, register.Email)

	// if user with email and role is existed, return error
	roles, err := s.storage.GetRolesByEmail(ctx, register.Email)
	if emailExist && slices.Contains(roles, register.Role) {
		return nil, fmt.Errorf("user with email %s and role %s is already registered", register.Email, register.Role)
	}

	// if user email registered but different role, create role
	if emailExist && !slices.Contains(roles, register.Role) {
		// register role
		userId, err = s.storage.RegisterRole(ctx, register)
		if err != nil {
			return nil, err
		}
	} else {
		// if user email not registered, create user and register role
		userId, err = s.storage.RegisterNewUser(ctx, register)
		if err != nil {
			return nil, err
		}
	}

	return &OutputRegister{
		ID:       userId,
		Username: register.Username,
		Email:    register.Email,
	}, nil
}

func NewService(storage UserStorage, jwtService encryption.JWTService) Service {
	return &service{
		storage:    storage,
		encryption: jwtService,
	}
}

func (s *service) LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, accessToken string, err error) {
	exist, err := s.storage.IsExistUsername(ctx, req.Username)
	if err != nil {
		return
	}
	if !exist {
		err = errors2.ErrInvalidUsernameOrPassword
		return
	}

	user = entity.User{
		Username: req.Username,
		Password: req.Password,
	}

	err = s.storage.LoginUser(ctx, &user)
	if err != nil {
		return
	}

	accessToken, err = s.encryption.GenerateToken(user)

	return user, accessToken, nil
}

func (s *service) GetListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error) {
	return s.storage.ListUser(ctx, limit, page, role)
}
