package user

import (
	"context"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
)

type UserStorage interface {
	RegisterNewUser(ctx context.Context, user *entity.User) (err error)
	LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, err error)
	ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error)
	IsExistEmail(ctx context.Context, email string) (exist bool, err error)
	IsExistUsername(ctx context.Context, username string) (exist bool, err error)
}
