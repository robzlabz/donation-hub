package entity

var (
	StatusNeedReview = "need_review"
	StatusApproved   = "approved"
	StatusCompleted  = "completed"
	StatusRejected   = "rejected"
)

type Project struct {
	ID               int64
	Name             string
	Description      string
	TargetAmount     float64
	CollectionAmount float64
	Currency         string
	Status           string
	RequesterID      int
	DueAt            int64
	CreatedAt        int64
	UpdatedAt        int64
}
