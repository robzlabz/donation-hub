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
