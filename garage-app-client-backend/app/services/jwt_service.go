package services

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type JWTService struct {
	Secret string
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Создание нового JWTService
func NewJWTService() *JWTService {
	return &JWTService{Secret: os.Getenv("JWT_SECRET")}
}

// Генерация токена
func (s *JWTService) GenerateToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.Secret))
}

// Метод для проверки и парсинга токена
func (s *JWTService) ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}

//// Валидация токена
//func (s *JWTService) ValidateToken(tokenStr string) (*Claims, error) {
//	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(s.Secret), nil
//	})
//	if err != nil || !token.Valid {
//		return nil, errors.New("invalid token")
//	}
//	claims, ok := token.Claims.(*Claims)
//	if !ok {
//		return nil, errors.New("could not parse claims")
//	}
//	return claims, nil
//}
