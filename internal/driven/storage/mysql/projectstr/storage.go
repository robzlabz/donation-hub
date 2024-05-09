package projectstr

import (
	"context"

	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/rest"
	"github.com/jmoiron/sqlx"
	"gopkg.in/validator.v2"
)

type Storage struct {
	sqlClient *sqlx.DB
}

type Config struct {
	SQLClient *sqlx.DB `validate:"nonnil"`
}

func (c Config) Validate() error {
	return validator.Validate(c)
}

func New(cfg Config) (*Storage, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	s := &Storage{sqlClient: cfg.SQLClient}
	return s, nil
}

func (s *Storage) RequestUploadURL(ctx context.Context) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) SubmitProject(ctx context.Context, project entity.Project) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) ReviewProject(ctx context.Context, model entity.Project, req rest.StatusUpdateRequestBody) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) ListProject(ctx context.Context, limit int, page int, status string) (projects []entity.Project, err error) {
	// implement your logic here
	return nil, nil
}

func (s *Storage) GetProjectDetail(ctx context.Context, model entity.Project) (project entity.Project, err error) {
	// implement your logic here
	return entity.Project{}, nil
}

func (s *Storage) DonateProject(ctx context.Context, model entity.Project, req rest.DonationRequestBody) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) ListDonation(ctx context.Context, model entity.Project, limit int, page int) (donations []entity.Donation, err error) {
	// implement your logic here
	return nil, nil
}
