package actor

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func (i impl) GetAllActors(ctx context.Context) (models.ActorSlice, error) {

	var actors models.ActorSlice
	err := orm.Actors().Bind(ctx, i.dbConn, &actors)
	if err != nil {
		return nil, err
	}

	return actors, nil
}
