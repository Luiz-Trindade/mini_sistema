package services

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mini_sistema/auth/models"
	"mini_sistema/auth/selectors"
)

type AuthService struct {
	db           *gorm.DB
	userSelector *selectors.UserSelector
}

func NewAuthService(db *gorm.DB, userSelector *selectors.UserSelector) *AuthService {
	return &AuthService{
		db:           db,
		userSelector: userSelector,
	}
}

func (s *AuthService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userSelector.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("credenciais inválidas")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(password)); err != nil {
		return nil, errors.New("credenciais inválidas")
	}

	now := time.Now()

	err = s.db.Model(user).Updates(map[string]interface{}{
		"ultimo_login":  now,
		"token_version": gorm.Expr("token_version + 1"),
	}).Error

	if err != nil {
		return nil, err
	}

	user.UltimoLogin = &now
	user.TokenVersion++

	return user, nil
}

// RevokeUserTokens incrementa a token_version no banco, invalidando todos os JWTs emitidos anteriormente
func (s *AuthService) RevokeUserTokens(userID uint) error {
	return s.db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("token_version", gorm.Expr("token_version + 1")).Error
}
