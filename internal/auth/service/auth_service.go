package service

import (
	"errors"

	"github.com/FANIMAN/chainforge/internal/auth/domain"
	"github.com/FANIMAN/chainforge/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepo
}

func NewAuthService(repo *repository.UserRepo) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(username, password, role string) (*domain.User, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &domain.User{
		ID:       username,
		Username: username,
		Password: string(hashed),
		Role:     role,
	}
	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) Login(username, password string) (*domain.User, error) {
	user, err := s.Repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
