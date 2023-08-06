package repository

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/repository/actorRepository"
	"github.com/hxcuber/hollywood-api/api/internal/repository/movieRepository"
	"time"

	"github.com/hxcuber/hollywood-api/api/internal/repository/systemRepository"
	"github.com/hxcuber/hollywood-api/api/pkg/db/pg"

	"github.com/cenkalti/backoff/v4"
	pkgerrors "github.com/pkg/errors"
)

// Registry is the registry of all the domain specific repositories and also provides transaction capabilities.
type Registry interface {
	// System returns the system repo
	System() systemRepository.Repository
	// Actor returns the actor repo
	Actor() actorRepository.Repository
	// Movie returns the movie repo
	Movie() movieRepository.Repository
	// DoInTx wraps operations within a db tx
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error
}

// New returns a new instance of Registry
func New(dbConn pg.BeginnerExecutor) Registry {
	return impl{
		dbConn: dbConn,
		system: systemRepository.New(dbConn),
		actor:  actorRepository.New(dbConn),
		movie:  movieRepository.New(dbConn),
	}
}

type impl struct {
	dbConn pg.BeginnerExecutor // Only used to start DB txns
	tx     pg.ContextExecutor  // Only used to keep track if txn has already been started to prevent devs from accidentally creating nested txns
	system systemRepository.Repository
	actor  actorRepository.Repository
	movie  movieRepository.Repository
}

// System returns the system repo
func (i impl) System() systemRepository.Repository {
	return i.system
}

func (i impl) Actor() actorRepository.Repository {
	return i.actor
}

func (i impl) Movie() movieRepository.Repository {
	return i.movie
}

// DoInTx wraps operations within a db tx
func (i impl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, txRepo Registry) error, overrideBackoffPolicy backoff.BackOff) error {
	if i.tx != nil {
		return pkgerrors.WithStack(errNestedTx)
	}

	if overrideBackoffPolicy == nil {
		overrideBackoffPolicy = pg.ExponentialBackOff(3, time.Minute)
	}

	return pg.TxWithBackOff(ctx, overrideBackoffPolicy, i.dbConn, func(tx pg.ContextExecutor) error {
		newI := impl{
			tx:     tx,
			system: systemRepository.New(tx),
			actor:  actorRepository.New(tx),
			movie:  movieRepository.New(tx),
		}
		return txFunc(ctx, newI)
	})
}
