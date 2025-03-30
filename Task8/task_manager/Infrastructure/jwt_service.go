package Infrastructure

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"task_manager/Domain"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(user Domain.User) (string, error)
	ValidateToken(tokenString string) (*Claims, error)
}

type JWTServiceImpl struct{}

func NewJWTService() JWTService {
	return &JWTServiceImpl{}
}

func (j *JWTServiceImpl) GenerateToken(user Domain.User) (string, error) {
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

func (j *JWTServiceImpl) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
