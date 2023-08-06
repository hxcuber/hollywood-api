package models

import (
	"time"
)

type Actor struct {
	ActorID   int64     `boil:"actor_id" json:"actor_id,string" toml:"actor_id" yaml:"actor_id"`
	FirstName string    `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName  string    `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Gender    string    `boil:"gender" json:"gender" toml:"gender" yaml:"gender"`
	Dob       time.Time `boil:"dob" json:"dob" toml:"dob" yaml:"dob"`
	Email     string    `boil:"email" json:"email" toml:"email" yaml:"email"`
}

type ActorSlice []Actor
