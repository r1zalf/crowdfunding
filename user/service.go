package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(userInput UserRegisterInput) (User, error)
	Login(userInput UserLoginInput) (User, error)
	GetUserById(id int64) (User, error)
}

type service struct {
	repository Repository
}

// GetUserById implements Service
func (s *service) GetUserById(id int64) (User, error) {
	userEntity, err := s.repository.FindById(id)
	if err != nil {
		return userEntity, err
	}

	if userEntity.Id == 0 {
		return userEntity, errors.New("User Not Found ")
	}
	return userEntity, nil
}

// Login implements Service
func (s *service) Login(userInput UserLoginInput) (User, error) {

	userEntity, err := s.repository.FindByEmail(userInput.Email)

	if err != nil {
		return userEntity, err
	}

	if userEntity.Id == 0 {
		return userEntity, errors.New("User Not Found ")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.PasswordHash), []byte(userInput.Password))
	if err != nil {
		return userEntity, err
	}

	return userEntity, nil
}

// Register implements Service
func (s *service) Register(userInput UserRegisterInput) (User, error) {

	passHash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.MinCost)
	user := User{
		Name:         userInput.Name,
		Occupation:   userInput.Occupation,
		Email:        userInput.Email,
		PasswordHash: string(passHash),
	}

	if err != nil {
		return user, err
	}

	user, err = s.repository.Save(user)

	if err != nil {
		return user, err
	}

	return user, nil

}

func NewService(r Repository) Service {
	return &service{r}
}
