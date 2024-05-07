package user

import (
	"context"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/rest"
)

type UserStorage interface {
	RegisterNewUser(ctx context.Context, req rest.RegisterRequestBody) (user entity.User, err error)
	LoginUser(ctx context.Context, req rest.LoginRequestBody) (user entity.User, err error)
	ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error)
	IsExist(ctx context.Context, email string) (isExist bool, err error)
}
