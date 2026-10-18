package selectors

import (
	"errors"

	"gorm.io/gorm"
	"mini_sistema/auth/models"
)

type UserSelector struct {
	db *gorm.DB
}

func NewUserSelector(db *gorm.DB) *UserSelector {
	return &UserSelector{db: db}
}

func (s *UserSelector) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.db.Preload("Group").Preload("Permissions").Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByID busca usuário pela chave primária
func (s *UserSelector) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
