// Package encryption provides functionality for JWT token generation and validation,
// as well as password hashing and checking hashed passwords.
package encryption

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// JWTService is an interface that defines methods for generating and validating JWT tokens.
type JWTService interface {
	// GenerateToken generates a JWT token for a given user.
	GenerateToken(user entity.User) (string, error)
	// ValidateToken validates a given JWT token and returns the parsed token if it's valid.
	ValidateToken(token string) (*jwt.Token, error)
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

// HashPassword hashes a given password using bcrypt with 12 rounds of hashing.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPasswordHash checks a given password against a hashed password.
// It returns true if the password matches the hashed password, and false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
