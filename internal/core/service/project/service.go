package project

import (
	"context"
	"errors"
	"github.com/isdzulqor/donation-hub/internal/driven/storage/mysql/projectstr"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"time"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
)

type Service interface {
	RequestUploadURL(ctx context.Context, mimeType string, fileSize int64) (err error)
	SubmitProject(ctx context.Context, req request.ProjectRequestBody) (err error)
	ReviewProjectByAdmin(ctx context.Context, req request.StatusUpdateRequestBody) (err error)
	ListProject(ctx context.Context, limit int, page int, status string) (projects []entity.Project, err error)
	GetProjectDetail(ctx context.Context, model *entity.Project) (err error)
	DonateProject(ctx context.Context, req request.DonationRequestBody) (err error)
	ListDonation(ctx context.Context, model entity.Project, limit int, page int) (donations []entity.Donation, err error)
}

type service struct {
	storage ProjectStorage
}

func NewService(storage *projectstr.Storage) Service {
	return &service{
		storage: storage,
	}
}

func (s *service) RequestUploadURL(context.Context, string, int64) (err error) {
	// todo: implement repository for get upload url
	return
}

func (s *service) SubmitProject(ctx context.Context, req request.ProjectRequestBody) (err error) {
	project := entity.Project{
		Name:             req.Title,
		Description:      req.Description,
		TargetAmount:     float64(req.TargetAmount),
		Currency:         req.Currency,
		CollectionAmount: 0,
		Status:           entity.StatusNeedReview,
		DueAt:            req.DueAt,
		RequesterID:      1, // todo: get from jwt
		CreatedAt:        time.Now().Unix(),
		UpdatedAt:        time.Now().Unix(),
	}

	err = s.storage.SubmitProject(ctx, project)
	if err != nil {
		return
	}

	return
}

func (s *service) ReviewProjectByAdmin(ctx context.Context, req request.StatusUpdateRequestBody) (err error) {
	// todo: implement repository for review project
	return
}

func (s *service) ListProject(ctx context.Context, limit int, page int, status string) (projects []entity.Project, err error) {
	return s.storage.ListProject(ctx, limit, page, status)
}

func (s *service) GetProjectDetail(ctx context.Context, project *entity.Project) (err error) {
	exist, err := s.storage.IsProjectExist(ctx, int(project.ID))
	if err != nil {
		return
	}
	if !exist {
		err = errors.New("project not exist")
		return
	}

	err = s.storage.GetProjectDetail(ctx, project)
	if err != nil {
		return
	}

	return
}

func (s *service) DonateProject(ctx context.Context, req request.DonationRequestBody) (err error) {
	donation := entity.Donation{
		ProjectID: 1, // add project id,
		DonorID:   1, // todo : get from jwt
		Message:   req.Message,
		Amount:    float64(req.Amount),
		Currency:  req.Currency,
		CreatedAt: time.Now().Unix(),
	}

	err = s.storage.DonateProject(ctx, entity.Project{ID: 1}, donation)
	if err != nil {
		return
	}

	return
}

func (s *service) ListDonation(ctx context.Context, model entity.Project, limit int, page int) (donations []entity.Donation, err error) {
	return s.storage.ListDonation(ctx, model, limit, page)
}
