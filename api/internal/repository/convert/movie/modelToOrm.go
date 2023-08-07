package movie

import (
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func ModelToOrm(m models.Movie) orm.Movie {
	return orm.Movie{
		MovieID:     m.MovieID,
		Name:        m.Name,
		MainActorID: m.MainActorID,
		ReleaseDate: m.ReleaseDate,
	}
}

func SliceModelToOrm(ms models.MovieSlice) orm.MovieSlice {
	var movies orm.MovieSlice
	for _, m := range ms {
		movie := ModelToOrm(m)
		movies = append(movies, &movie)
	}
	return movies
}
