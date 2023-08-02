package model

import (
	"fmt"
	"strings"
	"time"
)

const (
	MovieIdWidth     = 5
	NameWidth        = 120
	MainActorIdWidth = 4
	ReleaseDateWidth = 11
)

type Movie struct {
	MovieId     int       `gorm:"column=movie_id;primaryKey"`
	MovieName   string    `gorm:"column=movie_name"`
	MainActorId int       `gorm:"column=main_actor_id"`
	ReleaseDate time.Time `gorm:"column=release_date"`
	Movie       Actor     `gorm:"foreignKey:MainActorId" gorm:"references:Actor_Id"`
}

type Movies []Movie

func (_ Movie) TableName() string {
	return "actors"
}

func (m Movie) String() string {
	return fmt.Sprintf("Movie ID: %-*d; Movie Name: %-*s; Main Actor ID: %-*d; Release Date: %-*s;\n", MovieIdWidth, m.MovieId, NameWidth, m.Name, MainActorIdWidth, m.MainActorId, ReleaseDateWidth, m.ReleaseDate.Format(time.DateOnly))
}

func (ms Movies) String() string {
	var sb strings.Builder
	for _, m := range ms {
		sb.WriteString(m.String())
	}
	return sb.String()
}
