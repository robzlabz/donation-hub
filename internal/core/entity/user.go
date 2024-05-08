package entity

import (
	"strings"
	"time"

	"github.com/isdzulqor/donation-hub/internal/driver/rest"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
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
		return rest.ErrUsernameIsRequired
	}
	return nil
}

func (u User) ValidatePassword() (err error) {
	if u.Password == "" {
		return rest.ErrPasswordIsRequired
	}
	return nil
}

func (u User) ValidateEmail() (err error) {
	if u.Email == "" {
		return rest.ErrEmailIsRequired
	}
	// split string by @ and ensure it has 2 parts for simple email validation
	if len(strings.Split(u.Email, "@")) != 2 {
		return rest.ErrInvalidEmail
	}
	return nil
}
