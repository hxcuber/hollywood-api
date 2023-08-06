package movieRepository

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/convert/movieConvert"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func (i impl) GetAllMovies(ctx context.Context) (models.MovieSlice, error) {
	moviesORM, err := orm.Movies().All(ctx, i.dbConn)
	if err != nil {
		return nil, err
	}

	return movieConvert.MovieSliceOrmToModel(moviesORM), nil
}
