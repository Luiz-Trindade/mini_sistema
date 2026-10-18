package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("sua-chave-secreta-super-segura-aqui")

type JWTClaims struct {
	UserID       uint   `json:"user_id"`
	Email        string `json:"email"`
	IsAdmin      bool   `json:"is_admin"`
	TokenVersion int    `json:"token_version"` // Incluído na claim
	jwt.RegisteredClaims
}

type JWTService struct{}

func NewJWTService() *JWTService {
	return &JWTService{}
}

func (s *JWTService) GenerateToken(userID uint, email string, isAdmin bool, tokenVersion int) (string, error) {
	claims := &JWTClaims{
		UserID:       userID,
		Email:        email,
		IsAdmin:      isAdmin,
		TokenVersion: tokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (s *JWTService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido ou expirado")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("falha ao ler as informações do token")
	}

	return claims, nil
}
