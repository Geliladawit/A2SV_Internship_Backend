package Infrastructure

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"task_manager/Domain"
)

type JWTService interface{
	GenerateToken(user *Domain.User) (string, error)
}
var jwtKey = []byte(os.Getenv("JWT_SECRET"))
type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

type JWTServiceImpl struct{}

func NewJWTService() JWTService {
	return &JWTServiceImpl{}
}
func (j *JWTServiceImpl) GenerateToken(user *Domain.User) (string, error) {
	expirationTime := jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims := &Claims{
		UserID:   user.ID.Hex(),
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}