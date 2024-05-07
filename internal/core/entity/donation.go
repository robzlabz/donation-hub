package entity

type Donation struct {
	ID        int64
	ProjectID int64
	DonorID   int64
	Message   string
	Amount    float64
	Currency  string
	CreatedAt int64
}
