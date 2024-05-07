package user

type Service interface {
	RegisterUser() (err error)
	LoginUser() (err error)
	GetListUser() (err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) RegisterUser() (err error) {
	// todo : implement repository for register user
	return
}

func (s *service) LoginUser() (err error) {
	// todo : implement repository for login user
	return
}

func (s *service) GetListUser() (err error) {
	// todo : implement repository for get list user
	return
}
