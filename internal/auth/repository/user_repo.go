package repository

import (
	"errors"
	"github.com/FANIMAN/chainforge/internal/auth/domain"
)

type UserRepo struct {
	users map[string]*domain.User
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[string]*domain.User),
	}
}

func (r *UserRepo) Create(user *domain.User) error {
	if _, exists := r.users[user.Username]; exists {
		return errors.New("username already exists")
	}
	r.users[user.Username] = user
	return nil
}

func (r *UserRepo) GetByUsername(username string) (*domain.User, error) {
	user, ok := r.users[username]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
