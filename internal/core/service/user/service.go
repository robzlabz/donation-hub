package user

type Service interface {
	RegisterUser() (err error)
	LoginUser() (err error)
	GetListUser() (err error)
}
