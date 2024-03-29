package actors

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

func (i impl) GetActorByID(ctx context.Context, id int64) (models.Actor, error) {
	var actor models.Actor

	err := i.repo.DoInTx(context.Background(), func(ctx context.Context, txRepo repository.Registry) error {
		var err error
		actor, err = txRepo.Actor().GetActorByID(ctx, id)
		return err
	}, nil)

	if err != nil {
		return models.Actor{}, err
	}
	return actor, nil
}
