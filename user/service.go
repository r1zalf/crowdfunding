package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	Register(userInput UserRegisterInput) (User, error)
}

type service struct {
	repository Repository
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
