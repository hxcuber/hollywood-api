package actorRepository

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/convert/actorConvert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateActor(ctx context.Context, actor models.Actor) (models.Actor, error) {
	insertColumns := boil.Blacklist()
	if actor.ActorID == -1 {
		insertColumns = boil.Blacklist("actor_id")
	}

	actorOrm := actorConvert.ActorModelToOrm(actor)

	if err := (&actorOrm).Insert(ctx, i.dbConn, insertColumns); err != nil {
		return models.Actor{}, err
	}

	return actorConvert.ActorOrmToModel(actorOrm), nil
}
