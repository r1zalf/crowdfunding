package user

import (
	"gorm.io/gorm"
)

// Contract
type Repository interface {
	Save(user User) (User, error)
}

// Like instance
type repository struct {
	db *gorm.DB
}

// Implemet
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		panic(err)
	}

	return user, nil

}
