package actors

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

func (i impl) CreateActor(ctx context.Context, actor models.Actor) (models.Actor, error) {
	// TODO
	var newActor models.Actor

	err := i.repo.DoInTx(context.Background(), func(ctx context.Context, txRepo repository.Registry) error {
		var err error
		newActor, err = txRepo.Actor().CreateActor(ctx, actor)
		return err
	}, nil)

	if err != nil {
		return models.Actor{}, err
	}
	return newActor, nil
}
