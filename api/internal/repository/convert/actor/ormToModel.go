package actor

import (
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func OrmToModel(a orm.Actor) models.Actor {
	return models.Actor{
		ActorID:   a.ActorID,
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Gender:    a.Gender,
		Dob:       a.Dob,
		Email:     a.Email,
	}
}

func SliceOrmToModel(as orm.ActorSlice) models.ActorSlice {
	var actors models.ActorSlice
	for _, a := range as {
		actors = append(actors, OrmToModel(*a))
	}
	return actors
}
