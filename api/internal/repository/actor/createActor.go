package actor

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	convert "github.com/hxcuber/hollywood-api/api/internal/repository/convert/actor"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (i impl) CreateActor(ctx context.Context, actor models.Actor) (models.Actor, error) {
	insertColumns := boil.Blacklist()
	if actor.ActorID == -1 {
		insertColumns = boil.Blacklist("actor_id")
	}

	actorOrm := convert.ModelToOrm(actor)

	if err := (&actorOrm).Insert(ctx, i.dbConn, insertColumns); err != nil {
		return models.Actor{}, err
	}

	return convert.OrmToModel(actorOrm), nil
}
