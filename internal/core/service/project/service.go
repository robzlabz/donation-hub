package project

type Service interface {
	RequestUploadURL() (err error)
	UplaodProjectImage() (err error)
	SubmitProject() (err error)
	ReviewProjectByAdmin() (err error)
	ListProject() (err error)
	GetProjectDetail() (err error)
	DonateProject() (err error)
	ListDonation() (err error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) RequestUploadURL() (err error) {
	// todo: implement repository for get upload url
	return
}

func (s *service) UplaodProjectImage() (err error) {
	// todo : implement repository for upload image
	return
}

func (s *service) SubmitProject() (err error) {
	// todo: implement repository for submit project
	return
}

func (s *service) ReviewProjectByAdmin() (err error) {
	// todo: implement repository for review project
	return
}

func (s *service) ListProject() (err error) {
	// todo: implement repository for list project
	return
}

func (s *service) GetProjectDetail() (err error) {
	// todo: implement repository for get project detail
	return
}

func (s *service) DonateProject() (err error) {
	// todo: implement repository for donate project
	return
}

func (s *service) ListDonation() (err error) {
	// todo : implement repository for list donation
	return
}
