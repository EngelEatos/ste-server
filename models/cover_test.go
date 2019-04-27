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

func testCovers(t *testing.T) {
	t.Parallel()

	query := Covers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testCoversDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
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

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCoversQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Covers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCoversSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CoverSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testCoversExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := CoverExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Cover exists: %s", err)
	}
	if !e {
		t.Errorf("Expected CoverExists to return true, but got false.")
	}
}

func testCoversFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	coverFound, err := FindCover(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if coverFound == nil {
		t.Error("want a record, got nil")
	}
}

func testCoversBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Covers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testCoversOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Covers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testCoversAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	coverOne := &Cover{}
	coverTwo := &Cover{}
	if err = randomize.Struct(seed, coverOne, coverDBTypes, false, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}
	if err = randomize.Struct(seed, coverTwo, coverDBTypes, false, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = coverOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = coverTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Covers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testCoversCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	coverOne := &Cover{}
	coverTwo := &Cover{}
	if err = randomize.Struct(seed, coverOne, coverDBTypes, false, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}
	if err = randomize.Struct(seed, coverTwo, coverDBTypes, false, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = coverOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = coverTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func coverBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func coverAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Cover) error {
	*o = Cover{}
	return nil
}

func testCoversHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Cover{}
	o := &Cover{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, coverDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Cover object: %s", err)
	}

	AddCoverHook(boil.BeforeInsertHook, coverBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	coverBeforeInsertHooks = []CoverHook{}

	AddCoverHook(boil.AfterInsertHook, coverAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	coverAfterInsertHooks = []CoverHook{}

	AddCoverHook(boil.AfterSelectHook, coverAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	coverAfterSelectHooks = []CoverHook{}

	AddCoverHook(boil.BeforeUpdateHook, coverBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	coverBeforeUpdateHooks = []CoverHook{}

	AddCoverHook(boil.AfterUpdateHook, coverAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	coverAfterUpdateHooks = []CoverHook{}

	AddCoverHook(boil.BeforeDeleteHook, coverBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	coverBeforeDeleteHooks = []CoverHook{}

	AddCoverHook(boil.AfterDeleteHook, coverAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	coverAfterDeleteHooks = []CoverHook{}

	AddCoverHook(boil.BeforeUpsertHook, coverBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	coverBeforeUpsertHooks = []CoverHook{}

	AddCoverHook(boil.AfterUpsertHook, coverAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	coverAfterUpsertHooks = []CoverHook{}
}

func testCoversInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCoversInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(coverColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testCoverToManyNovels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Cover
	var b, c Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, novelDBTypes, false, novelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, novelDBTypes, false, novelColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.CoverID, a.ID)
	queries.Assign(&c.CoverID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Novels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.CoverID, b.CoverID) {
			bFound = true
		}
		if queries.Equal(v.CoverID, c.CoverID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := CoverSlice{&a}
	if err = a.L.LoadNovels(ctx, tx, false, (*[]*Cover)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Novels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Novels = nil
	if err = a.L.LoadNovels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Novels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testCoverToManyAddOpNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Cover
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, coverDBTypes, false, strmangle.SetComplement(coverPrimaryKeyColumns, coverColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Novel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, novelDBTypes, false, strmangle.SetComplement(novelPrimaryKeyColumns, novelColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Novel{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddNovels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.CoverID) {
			t.Error("foreign key was wrong value", a.ID, first.CoverID)
		}
		if !queries.Equal(a.ID, second.CoverID) {
			t.Error("foreign key was wrong value", a.ID, second.CoverID)
		}

		if first.R.Cover != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Cover != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Novels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Novels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Novels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testCoverToManySetOpNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Cover
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, coverDBTypes, false, strmangle.SetComplement(coverPrimaryKeyColumns, coverColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Novel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, novelDBTypes, false, strmangle.SetComplement(novelPrimaryKeyColumns, novelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetNovels(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Novels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetNovels(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Novels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CoverID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CoverID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.CoverID) {
		t.Error("foreign key was wrong value", a.ID, d.CoverID)
	}
	if !queries.Equal(a.ID, e.CoverID) {
		t.Error("foreign key was wrong value", a.ID, e.CoverID)
	}

	if b.R.Cover != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Cover != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Cover != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Cover != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Novels[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Novels[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testCoverToManyRemoveOpNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Cover
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, coverDBTypes, false, strmangle.SetComplement(coverPrimaryKeyColumns, coverColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Novel{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, novelDBTypes, false, strmangle.SetComplement(novelPrimaryKeyColumns, novelColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddNovels(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Novels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveNovels(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Novels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.CoverID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.CoverID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Cover != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Cover != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Cover != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Cover != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Novels) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Novels[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Novels[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testCoversReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
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

func testCoversReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := CoverSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testCoversSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Covers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	coverDBTypes = map[string]string{`ID`: `integer`, `URL`: `text`, `Downloaded`: `boolean`, `Path`: `text`}
	_            = bytes.MinRead
)

func testCoversUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(coverPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(coverColumns) == len(coverPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, coverDBTypes, true, coverPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testCoversSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(coverColumns) == len(coverPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Cover{}
	if err = randomize.Struct(seed, o, coverDBTypes, true, coverColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, coverDBTypes, true, coverPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(coverColumns, coverPrimaryKeyColumns) {
		fields = coverColumns
	} else {
		fields = strmangle.SetComplement(
			coverColumns,
			coverPrimaryKeyColumns,
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

	slice := CoverSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testCoversUpsert(t *testing.T) {
	t.Parallel()

	if len(coverColumns) == len(coverPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Cover{}
	if err = randomize.Struct(seed, &o, coverDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Cover: %s", err)
	}

	count, err := Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, coverDBTypes, false, coverPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Cover struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Cover: %s", err)
	}

	count, err = Covers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
