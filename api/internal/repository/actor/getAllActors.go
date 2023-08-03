package actor

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
)

func (i impl) GetAllActors(ctx context.Context) (models.ActorSlice, error) {
	return models.Actors().All(ctx, i.dbConn)
}
