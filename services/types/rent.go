package types

import "time"

type Rent struct {
	ID           int        `json:"id"`
	Resource     string     `json:"resource"`
	TypeImmobile string     `json:"typeImmobile"`
	DateBegin    time.Time  `json:"dateBegin"`
	DateEnd      time.Time  `json:"dateEnd"`
	Price        float64    `json:"price"`
	Review       *Review    `json:"reviews"`
	Comment      string     `json:"comment"`
	Location     Location   `json:"location"`
	Signature    *Signature `json:"signatures"`
}
