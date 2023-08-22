package model

type Movie struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Image  *string  `json:"image,omitempty"`
	Qoutes []*Qoute `json:"qoutes,omitempty"`
}
