package actorRepository

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/convert/actorConvert"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func (i impl) GetAllActors(ctx context.Context) (models.ActorSlice, error) {

	actorsORM, err := orm.Actors().All(ctx, i.dbConn)
	if err != nil {
		return nil, err
	}

	return actorConvert.ActorSliceOrmToModel(actorsORM), nil
}
