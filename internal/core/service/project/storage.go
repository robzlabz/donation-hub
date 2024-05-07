package project

import (
	"context"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/rest"
)

type ProjectStorage interface {
	RequestUploadURL(ctx context.Context) (err error)
	SubmitProject(ctx context.Context, req rest.ProjectRequestBody) (err error)
	ReviewProject(ctx context.Context, model entity.Project, req rest.StatusUpdateRequestBody) (err error)
	ListProject(ctx context.Context, limit int, page int, status string) (projects []entity.Project, err error)
	GetProjectDetail(ctx context.Context, model entity.Project) (project entity.Project, err error)
	DonateProject(ctx context.Context, model entity.Project, req rest.DonationRequestBody) (err error)
	ListDonation(ctx context.Context, model entity.Project, limit int, page int) (donations []entity.Donation, err error)
}
