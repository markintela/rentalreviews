package storage

import "rentalreviewspt/services/types"

type Storage interface {
	Get(int) *types.User
}
