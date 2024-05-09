package rest

import "errors"

var (
	ErrUsernameNotFound            = errors.New("Username/Password not found")
	ErrInvalidUsernameOrPassword   = errors.New("invalid username or password")
	ErrDueTimeAtLeastOneWeek       = errors.New("due time must be at least 1 week from now")
	ErrImageMustBetween1And4       = errors.New("image URLs must be between 1 and 4 URLs")
	ErrDonationAmountExceedsTarget = errors.New("donation amount exceeds target amount")
	ErrUsernameIsRequired          = errors.New("username is required")
	ErrPasswordIsRequired          = errors.New("password is required")
	ErrEmailIsRequired             = errors.New("email is required")
	ErrInvalidEmail                = errors.New("invalid email")
)

var (
	ErrBucketNameIsRequired = errors.New("BucketName is required")
	ErrObjectKeyIsRequired  = errors.New("ObjectKey is required")
)

var (
	ErrSqlClientIsRequired = errors.New("SQLClient is required")
)

var (
	ERR_INVALID_CREDS     = "ERR_UNSUPPORTED_TYPE"
	ERR_FILE_TOO_LARGE    = "ERR_FILE_TOO_LARGE"
	ERR_UNSUPPORTED_TYPE  = "ERR_UNSUPPORTED_TYPE"
	ERR_BAD_REQUEST       = "ERR_BAD_REQUEST"
	ERR_TOO_MUCH_DONATION = "ERR_TOO_MUCH_DONATION"
)
