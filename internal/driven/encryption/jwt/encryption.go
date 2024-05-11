package encryption

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/isdzulqor/donation-hub/internal/core/entity"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JWTService interface {
	GenerateToken(user entity.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService(secretKey string, issuer string) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
	}
}

type CustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(user entity.User) (string, error) {
	claims := &CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	// one year expiry time
	claims.ExpiresAt = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.secretKey), nil
	})
}

// implement hash password if needed
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// implement check password hash if needed
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
