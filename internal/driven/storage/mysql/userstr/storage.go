package userstr

import (
	"context"
	"fmt"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
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

func New(cfg Config) *Storage {
	err := cfg.Validate()
	if err != nil {
		return nil
	}
	s := &Storage{sqlClient: cfg.SQLClient}
	return s
}

func (s *Storage) RegisterNewUser(ctx context.Context, user *entity.User, role string) (err error) {
	query := `INSERT INTO users (username, email, password, created_at) VALUES (:username, :email, :password, :created_at)`
	_, err = s.sqlClient.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("unable to execute query due: %w", err)
	}

	// add role to user
	query = `INSERT INTO user_roles (user_id, role) VALUES (:user_id, :role)`
	_, err = s.sqlClient.NamedExecContext(ctx, query, map[string]interface{}{"user_id": user.ID, "role": role})
	if err != nil {
		return fmt.Errorf("unable to execute query due: %w", err)
	}

	return nil
}

func (s *Storage) LoginUser(ctx context.Context, user *entity.User) (err error) {
	// note: this is not a good practice to store password in plain text but this is a requirement
	query := `SELECT * FROM users WHERE username = ? AND password = ?`
	err = s.sqlClient.GetContext(ctx, user, query, user.Username, user.Password)
	if err != nil {
		return fmt.Errorf("unable to execute query: %w", err)
	}

	return nil
}

func (s *Storage) ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error) {
	// Calculate the offset for the SQL query
	offset := (page - 1) * limit

	// Prepare the SQL query
	query := `SELECT users.*, GROUP_CONCAT(user_roles.role SEPARATOR ',') AS roles
				FROM users 
				JOIN user_roles ON users.id = user_roles.user_id
				WHERE user_roles.role = ? GROUP BY users.id LIMIT ? OFFSET ? `

	// Execute the SQL query
	err = s.sqlClient.SelectContext(ctx, &users, query, role, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query: %w", err)
	}

	return users, nil
}

func (s *Storage) IsExistEmail(ctx context.Context, email string) (exist bool, err error) {
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	var count int
	err = s.sqlClient.GetContext(ctx, &count, query, email)
	if err != nil {
		return
	}
	return count > 0, nil
}

func (s *Storage) IsExistUsername(ctx context.Context, username string) (exist bool, err error) {
	query := `SELECT COUNT(*) FROM users WHERE username = ?`
	var count int
	err = s.sqlClient.GetContext(ctx, &count, query, username)
	if err != nil {
		return
	}
	return count > 0, nil
}
