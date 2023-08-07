package movie

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/convert/movie"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func (i impl) GetMovieByID(ctx context.Context, id int64) (models.Movie, error) {
	movieORMPointer, err := orm.FindMovie(ctx, i.dbConn, id)
	if err != nil {
		return models.Movie{}, err
	}

	return movie.MovieOrmToModel(*movieORMPointer), nil
}
