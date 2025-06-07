package utils

type Peer struct {
	Measurable bool     `json:"measurable"`
	Identifier string   `json:"identifier"`
	Rtt        *int64   `json:"rtt"`
	Loss       *float64 `json:"loss"`
}
