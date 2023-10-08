package types

type Location struct {
	ID       int     `json:"id"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Address  string  `json:"address"`
	Level    string  `json:"level"`
	Codepost string  `json:"codepost"`
}
