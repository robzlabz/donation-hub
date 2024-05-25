package user

import (
	"context"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
)

type UserStorage interface {
	RegisterNewUser(context.Context, InputRegister) (int64, error)
	LoginUser(ctx context.Context, user *entity.User) (err error)
	ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error)
	IsExistEmail(ctx context.Context, email string) (exist bool, err error)
	IsExistUsername(ctx context.Context, username string) (exist bool, err error)
	GetRolesByEmail(ctx context.Context, email string) (roles []string, err error)
	RegisterRole(context.Context, InputRegister) (int64, error)
}
