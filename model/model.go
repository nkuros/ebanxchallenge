package model

type Event struct {
	Type        string  `json:"type"`
	Amount      int     `json:"amount"`
	Origin      string  `json:"origin"`
	Destination *string `json:"destination"`
}
