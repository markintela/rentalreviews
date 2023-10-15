package services

import types "rentalreviewspt/services/types"

type IUserSevice interface {
	CreateUser(*types.User) error
	GetUserByName(*string) (*types.User, error)
	GetAllUsers() ([]*types.User, error)
	UpdateUser(*types.User) error
	DeleteUser(*string) error
}
