package entity

type Apartment struct {
	Id       int     `json:"Id"`
	Phone    string  `json:"Phone"`
	Address  string  `json:"Address"`
	Category string  `json:"Category"`
	Rating   float64 `json:"Rating"`
}
