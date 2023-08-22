// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Meme struct {
	ID    string `json:"id"`
	Image string `json:"image"`
	Movie *Movie `json:"movie,omitempty"`
}

type NewMeme struct {
	Image   string  `json:"image"`
	MovieID *string `json:"movieId,omitempty"`
}

type NewMovie struct {
	Name  string  `json:"name"`
	Image *string `json:"image,omitempty"`
}

type NewQoute struct {
	Text    string `json:"text"`
	MovieID string `json:"movieId"`
}

type Qoute struct {
	ID    string `json:"id"`
	Text  string `json:"text"`
	Movie *Movie `json:"movie,omitempty"`
}
