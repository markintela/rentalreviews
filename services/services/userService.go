package services

import (
	context "context"
	types "rentalreviewspt/services/types"

	mongo "go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func (u *UserService) CreateUser(user *types.User) error {

	_, err := u.userCollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserService) GetUser(user *types.User) (*types.User, error) {

	return nil, nil
}

func (u *UserService) GetUserByName(name string) error {

	return nil
}

func (u *UserService) UpdateUser(user *types.User) error {

	return nil
}

func (u *UserService) GetAllUsers() ([]*types.User, error) {

	return nil, nil
}

func (u *UserService) DeleteUser(name string) error {

	return nil
}
