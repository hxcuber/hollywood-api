package movies

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

func (i impl) GetMovieByID(ctx context.Context, id int64) (models.Movie, error) {
	var movie models.Movie

	err := i.repo.DoInTx(context.Background(), func(ctx context.Context, txRepo repository.Registry) error {
		var err error
		movie, err = txRepo.Movie().GetMovieByID(ctx, id)
		return err
	}, nil)

	if err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}
