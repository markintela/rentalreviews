package types

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	From     string `json:"from"`
	AccessId string `json:"accessId"`
	Rent     *Rent  `json:"rents"`
}

func ValidateUser(u *User) bool {
	return true
}
