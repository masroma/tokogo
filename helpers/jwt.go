package helpers

import (
	"errors"
	"time"

	"tokogo/models"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("your-secret-key") // Gunakan dari environment variable

// Claims struct untuk JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken menghasilkan JWT token untuk user
func GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam

	claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
