package userstr

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/core/service/user"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

type Storage struct {
	sqlClient *sqlx.DB
}

type Config struct {
	SQLClient *sqlx.DB `validate:"required"`
}

func (c Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(c)
	if err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}
	return nil
}

func New(cfg Config) (*Storage, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	s := &Storage{sqlClient: cfg.SQLClient}
	return s, nil
}

func (s *Storage) RegisterNewUser(ctx context.Context, input user.InputRegister) (int64, error) {
	// insert user data
	query := `INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)`
	res, err := s.sqlClient.ExecContext(ctx, query, input.Username, input.Email, input.Password, time.Now().Unix())
	if err != nil {
		return 0, fmt.Errorf("unable to execute query: %w", err)
	}

	// get user id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("unable to get last insert id: %w", err)
	}

	// create user role
	id, err = s.RegisterRole(ctx, input)
	if err != nil {
		return 0, err
	}

	return id, nil
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

type DBUser struct {
	ID        int64  `db:"id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt int64  `db:"created_at"`
	Roles     string `db:"roles"`
}

func (s *Storage) ListUser(ctx context.Context, limit int, page int, role string) (users []entity.User, err error) {
	// Calculate the offset for the SQL query
	offset := (page - 1) * limit

	// Prepare the SQL query
	query := `SELECT users.*, GROUP_CONCAT(user_roles.role SEPARATOR ',') AS roles
				FROM users 
				JOIN user_roles ON users.id = user_roles.user_id
				WHERE user_roles.role = ? GROUP BY users.id LIMIT ? OFFSET ? `

	dbuser := []DBUser{}

	// Execute the SQL query
	err = s.sqlClient.SelectContext(ctx, &dbuser, query, role, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query: %w", err)
	}

	// Convert the DBUser struct to the entity.User struct
	for _, u := range dbuser {
		users = append(users, entity.User{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			Password:  u.Password,
			CreatedAt: u.CreatedAt,
			Roles:     strings.Split(u.Roles, ","),
		})
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

func (s *Storage) GetRolesByEmail(ctx context.Context, email string) (roles []string, err error) {
	query := `SELECT role FROM user_roles JOIN users ON users.id = user_roles.user_id WHERE users.email = ?`
	err = s.sqlClient.SelectContext(ctx, &roles, query, email)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query: %w", err)
	}
	return roles, nil
}

func (s *Storage) RegisterRole(ctx context.Context, input user.InputRegister) (int64, error) {
	// get user id by email
	query := `SELECT id FROM users WHERE email = ?`
	var id int64
	err := s.sqlClient.GetContext(ctx, &id, query, input.Email)
	if err != nil {
		return 0, fmt.Errorf("unable to execute query: %w", err)
	}

	// add role to user
	query = `INSERT INTO user_roles (user_id, role) VALUES (?, ?)`
	_, err = s.sqlClient.ExecContext(ctx, query, id, input.Role)
	if err != nil {
		return 0, fmt.Errorf("unable to execute query: %w", err)
	}

	return id, nil
}
