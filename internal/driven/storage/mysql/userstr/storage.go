package userstr

import (
	"context"
	"fmt"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/rest"
	"github.com/jmoiron/sqlx"
	"gopkg.in/validator.v2"
)

type Storage struct {
	sqlClient *sqlx.DB
}

type Config struct {
	SQLClient *sqlx.DB `validate:"nonnil"`
}

func (c Config) Validate() error {
	return validator.Validate(c)
}

func New(cfg Config) (*Storage, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	s := &Storage{sqlClient: cfg.SQLClient}
	return s, nil
}

func (s *Storage) RegisterNewUser(ctx context.Context, user *entity.User) (err error) {
	query := `INSERT INTO users (username, email, password, created_at) VALUES (:username, :email, :password, :created_at)`
	_, err = s.sqlClient.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("unable to execute query due: %w", err)
	}
	return nil
}

func (s *Storage) LoginUser(ctx context.Context, req rest.LoginRequestBody) (user entity.User, err error) {
	// implement your logic here
	return entity.User{}, nil
}

func (s *Storage) ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error) {
	// implement your logic here
	return nil, nil
}

func (s *Storage) IsExist(ctx context.Context, email string) (exist bool, err error) {
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	var count int
	err = s.sqlClient.GetContext(ctx, &count, query, email)
	if err != nil {
		return
	}
	return count > 0, nil
}
