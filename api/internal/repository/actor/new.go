package actor

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/pkg/db/pg"
)

type Repository interface {
	GetAllActors(ctx context.Context) (models.ActorSlice, error)
	GetActorByID(ctx context.Context, id int64) (models.Actor, error)
	CreateActor(ctx context.Context, actor models.Actor) (models.Actor, error)
}

func New(dbConn pg.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn pg.ContextExecutor
}
