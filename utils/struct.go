package utils

type Peer struct {
	Measurable bool     `json:"measurable"`
	Identifier string   `json:"identifier"`
	Rtt        *float64 `json:"rtt"`
	Loss       *float64 `json:"loss"`
	Min        *float64 `json:"min"`
	Max        *float64 `json:"max"`
	Mdev       *float64 `json:"mdev"`
	Responding *bool    `json:"responding"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
