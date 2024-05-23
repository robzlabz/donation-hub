package projectstr

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"github.com/isdzulqor/donation-hub/internal/driver/request"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	sqlClient *sqlx.DB
}

func (s *Storage) IsProjectExist(ctx context.Context, id int) (exist bool, err error) {
	//TODO implement me
	panic("implement me")
}

type Config struct {
	SQLClient *sqlx.DB `validate:"required"`
}

func (c Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}

func New(cfg Config) *Storage {
	err := cfg.Validate()
	if err != nil {
		return nil
	}
	s := &Storage{sqlClient: cfg.SQLClient}
	return s
}

func (s *Storage) RequestUploadURL(ctx context.Context) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) SubmitProject(ctx context.Context, input *entity.Project) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) ReviewProject(ctx context.Context, model entity.Project, req request.StatusUpdateRequestBody) (err error) {
	// implement your logic here
	return nil
}

func (s *Storage) ListProject(ctx context.Context, limit int, page int, status string) (projects []entity.Project, err error) {
	// implement your logic here
	return nil, nil
}

func (s *Storage) ListDonation(ctx context.Context, model entity.Project, limit int, page int) (donations []entity.Donation, err error) {
	// implement your logic here
	return nil, nil
}

func (s *Storage) GetProjectDetail(ctx context.Context, project *entity.Project) (err error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) DonateProject(ctx context.Context, project entity.Project, donation entity.Donation) (err error) {
	//TODO implement me
	panic("implement me")
}
