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

func testNovelTypes(t *testing.T) {
	t.Parallel()

	query := NovelTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testNovelTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
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

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNovelTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := NovelTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNovelTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NovelTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testNovelTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := NovelTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if NovelType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected NovelTypeExists to return true, but got false.")
	}
}

func testNovelTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	novelTypeFound, err := FindNovelType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if novelTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testNovelTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = NovelTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testNovelTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := NovelTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testNovelTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	novelTypeOne := &NovelType{}
	novelTypeTwo := &NovelType{}
	if err = randomize.Struct(seed, novelTypeOne, novelTypeDBTypes, false, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}
	if err = randomize.Struct(seed, novelTypeTwo, novelTypeDBTypes, false, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = novelTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = novelTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NovelTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testNovelTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	novelTypeOne := &NovelType{}
	novelTypeTwo := &NovelType{}
	if err = randomize.Struct(seed, novelTypeOne, novelTypeDBTypes, false, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}
	if err = randomize.Struct(seed, novelTypeTwo, novelTypeDBTypes, false, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = novelTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = novelTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func novelTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func novelTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *NovelType) error {
	*o = NovelType{}
	return nil
}

func testNovelTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &NovelType{}
	o := &NovelType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, novelTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize NovelType object: %s", err)
	}

	AddNovelTypeHook(boil.BeforeInsertHook, novelTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	novelTypeBeforeInsertHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.AfterInsertHook, novelTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	novelTypeAfterInsertHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.AfterSelectHook, novelTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	novelTypeAfterSelectHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.BeforeUpdateHook, novelTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	novelTypeBeforeUpdateHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.AfterUpdateHook, novelTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	novelTypeAfterUpdateHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.BeforeDeleteHook, novelTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	novelTypeBeforeDeleteHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.AfterDeleteHook, novelTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	novelTypeAfterDeleteHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.BeforeUpsertHook, novelTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	novelTypeBeforeUpsertHooks = []NovelTypeHook{}

	AddNovelTypeHook(boil.AfterUpsertHook, novelTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	novelTypeAfterUpsertHooks = []NovelTypeHook{}
}

func testNovelTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNovelTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(novelTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testNovelTypeToManyNtypeNovels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NovelType
	var b, c Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
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

	queries.Assign(&b.NtypeID, a.ID)
	queries.Assign(&c.NtypeID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.NtypeNovels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.NtypeID, b.NtypeID) {
			bFound = true
		}
		if queries.Equal(v.NtypeID, c.NtypeID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := NovelTypeSlice{&a}
	if err = a.L.LoadNtypeNovels(ctx, tx, false, (*[]*NovelType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NtypeNovels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.NtypeNovels = nil
	if err = a.L.LoadNtypeNovels(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.NtypeNovels); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testNovelTypeToManyAddOpNtypeNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NovelType
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, novelTypeDBTypes, false, strmangle.SetComplement(novelTypePrimaryKeyColumns, novelTypeColumnsWithoutDefault)...); err != nil {
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
		err = a.AddNtypeNovels(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.NtypeID) {
			t.Error("foreign key was wrong value", a.ID, first.NtypeID)
		}
		if !queries.Equal(a.ID, second.NtypeID) {
			t.Error("foreign key was wrong value", a.ID, second.NtypeID)
		}

		if first.R.Ntype != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Ntype != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.NtypeNovels[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.NtypeNovels[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.NtypeNovels().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testNovelTypeToManySetOpNtypeNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NovelType
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, novelTypeDBTypes, false, strmangle.SetComplement(novelTypePrimaryKeyColumns, novelTypeColumnsWithoutDefault)...); err != nil {
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

	err = a.SetNtypeNovels(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.NtypeNovels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetNtypeNovels(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.NtypeNovels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.NtypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.NtypeID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.NtypeID) {
		t.Error("foreign key was wrong value", a.ID, d.NtypeID)
	}
	if !queries.Equal(a.ID, e.NtypeID) {
		t.Error("foreign key was wrong value", a.ID, e.NtypeID)
	}

	if b.R.Ntype != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Ntype != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Ntype != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Ntype != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.NtypeNovels[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.NtypeNovels[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testNovelTypeToManyRemoveOpNtypeNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a NovelType
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, novelTypeDBTypes, false, strmangle.SetComplement(novelTypePrimaryKeyColumns, novelTypeColumnsWithoutDefault)...); err != nil {
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

	err = a.AddNtypeNovels(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.NtypeNovels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveNtypeNovels(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.NtypeNovels().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.NtypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.NtypeID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Ntype != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Ntype != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Ntype != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Ntype != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.NtypeNovels) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.NtypeNovels[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.NtypeNovels[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testNovelTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
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

func testNovelTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := NovelTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testNovelTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := NovelTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	novelTypeDBTypes = map[string]string{`ID`: `integer`, `Name`: `text`}
	_                = bytes.MinRead
)

func testNovelTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(novelTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(novelTypeColumns) == len(novelTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testNovelTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(novelTypeColumns) == len(novelTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &NovelType{}
	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, novelTypeDBTypes, true, novelTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(novelTypeColumns, novelTypePrimaryKeyColumns) {
		fields = novelTypeColumns
	} else {
		fields = strmangle.SetComplement(
			novelTypeColumns,
			novelTypePrimaryKeyColumns,
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

	slice := NovelTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testNovelTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(novelTypeColumns) == len(novelTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := NovelType{}
	if err = randomize.Struct(seed, &o, novelTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NovelType: %s", err)
	}

	count, err := NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, novelTypeDBTypes, false, novelTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize NovelType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert NovelType: %s", err)
	}

	count, err = NovelTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
