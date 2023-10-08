package types

import "time"

type Review struct {
	ID          int       `json:"id"`
	TypeReview  string    `json:"typeReview"`
	Description string    `json:"description"`
	Score       int       `json:"score"`
	DateReview  time.Time `json:"dateReview"`
}
