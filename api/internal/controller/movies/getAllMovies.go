package movies

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

func (i impl) GetAllMovies(ctx context.Context) (models.MovieSlice, error) {
	var movies models.MovieSlice

	err := i.repo.DoInTx(context.Background(), func(ctx context.Context, txRepo repository.Registry) error {
		var err error
		movies, err = txRepo.Movie().GetAllMovies(ctx)
		return err
	}, nil)

	if err != nil {
		return nil, err
	}
	return movies, nil
}
