package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

type respRegister struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type respSuccessLogin struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

type reqRegister struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	ImageUrls    []string `json:"image_urls"`
	DueAt        int64    `json:"due_at"`
	TargetAmount int64    `json:"target_amount"`
	Currency     string   `json:"currency"`
}

type ResponseBodySuccess struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

func SuccessResponse(w http.ResponseWriter, data interface{}) {
	resp := ResponseBodySuccess{
		Ok:   true,
		Data: data,
		Ts:   time.Now().Unix(),
	}

	// Convert the response struct to JSON
	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
