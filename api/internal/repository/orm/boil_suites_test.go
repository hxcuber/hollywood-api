// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package orm

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Actors", testActors)
	t.Run("Movies", testMovies)
}

func TestDelete(t *testing.T) {
	t.Run("Actors", testActorsDelete)
	t.Run("Movies", testMoviesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Actors", testActorsQueryDeleteAll)
	t.Run("Movies", testMoviesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Actors", testActorsSliceDeleteAll)
	t.Run("Movies", testMoviesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Actors", testActorsExists)
	t.Run("Movies", testMoviesExists)
}

func TestFind(t *testing.T) {
	t.Run("Actors", testActorsFind)
	t.Run("Movies", testMoviesFind)
}

func TestBind(t *testing.T) {
	t.Run("Actors", testActorsBind)
	t.Run("Movies", testMoviesBind)
}

func TestOne(t *testing.T) {
	t.Run("Actors", testActorsOne)
	t.Run("Movies", testMoviesOne)
}

func TestAll(t *testing.T) {
	t.Run("Actors", testActorsAll)
	t.Run("Movies", testMoviesAll)
}

func TestCount(t *testing.T) {
	t.Run("Actors", testActorsCount)
	t.Run("Movies", testMoviesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Actors", testActorsHooks)
	t.Run("Movies", testMoviesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Actors", testActorsInsert)
	t.Run("Actors", testActorsInsertWhitelist)
	t.Run("Movies", testMoviesInsert)
	t.Run("Movies", testMoviesInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("MovieToActorUsingMainActor", testMovieToOneActorUsingMainActor)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ActorToMainActorMovies", testActorToManyMainActorMovies)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("MovieToActorUsingMainActorMovies", testMovieToOneSetOpActorUsingMainActor)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ActorToMainActorMovies", testActorToManyAddOpMainActorMovies)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Actors", testActorsReload)
	t.Run("Movies", testMoviesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Actors", testActorsReloadAll)
	t.Run("Movies", testMoviesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Actors", testActorsSelect)
	t.Run("Movies", testMoviesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Actors", testActorsUpdate)
	t.Run("Movies", testMoviesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Actors", testActorsSliceUpdateAll)
	t.Run("Movies", testMoviesSliceUpdateAll)
}
