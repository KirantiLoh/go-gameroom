package users

import (
	"database/sql"
	"errors"

	"github.com/kirantiloh/gameroom/pkg/auth"
	"github.com/kirantiloh/gameroom/pkg/database"
)

type UserService interface {
	RegisterUser(userDto *UserDto) error
	VerifyUserCredentials(loginDto *LoginDto) (*User, error)
	GetUser(email string) (*User, error)
	UpdateUser(user *User) error
}

type userService struct {
	repo UserRepository
}

func NewService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) RegisterUser(userDto *UserDto) error {
	hash, err := auth.HashPassword(userDto.Password)

	if err != nil {
		return err
	}

	userDto.Password = hash

	if err := s.repo.Create(userDto); err != nil {
		if database.IsUniqueConstraintError(err) {
			return errors.New("Email is already taken")
		}
		return errors.New("Database error: " + err.Error())
	}
	return nil
}

func (s *userService) VerifyUserCredentials(loginDto *LoginDto) (*User, error) {
	user, err := s.repo.Get(loginDto.Email)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	} else if err == sql.ErrNoRows {
    return nil, errors.New("User not found")
  }

	if !auth.VerifyPassword(loginDto.Password, user.Password) {
		return nil, errors.New("Email / password doesn't match")
	}

	return user, nil
}

func (s *userService) GetUser(email string) (*User, error) {
	return s.repo.Get(email)
}

func (s *userService) UpdateUser(user *User) error {
	return s.repo.Update(user)
}
