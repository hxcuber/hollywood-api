package models

import "time"

type Movie struct {
	MovieID     int64     `boil:"movie_id" json:"movie_id" toml:"movie_id" yaml:"movie_id"`
	Name        string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	MainActorID int64     `boil:"main_actor_id" json:"main_actor_id" toml:"main_actor_id" yaml:"main_actor_id"`
	ReleaseDate time.Time `boil:"release_date" json:"release_date" toml:"release_date" yaml:"release_date"`
}

type MovieSlice []Movie
