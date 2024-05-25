package projectstr

import (
	"context"
	"fmt"

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

type imageUrls struct {
	ID        int64  `db:"project_id"`
	ProjectId int64  `db:"project_id"`
	Url       string `db:"url"`
}

type dbRequester struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

type dbProject struct {
	ID           int64  `db:"id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	DueAt        int64  `db:"due_at"`
	TargetAmount int64  `db:"target_amount"`
	Currency     string `db:"currency"`
	Status       string `db:"status"`
	RequesterID  int64  `db:"requester_id"`
	Requester    dbRequester
	ImageUrls    imageUrls
}

func (s *Storage) GetProjects(ctx context.Context, limit int, page int, status string) ([]entity.Project, error) {
	projects := make([]dbProject, 0)

	lastId := limit * (page - 1)

	queryProjects := `SELECT * FROM projects WHERE id > ? AND status = ? LIMIT ? `
	err := s.sqlClient.SelectContext(ctx, &projects, queryProjects, lastId, status, limit)
	if err != nil {
		return nil, err
	}

	fmt.Println(projects)

	// get requester from requesters table
	requesterIds := make([]int, 0)
	for _, project := range projects {
		requesterIds = append(requesterIds, int(project.RequesterID))
	}

	queryReeuster := `SELECT * FROM users WHERE id IN (?)`
	queryReeuster = s.sqlClient.Rebind(queryReeuster)
	requesters := make([]dbRequester, 0)
	s.sqlClient.Select(&requesters, queryReeuster, requesterIds)

	// get image urls from image_urls table
	projectIds := make([]int, 0)
	for _, project := range projects {
		projectIds = append(projectIds, int(project.ID))
	}

	queryImageUrls := `SELECT * FROM project_images WHERE project_id IN (?)`
	queryImageUrls = s.sqlClient.Rebind(queryImageUrls)
	imageUrls := make([]imageUrls, 0)
	s.sqlClient.Select(&imageUrls, queryImageUrls, projectIds)

	for _, project := range projects {
		for _, requester := range requesters {
			if project.RequesterID == requester.ID {
				project.Requester = requester
				break
			}
		}

		for _, imageUrl := range imageUrls {
			if project.ID == imageUrl.ID {
				project.ImageUrls = imageUrl
				break
			}
		}

	}

	fmt.Println(projects)

	// todo : convert dbProject to entity.Project

	return []entity.Project{}, nil
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
