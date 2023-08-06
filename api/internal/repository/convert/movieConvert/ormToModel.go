package movieConvert

import (
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func MovieOrmToModel(m orm.Movie) models.Movie {
	return models.Movie{
		MovieID:     m.MovieID,
		Name:        m.Name,
		MainActorID: m.MainActorID,
		ReleaseDate: m.ReleaseDate,
	}
}

func MovieSliceOrmToModel(ms orm.MovieSlice) models.MovieSlice {
	var movies models.MovieSlice
	for _, m := range ms {
		movies = append(movies, MovieOrmToModel(*m))
	}
	return movies
}
