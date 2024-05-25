// Package encryption provides functionality for JWT token generation and validation,
// as well as password hashing and checking hashed passwords.
package encryption

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
)

type ResponseBodyError struct {
	Ok  bool   `json:"ok"`
	Err string `json:"err"`
	Msg string `json:"msg"`
	Ts  int64  `json:"ts"`
}

// JWTService is an interface that defines methods for generating and validating JWT tokens.
type JWTService interface {
	// GenerateToken generates a JWT token for a given user.
	GenerateToken(user entity.User) (string, error)
	// ValidateToken validates a given JWT token and returns the parsed token if it's valid.
	ValidateToken(token string) (*jwt.Token, error)

	Middleware(next http.Handler, optional bool) http.Handler
}

// jwtService is a struct that implements the JWTService interface.
type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService creates a new JWTService with the given secret key and issuer.
func NewJWTService(secretKey string, issuer string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

// CustomClaims is a struct that defines the custom claims for the JWT token.
type CustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for a given user.
// The token includes custom claims (user ID, username, and email) and standard claims (expiry time and issuer).
func (s *jwtService) GenerateToken(user entity.User) (string, error) {

	// one year expiry time token
	claims := &CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    s.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

// ValidateToken validates a given JWT token and returns the parsed token if it's valid.
// It checks if the token's signing method is HMAC and if the token's signature is valid.
func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
}

func (s *jwtService) ExtractClaims(tokenStr string) (CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return *claims, nil
	} else {
		return CustomClaims{}, err
	}
}

func (s *jwtService) Middleware(next http.Handler, canBeOptional bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// If the endpoint can be optional and there's no Authorization header, pass the request to the next handler
		if canBeOptional && authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		// If the Authorization header does not contain "Bearer ", respond with an error
		if !strings.Contains(authHeader, "Bearer ") {
			responseError := ResponseBodyError{
				Ok:  false,
				Err: "ERR_INVALID_ACCESS_TOKEN",
				Msg: "invalid access token",
				Ts:  time.Now().Unix(),
			}
			JSONResponse(w, http.StatusUnauthorized, responseError)
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := s.ValidateToken(tokenString)

		// If there's an error in parsing the token or the token is not valid, respond with an error
		if err != nil || !token.Valid {
			responseError := ResponseBodyError{
				Ok:  false,
				Err: "ERR_INVALID_ACCESS_TOKEN",
				Msg: "invalid access token",
				Ts:  time.Now().Unix(),
			}
			JSONResponse(w, http.StatusUnauthorized, responseError)
			return
		}

		// If the token is valid, pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}

func JSONResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
