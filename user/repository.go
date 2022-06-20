package user

import (
	"gorm.io/gorm"
)

// Contract
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int64) (User, error)
}

// Like instance
type repository struct {
	db *gorm.DB
}

// FindById implements Repository
func (r *repository) FindById(id int64) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

// FindByEmail implements Repository
func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Save(user User) (User, error) {

	err := r.db.Create(&user)

	if err != nil {
		return user, err.Error
	}

	return user, nil
}

// Implemet
func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
