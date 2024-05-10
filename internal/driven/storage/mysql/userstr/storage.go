package userstr

import (
	"context"
	"fmt"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"strings"

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

func (s *Storage) RegisterNewUser(ctx context.Context, user *entity.User) (err error) {
	query := `INSERT INTO users (username, email, password, created_at) VALUES (:username, :email, :password, :created_at)`
	_, err = s.sqlClient.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("unable to execute query due: %w", err)
	}
	return nil
}

func (s *Storage) LoginUser(ctx context.Context, req request.LoginRequestBody) (user entity.User, err error) {
	// implement your logic here
	return entity.User{}, nil
}

func (s *Storage) ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error) {
	// Calculate the offset for the SQL query
	offset := (page - 1) * limit

	// Prepare the SQL query
	query := `SELECT users.id, users.username, users.email, users.password, users.created_at, GROUP_CONCAT(user_roles.role) AS roles
        FROM users
        JOIN user_roles ON users.id = user_roles.user_id
        WHERE user_roles.role = ? GROUP BY users.id LIMIT ? OFFSET ? `

	rows, err := s.sqlClient.QueryxContext(ctx, query, role, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query: %w", err)
	}

	for rows.Next() {
		var user entity.User
		var roles string
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &roles)
		if err != nil {
			return nil, err
		}
		user.Roles = strings.Split(roles, ",")
		users = append(users, user)
	}

	return users, nil
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
