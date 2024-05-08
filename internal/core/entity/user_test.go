package entity

import (
	"testing"

	"github.com/isdzulqor/donation-hub/internal/driver/rest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateUser(t *testing.T) {
	t.Run("validate user with empty username", func(t *testing.T) {
		u := User{
			Username: "",
		}
		err := u.ValidateUsername()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrUsernameIsRequired, err)
	})

	t.Run("Successfully validate user", func(t *testing.T) {
		u := User{
			Username: "usernamevalid",
		}
		err := u.ValidateUsername()
		require.Nil(t, err)
		require.Equal(t, "usernamevalid", u.Username)
	})

	t.Run("validate user with empty password", func(t *testing.T) {
		u := User{Password: ""}
		err := u.ValidatePassword()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrPasswordIsRequired, err)
	})

	t.Run("valid password", func(t *testing.T) {
		u := User{Password: "password"}
		err := u.ValidatePassword()
		require.Nil(t, err)
		require.Equal(t, "password", u.Password)
	})

	t.Run("validate user with empty email", func(t *testing.T) {
		u := User{Email: ""}
		err := u.ValidateEmail()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrEmailIsRequired, err)
	})

	t.Run("validate using valid email", func(t *testing.T) {
		u := User{Email: "test@test.com"}
		err := u.ValidateEmail()
		require.Nil(t, err)
		require.Equal(t, "test@test.com", u.Email)
	})

	t.Run("validate user with invalid email", func(t *testing.T) {
		u := User{Email: "test"}
		err := u.ValidateEmail()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrInvalidEmail, err)
	})

	t.Run("validate with all success", func(t *testing.T) {
		u := User{
			Username: "username",
			Password: "password",
			Email:    "valid@mail.com",
		}

		err := u.Validate()
		require.Nil(t, err)
	})

	t.Run("error when one of field is null", func(t *testing.T) {
		users := []struct {
			user User
			err  error
		}{
			{
				user: User{
					Username: "username", Password: "password", Email: "",
				},
				err: rest.ErrEmailIsRequired,
			},
			{
				user: User{
					Username: "username", Password: "", Email: "email",
				},
				err: rest.ErrPasswordIsRequired,
			},
			{
				user: User{
					Username: "", Password: "password", Email: "email",
				},
				err: rest.ErrUsernameIsRequired,
			},
		}

		for _, u := range users {
			err := u.user.Validate()
			require.NotNil(t, err)
			assert.Equal(t, u.err, err)
		}
	})

}
