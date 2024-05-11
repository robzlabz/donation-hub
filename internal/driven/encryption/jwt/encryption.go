package encryption

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JWTService interface {
	GenerateToken(userID int64, role string) (string, error)
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
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(userID int64, role string) (string, error) {
	claims := &CustomClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    s.issuer,
		},
	}

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
