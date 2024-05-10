package user

import (
	"context"
	"github.com/isdzulqor/donation-hub/internal/driver/request"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
)

type UserStorage interface {
	RegisterNewUser(ctx context.Context, user *entity.User) (err error)
	LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, err error)
	ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error)
	IsExist(ctx context.Context, email string) (exist bool, err error)
}
