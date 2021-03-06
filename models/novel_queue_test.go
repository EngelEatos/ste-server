// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testNovelQueues(t *testing.T) {
	t.Parallel()

	query := NovelQueues()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNovelQueuesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
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

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNovelQueuesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NovelQueues().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNovelQueuesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NovelQueueSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNovelQueuesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NovelQueueExists(ctx, tx, o.NovelID, o.QueuedAt)
	if err != nil {
		t.Errorf("Unable to check if NovelQueue exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NovelQueueExists to return true, but got false.")
	}
}

func testNovelQueuesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	novelQueueFound, err := FindNovelQueue(ctx, tx, o.NovelID, o.QueuedAt)
	if err != nil {
		t.Error(err)
	}

	if novelQueueFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNovelQueuesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NovelQueues().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNovelQueuesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NovelQueues().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNovelQueuesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	novelQueueOne := &NovelQueue{}
	novelQueueTwo := &NovelQueue{}
	if err = randomize.Struct(seed, novelQueueOne, novelQueueDBTypes, false, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}
	if err = randomize.Struct(seed, novelQueueTwo, novelQueueDBTypes, false, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = novelQueueOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = novelQueueTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NovelQueues().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNovelQueuesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	novelQueueOne := &NovelQueue{}
	novelQueueTwo := &NovelQueue{}
	if err = randomize.Struct(seed, novelQueueOne, novelQueueDBTypes, false, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}
	if err = randomize.Struct(seed, novelQueueTwo, novelQueueDBTypes, false, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = novelQueueOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = novelQueueTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func novelQueueBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func novelQueueAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelQueue) error {
	*o = NovelQueue{}
	return nil
}

func testNovelQueuesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &NovelQueue{}
	o := &NovelQueue{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, novelQueueDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NovelQueue object: %s", err)
	}

	AddNovelQueueHook(boil.BeforeInsertHook, novelQueueBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	novelQueueBeforeInsertHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.AfterInsertHook, novelQueueAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	novelQueueAfterInsertHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.AfterSelectHook, novelQueueAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	novelQueueAfterSelectHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.BeforeUpdateHook, novelQueueBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	novelQueueBeforeUpdateHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.AfterUpdateHook, novelQueueAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	novelQueueAfterUpdateHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.BeforeDeleteHook, novelQueueBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	novelQueueBeforeDeleteHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.AfterDeleteHook, novelQueueAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	novelQueueAfterDeleteHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.BeforeUpsertHook, novelQueueBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	novelQueueBeforeUpsertHooks = []NovelQueueHook{}

	AddNovelQueueHook(boil.AfterUpsertHook, novelQueueAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	novelQueueAfterUpsertHooks = []NovelQueueHook{}
}

func testNovelQueuesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNovelQueuesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(novelQueueColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNovelQueueToOneNovelUsingNovel(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local NovelQueue
	var foreign Novel

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, novelQueueDBTypes, false, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, novelDBTypes, false, novelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Novel struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.NovelID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Novel().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := NovelQueueSlice{&local}
	if err = local.L.LoadNovel(ctx, tx, false, (*[]*NovelQueue)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Novel == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Novel = nil
	if err = local.L.LoadNovel(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Novel == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testNovelQueueToOneSetOpNovelUsingNovel(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NovelQueue
	var b, c Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, novelQueueDBTypes, false, strmangle.SetComplement(novelQueuePrimaryKeyColumns, novelQueueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, novelDBTypes, false, strmangle.SetComplement(novelPrimaryKeyColumns, novelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, novelDBTypes, false, strmangle.SetComplement(novelPrimaryKeyColumns, novelColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Novel{&b, &c} {
		err = a.SetNovel(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Novel != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.NovelQueues[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.NovelID != x.ID {
			t.Error("foreign key was wrong value", a.NovelID)
		}

		if exists, err := NovelQueueExists(ctx, tx, a.NovelID, a.QueuedAt); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testNovelQueuesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
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

func testNovelQueuesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NovelQueueSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNovelQueuesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NovelQueues().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	novelQueueDBTypes = map[string]string{`Finished`: `boolean`, `FinishedAt`: `date`, `NovelID`: `integer`, `QueuedAt`: `date`, `ScheduledAt`: `date`}
	_                 = bytes.MinRead
)

func testNovelQueuesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(novelQueuePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(novelQueueColumns) == len(novelQueuePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueuePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNovelQueuesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(novelQueueColumns) == len(novelQueuePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NovelQueue{}
	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, novelQueueDBTypes, true, novelQueuePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(novelQueueColumns, novelQueuePrimaryKeyColumns) {
		fields = novelQueueColumns
	} else {
		fields = strmangle.SetComplement(
			novelQueueColumns,
			novelQueuePrimaryKeyColumns,
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

	slice := NovelQueueSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNovelQueuesUpsert(t *testing.T) {
	t.Parallel()

	if len(novelQueueColumns) == len(novelQueuePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NovelQueue{}
	if err = randomize.Struct(seed, &o, novelQueueDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NovelQueue: %s", err)
	}

	count, err := NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, novelQueueDBTypes, false, novelQueuePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NovelQueue struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NovelQueue: %s", err)
	}

	count, err = NovelQueues().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
