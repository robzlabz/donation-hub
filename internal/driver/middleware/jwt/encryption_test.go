package encryption

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithoutHttpToken(t *testing.T) {
	// Initialize the jwtService object
	jwt := &jwtService{
		secretKey: "secret",
		issuer:    "issuer",
	}

	// Create a mock HTTP handler for the next handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a mock HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the Middleware method with the mock HTTP handler and optional flag set to false
	jwt.Middleware(nextHandler, true).ServeHTTP(rr, req)

	// Assert that the next handler was called
	assert.Equal(t, http.StatusOK, rr.Code)
}

func TestExpiredToken(t *testing.T) {
	// Initialize the jwtService object
	jwt := &jwtService{
		secretKey: "secret",
		issuer:    "issuer",
	}

	// Create a mock HTTP handler for the next handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Create a mock HTTP request with an expired token in the Authorization header
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer expired_token")

	// Create a mock HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the Middleware method with the mock HTTP handler and optional flag set to false
	jwt.Middleware(nextHandler, false).ServeHTTP(rr, req)

	// Assert that the response status code is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
