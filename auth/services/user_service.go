package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"mini_sistema/auth/models"
	"mini_sistema/auth/selectors"
)

type UserService struct {
	db           *gorm.DB
	userSelector *selectors.UserSelector
}

func NewUserService(db *gorm.DB, userSelector *selectors.UserSelector) *UserService {
	return &UserService{
		db:           db,
		userSelector: userSelector,
	}
}

// CreateSuperuser cria um usuário com is_admin = true e is_staff = true
func (s *UserService) CreateSuperuser(nome, email, password string) (*models.User, error) {
	// Verifica se já existe usuário com o e-mail informado
	existingUser, err := s.userSelector.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("já existe um usuário cadastrado com este e-mail")
	}

	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	superuser := &models.User{
		Nome:    nome,
		Email:   email,
		Senha:   string(hashedPassword),
		IsAdmin: true,
		IsStaff: true,
	}

	if err := s.db.Create(superuser).Error; err != nil {
		return nil, err
	}

	return superuser, nil
}
