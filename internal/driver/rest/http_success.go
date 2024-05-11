package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponseBodySuccess struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
	Ts   int64       `json:"ts"`
}

type HttpSuccess struct{}

func (h *HttpSuccess) SuccessResponse(w http.ResponseWriter, data interface{}) {
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
