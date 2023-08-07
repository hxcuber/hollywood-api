package actors

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository"
)

type Controller interface {
	GetAllActors(ctx context.Context) (models.ActorSlice, error)
	GetActorByID(ctx context.Context, id int64) (models.Actor, error)
	CreateActor(ctx context.Context, actor models.Actor) (models.Actor, error)
}

func New(repo repository.Registry) Controller {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
