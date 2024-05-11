// Package entity defines the User entity and its validation methods.
package entity

import (
	"github.com/isdzulqor/donation-hub/internal/common/errors"
	"strings"
)

// User is a struct that represents a user in the system.
// It includes fields for the user's ID, username, email, password, creation time, and roles.
// The struct fields have db and json tags for database and JSON marshaling/unmarshaling respectively.
type User struct {
	ID        int64   `db:"id" json:"id"`
	Username  string  `db:"username" json:"username"`
	Email     string  `db:"email" json:"email"`
	Password  string  `db:"password" json:"password"`
	CreatedAt int64   `db:"created_at" json:"created_at"`
	Roles     []uint8 `db:"roles" json:"roles"`
}

// Validate is a method of User that validates the user's username, password, and email.
// It returns an error if any of the validations fail.
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

// ValidateUsername is a method of User that validates the user's username.
// It returns an error if the username is empty.
func (u User) ValidateUsername() (err error) {
	if u.Username == "" {
		return errors.ErrUsernameIsRequired
	}
	return nil
}

// ValidatePassword is a method of User that validates the user's password.
// It returns an error if the password is empty.
func (u User) ValidatePassword() (err error) {
	if u.Password == "" {
		return errors.ErrPasswordIsRequired
	}
	return nil
}

// ValidateEmail is a method of User that validates the user's email.
// It returns an error if the email is empty or if it doesn't contain an "@" symbol.
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
