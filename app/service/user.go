package service

import (
	"go_myapp/app/model"
)

type UserService struct{}

func (UserService) GetUserIndex() ([]model.User, error) {
	var Users []model.User
	if err := db.Find(&Users).Error; err != nil {
		return nil, err
	}
	return Users, nil
}

func (UserService) CreateUser(User *model.User) error {
	if err := validate.Struct(User); err != nil {
		return err
	}
	if err := db.Create(User).Error; err != nil {
		return err
	}
	return nil
}
