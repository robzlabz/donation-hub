package entity

import (
	"github.com/isdzulqor/donation-hub/internal/common/errors"
	"strings"
)

type User struct {
	ID        int64   `db:"id" json:"id"`
	Username  string  `db:"username" json:"username"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"password"`
	CreatedAt int64   `db:"created_at" json:"created_at"`
	Roles     []uint8 `db:"roles" json:"roles"`
}

func (u User) Validate() (err error) {
	if err = u.ValidateUsername(); err != nil {
		return
	}
	if err = u.ValidatePassword(); err != nil {
		return
	}
	if err = u.ValidateEmail(); err != nil {
		return
	}

	return nil
}

func (u User) ValidateUsername() (err error) {
	if u.Username == "" {
		return errors.ErrUsernameIsRequired
	}
	return nil
}

func (u User) ValidatePassword() (err error) {
	if u.Password == "" {
		return errors.ErrPasswordIsRequired
	}
	return nil
}

func (u User) ValidateEmail() (err error) {
	if u.Email == "" {
		return errors.ErrEmailIsRequired
	}
	// split string by @ and ensure it has 2 parts for simple email validation
	if len(strings.Split(u.Email, "@")) != 2 {
		return errors.ErrInvalidEmail
	}
	return nil
}
