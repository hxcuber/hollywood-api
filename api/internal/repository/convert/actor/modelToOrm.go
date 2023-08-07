package actor

import (
	"github.com/hxcuber/hollywood-api/api/internal/models"
	"github.com/hxcuber/hollywood-api/api/internal/repository/orm"
)

func ModelToOrm(a models.Actor) orm.Actor {
	return orm.Actor{
		ActorID:   a.ActorID,
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Gender:    a.Gender,
		Dob:       a.Dob,
		Email:     a.Email,
	}
}

func SliceModelToOrm(as models.ActorSlice) orm.ActorSlice {
	var actors orm.ActorSlice
	for _, a := range as {
		actor := ModelToOrm(a)
		actors = append(actors, &actor)
	}
	return actors
}
