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

	t.Run("validate user with empty password", func(t *testing.T) {
		u := User{Password: ""}
		err := u.ValidatePassword()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrPasswordIsRequired, err)
	})

	t.Run("validate user with empty email", func(t *testing.T) {
		u := User{Email: ""}
		err := u.ValidateEmail()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrEmailIsRequired, err)
	})

	t.Run("validate user with invalid email", func(t *testing.T) {
		u := User{Email: "test"}
		err := u.ValidateEmail()
		require.NotNil(t, err)
		assert.Equal(t, rest.ErrInvalidEmail, err)
	})
}
