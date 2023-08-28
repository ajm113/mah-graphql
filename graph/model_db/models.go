package modeldb

import (
	"github.com/ajm113/mah-graphql/graph/model"
	"github.com/google/uuid"
)

type Meme struct {
	ID      uuid.UUID
	Image   string
	MovieID *uuid.UUID
}

func (m *Meme) ToDB(x *model.Meme) {
	m.ID = uuid.MustParse(x.ID)
	m.Image = x.Image

	if x.Movie != nil {
		id := uuid.MustParse(x.Movie.ID)
		m.MovieID = &id
	}
}

func (m *Meme) ToModel() *model.Meme {
	x := &model.Meme{
		ID:    m.ID.String(),
		Image: m.Image,
	}

	if m.MovieID != nil {
		x.Movie = &model.Movie{
			ID: m.MovieID.String(),
		}
	}

	return x
}

type Movie struct {
	ID       uuid.UUID
	Name     string
	Image    *string
	SitCount int
}

func (m *Movie) ToDB(x *model.Movie) {
	m.ID = uuid.MustParse(x.ID)
	m.Image = x.Image
	m.Name = x.Name
	m.SitCount = x.SitCount
}

func (m *Movie) ToModel() *model.Movie {
	x := &model.Movie{
		ID:       m.ID.String(),
		Image:    m.Image,
		Name:     m.Name,
		SitCount: m.SitCount,
	}
	return x
}

type Qoute struct {
	ID      uuid.UUID
	Text    string
	MovieID *uuid.UUID
}

func (m *Qoute) ToDB(x *model.Qoute) {
	m.ID = uuid.MustParse(x.ID)
	m.Text = x.Text

	if x.Movie != nil {
		id := uuid.MustParse(x.Movie.ID)
		m.MovieID = &id
	}
}

func (m *Qoute) ToModel() *model.Qoute {
	x := &model.Qoute{
		ID:   m.ID.String(),
		Text: m.Text,
	}

	if m.MovieID != nil {
		x.Movie = &model.Movie{
			ID: m.MovieID.String(),
		}
	}

	return x
}
