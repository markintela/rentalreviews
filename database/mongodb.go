package storage

import "rentalreviewspt/services/types"

type MongoStorage struct {
}

func (s *MongoStorage) Get(id int) *types.User {

	return &types.User{
		ID:   1,
		Name: "Foo",
	}

}
