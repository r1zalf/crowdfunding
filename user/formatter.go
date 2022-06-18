package user

type UserFormatter struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormaterUser(user User, token string) UserFormatter {
	return UserFormatter{
		Id:         user.Id,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
}
