package request

type RegisterRequestBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ProjectRequestBody struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	ImageURLs    []string `json:"image_urls"`
	DueAt        int64    `json:"due_at"`
	TargetAmount int      `json:"target_amount"`
	Currency     string   `json:"currency"`
}

type StatusUpdateRequestBody struct {
	Status string `json:"status"`
}

type DonationRequestBody struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
	Message  string `json:"message"`
}
