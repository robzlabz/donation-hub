package project

import (
	"context"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/rest"
)

type Service interface {
	RequestUploadURL(ctx context.Context, mime_type string, file_size int64) (err error)
	SubmitProject(ctx context.Context, project entity.Project, req rest.ProjectRequestBody) (err error)
	ReviewProjectByAdmin(ctx context.Context, req rest.StatusUpdateRequestBody) (err error)
	ListProject(ctx context.Context) (projects []entity.Project, err error)
	GetProjectDetail(ctx context.Context, model entity.Project) (project entity.Project, err error)
	DonateProject(ctx context.Context, req rest.DonationRequestBody) (err error)
	ListDonation(ctx context.Context, model entity.Project) (donations []entity.Donation, err error)
}

type service struct {
	storage ProjectStorage
}

func NewService(storage ProjectStorage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) RequestUploadURL(ctx context.Context, mime_type string, file_size int64) (err error) {
	// todo: implement repository for get upload url
	return
}

func (s *service) SubmitProject(ctx context.Context, project entity.Project, req rest.ProjectRequestBody) (err error) {
	// todo: implement repository for submit project
	return
}

func (s *service) ReviewProjectByAdmin(ctx context.Context, req rest.StatusUpdateRequestBody) (err error) {
	// todo: implement repository for review project
	return
}

func (s *service) ListProject(ctx context.Context) (projects []entity.Project, err error) {
	// todo: implement repository for list project
	return
}

func (s *service) GetProjectDetail(ctx context.Context, model entity.Project) (project entity.Project, err error) {
	// todo: implement repository for get project detail
	return
}

func (s *service) DonateProject(ctx context.Context, req rest.DonationRequestBody) (err error) {
	// todo: implement repository for donate project
	return
}

func (s *service) ListDonation(ctx context.Context, model entity.Project) (donations []entity.Donation, err error) {
	// todo : implement repository for list donation
	return
}
