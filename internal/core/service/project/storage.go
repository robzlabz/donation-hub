package project

import (
	"context"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
)

type ProjectStorage interface {
	RequestUploadURL(ctx context.Context) (err error)
	SubmitProject(ctx context.Context, proejct entity.Project) (err error)
	ReviewProject(ctx context.Context, model entity.Project, req request.StatusUpdateRequestBody) (err error)
	ListProject(ctx context.Context, limit int, page int, status string) (projects []entity.Project, err error)
	GetProjectDetail(ctx context.Context, project *entity.Project) (err error)
	DonateProject(ctx context.Context, project entity.Project, donation entity.Donation) (err error)
	ListDonation(ctx context.Context, model entity.Project, limit int, page int) (donations []entity.Donation, err error)
	IsProjectExist(ctx context.Context, id int) (exist bool, err error)
}

type ObjectStorage interface {
	GetUploadedUrl() (url string, err error)
}
