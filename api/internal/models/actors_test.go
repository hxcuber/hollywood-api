// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testActors(t *testing.T) {
	t.Parallel()

	query := Actors()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testActorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testActorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Actors().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testActorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ActorSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testActorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ActorExists(ctx, tx, o.ActorID)
	if err != nil {
		t.Errorf("Unable to check if Actor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ActorExists to return true, but got false.")
	}
}

func testActorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	actorFound, err := FindActor(ctx, tx, o.ActorID)
	if err != nil {
		t.Error(err)
	}

	if actorFound == nil {
		t.Error("want a record, got nil")
	}
}

func testActorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Actors().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testActorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Actors().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testActorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	actorOne := &Actor{}
	actorTwo := &Actor{}
	if err = randomize.Struct(seed, actorOne, actorDBTypes, false, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}
	if err = randomize.Struct(seed, actorTwo, actorDBTypes, false, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = actorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = actorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Actors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testActorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	actorOne := &Actor{}
	actorTwo := &Actor{}
	if err = randomize.Struct(seed, actorOne, actorDBTypes, false, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}
	if err = randomize.Struct(seed, actorTwo, actorDBTypes, false, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = actorOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = actorTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func actorBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func actorAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Actor) error {
	*o = Actor{}
	return nil
}

func testActorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Actor{}
	o := &Actor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, actorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Actor object: %s", err)
	}

	AddActorHook(boil.BeforeInsertHook, actorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	actorBeforeInsertHooks = []ActorHook{}

	AddActorHook(boil.AfterInsertHook, actorAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	actorAfterInsertHooks = []ActorHook{}

	AddActorHook(boil.AfterSelectHook, actorAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	actorAfterSelectHooks = []ActorHook{}

	AddActorHook(boil.BeforeUpdateHook, actorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	actorBeforeUpdateHooks = []ActorHook{}

	AddActorHook(boil.AfterUpdateHook, actorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	actorAfterUpdateHooks = []ActorHook{}

	AddActorHook(boil.BeforeDeleteHook, actorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	actorBeforeDeleteHooks = []ActorHook{}

	AddActorHook(boil.AfterDeleteHook, actorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	actorAfterDeleteHooks = []ActorHook{}

	AddActorHook(boil.BeforeUpsertHook, actorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	actorBeforeUpsertHooks = []ActorHook{}

	AddActorHook(boil.AfterUpsertHook, actorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	actorAfterUpsertHooks = []ActorHook{}
}

func testActorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testActorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(actorColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testActorToManyMainActorMovies(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Actor
	var b, c Movie

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, movieDBTypes, false, movieColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, movieDBTypes, false, movieColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.MainActorID = a.ActorID
	c.MainActorID = a.ActorID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.MainActorMovies().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.MainActorID == b.MainActorID {
			bFound = true
		}
		if v.MainActorID == c.MainActorID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ActorSlice{&a}
	if err = a.L.LoadMainActorMovies(ctx, tx, false, (*[]*Actor)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MainActorMovies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MainActorMovies = nil
	if err = a.L.LoadMainActorMovies(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MainActorMovies); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testActorToManyAddOpMainActorMovies(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Actor
	var b, c, d, e Movie

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, actorDBTypes, false, strmangle.SetComplement(actorPrimaryKeyColumns, actorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Movie{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, movieDBTypes, false, strmangle.SetComplement(moviePrimaryKeyColumns, movieColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Movie{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMainActorMovies(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ActorID != first.MainActorID {
			t.Error("foreign key was wrong value", a.ActorID, first.MainActorID)
		}
		if a.ActorID != second.MainActorID {
			t.Error("foreign key was wrong value", a.ActorID, second.MainActorID)
		}

		if first.R.MainActor != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.MainActor != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MainActorMovies[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MainActorMovies[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MainActorMovies().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testActorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testActorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ActorSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testActorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Actors().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	actorDBTypes = map[string]string{`ActorID`: `bigint`, `FirstName`: `character varying`, `LastName`: `character varying`, `Gender`: `character varying`, `Dob`: `date`, `Email`: `character varying`}
	_            = bytes.MinRead
)

func testActorsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(actorPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(actorAllColumns) == len(actorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, actorDBTypes, true, actorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testActorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(actorAllColumns) == len(actorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Actor{}
	if err = randomize.Struct(seed, o, actorDBTypes, true, actorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, actorDBTypes, true, actorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(actorAllColumns, actorPrimaryKeyColumns) {
		fields = actorAllColumns
	} else {
		fields = strmangle.SetComplement(
			actorAllColumns,
			actorPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ActorSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testActorsUpsert(t *testing.T) {
	t.Parallel()

	if len(actorAllColumns) == len(actorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Actor{}
	if err = randomize.Struct(seed, &o, actorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Actor: %s", err)
	}

	count, err := Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, actorDBTypes, false, actorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Actor struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Actor: %s", err)
	}

	count, err = Actors().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
