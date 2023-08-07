package movies

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

type Controller interface {
	GetAllMovies(ctx context.Context) (models.MovieSlice, error)
	GetMovieByID(ctx context.Context, id int64) (models.Movie, error)
}

func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
