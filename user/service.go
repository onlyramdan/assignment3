package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, file_loc string) (User, error)
	GetUserById(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}

	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(PasswordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No Found User on That Email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, file_loc string) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, nil
	}
	user.AvatarFileName = file_loc
	updateUser, err := s.repository.UpdateAvatar(user)

	if err != nil {
		return user, nil
	}
	return updateUser, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repository.FindById(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No Found User with That ID")
	}
	return user, nil
}
