package entity

var (
	UserRoleAdmin     = "admin"
	UserRoleDonor     = "donor"
	UserRoleRequester = "requester"
)

type UserRole struct {
	UserID int64
	Role   string
}
