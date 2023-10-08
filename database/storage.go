package storage

import "rentalreviewspt/types"

type Storage interface {
	Get(int) *types.User
}
