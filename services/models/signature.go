package types

import "time"

type Signature struct {
	SignedById    int       `json:"signedById"`
	SignedByName  int       `json:"signedByName"`
	DateSignature time.Time `json:"dateSignature"`
	Comment       string    `json:"comment"`
}
