package rest

import (
	"net/http"
	"strings"
)

func parseBearerToken(r *http.Request) string {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	if len(token) != 2 {
		return ""
	}
	if strings.ToLower(token[0]) != "bearer" {
		return ""
	}
	return token[1]
}

func GetPageLimit(r *http.Request) (page, limit int) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	if pageStr == "" {
		page = 1
	} else {
		page = 1
	}
	if limitStr == "" {
		limit = 10
	} else {
		limit = 10
	}
	return
}
