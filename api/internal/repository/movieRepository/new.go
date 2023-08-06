package movieRepository

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/pkg/db/pg"
)

type Repository interface {
	GetAllMovies(ctx context.Context) (models.MovieSlice, error)
	GetMovieByID(ctx context.Context, id int64) (models.Movie, error)
}

func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
