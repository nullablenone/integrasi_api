package user

import "gorm.io/gorm"

type Repository interface {
	SaveUser(user *User) error
	GetAllUsers() ([]User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) SaveUser(user *User) error {
	return r.DB.Save(user).Error
}

func (r *repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.DB.Preload("Address").Preload("Company").Find(&users).Error
	return users, err
}
