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
