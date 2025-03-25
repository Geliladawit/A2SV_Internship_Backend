package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
    "errors"
    "time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"task_manager/data"
	"task_manager/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET")) 

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User) (string, error) {
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

func Authenticate(c *gin.Context) {
	tokenString := extractToken(c)

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.Set("userID", claims.UserID)
	c.Set("username", claims.Username)
	c.Set("isAdmin", claims.IsAdmin)

	c.Next()
}

func AuthorizeAdmin(c *gin.Context) {
	isAdmin, _ := c.Get("isAdmin")
	if !isAdmin.(bool) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin access required"})
		return
	}
	c.Next()
}

func extractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetUserID(c *gin.Context) string {
	userID, exists := c.Get("userID")
	if !exists {
		return "" 
	}
	return userID.(string)
}

func GetUser(c *gin.Context) (*models.User, error) {
	userIDString := GetUserID(c)
	objID, err := primitive.ObjectIDFromHex(userIDString)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}
	user, err := data.GetUserByID(objID) 
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}