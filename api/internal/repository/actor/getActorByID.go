package actor

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/convert/actor"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func (i impl) GetActorByID(ctx context.Context, id int64) (models.Actor, error) {
	actorORMPointer, err := orm.FindActor(ctx, i.dbConn, id)
	if err != nil {
		return models.Actor{}, err
	}

	return actor.OrmToModel(*actorORMPointer), nil
}
