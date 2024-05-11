package rest

import (
	"encoding/json"
	"net/http"
	"time"
)

type ResponseBodyError struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
	Msg string `json:"msg"`
	Ts  int64  `json:"ts"`
}

type HttpError struct{}

func (h *HttpError) ErrBadRequest(w http.ResponseWriter, msg string) {
	resp := ResponseBodyError{
		Ok:  false,
		Err: "ERR_BAD_REQUEST",
		Msg: msg,
		Ts:  time.Now().Unix(),
	}

	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResponse)
}

// invalid access token
func (h *HttpError) ErrUnauthorized(w http.ResponseWriter, msg string) {
	resp := ResponseBodyError{
		Ok:  false,
		Err: "ERR_INVALID_ACCESS_TOKEN",
		Msg: msg,
		Ts:  time.Now().Unix(),
	}

	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(jsonResponse)
}

// forbidden access
func (h *HttpError) ErrForbidden(w http.ResponseWriter) {
	resp := ResponseBodyError{
		Ok:  false,
		Err: "ERR_FORBIDDEN",
		Msg: "user doesn't have enough authorization",
		Ts:  time.Now().Unix(),
	}

	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write(jsonResponse)
}

// not found
func (h *HttpError) ErrNotFound(w http.ResponseWriter) {
	resp := ResponseBodyError{
		Ok:  false,
		Err: "ERR_NOT_FOUND",
		Msg: "resource not found",
		Ts:  time.Now().Unix(),
	}

	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResponse)
}

// internal error
func (h *HttpError) ErrInternalServer(w http.ResponseWriter, msg string) {
	resp := ResponseBodyError{
		Ok:  false,
		Err: "ERR_INTERNAL_SERVER",
		Msg: msg,
		Ts:  time.Now().Unix(),
	}

	jsonResponse, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResponse)
}
