package actors

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

func (i impl) GetAllActors(ctx context.Context) (models.ActorSlice, error) {
	var actors models.ActorSlice

	err := i.repo.DoInTx(context.Background(), func(ctx context.Context, txRepo repository.Registry) error {
		var err error
		actors, err = txRepo.Actor().GetAllActors(ctx)
		return err
	}, nil)

	if err != nil {
		return nil, err
	}
	return actors, nil
}
