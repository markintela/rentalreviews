package storage

import "rentalreviewspt/types"

type MemoryStorage struct {
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Get(id int) *types.User {

	return &types.User{
		ID:   3,
		Name: "Foo",
	}

}
