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

func testChapters(t *testing.T) {
	t.Parallel()

	query := Chapters()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testChaptersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
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

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChaptersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Chapters().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChaptersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChapterSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChaptersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ChapterExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Chapter exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChapterExists to return true, but got false.")
	}
}

func testChaptersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	chapterFound, err := FindChapter(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if chapterFound == nil {
		t.Error("want a record, got nil")
	}
}

func testChaptersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Chapters().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testChaptersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Chapters().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChaptersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	chapterOne := &Chapter{}
	chapterTwo := &Chapter{}
	if err = randomize.Struct(seed, chapterOne, chapterDBTypes, false, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}
	if err = randomize.Struct(seed, chapterTwo, chapterDBTypes, false, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = chapterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = chapterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Chapters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChaptersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	chapterOne := &Chapter{}
	chapterTwo := &Chapter{}
	if err = randomize.Struct(seed, chapterOne, chapterDBTypes, false, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}
	if err = randomize.Struct(seed, chapterTwo, chapterDBTypes, false, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = chapterOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = chapterTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func chapterBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func chapterAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Chapter) error {
	*o = Chapter{}
	return nil
}

func testChaptersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Chapter{}
	o := &Chapter{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, chapterDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Chapter object: %s", err)
	}

	AddChapterHook(boil.BeforeInsertHook, chapterBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	chapterBeforeInsertHooks = []ChapterHook{}

	AddChapterHook(boil.AfterInsertHook, chapterAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	chapterAfterInsertHooks = []ChapterHook{}

	AddChapterHook(boil.AfterSelectHook, chapterAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	chapterAfterSelectHooks = []ChapterHook{}

	AddChapterHook(boil.BeforeUpdateHook, chapterBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	chapterBeforeUpdateHooks = []ChapterHook{}

	AddChapterHook(boil.AfterUpdateHook, chapterAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	chapterAfterUpdateHooks = []ChapterHook{}

	AddChapterHook(boil.BeforeDeleteHook, chapterBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	chapterBeforeDeleteHooks = []ChapterHook{}

	AddChapterHook(boil.AfterDeleteHook, chapterAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	chapterAfterDeleteHooks = []ChapterHook{}

	AddChapterHook(boil.BeforeUpsertHook, chapterBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	chapterBeforeUpsertHooks = []ChapterHook{}

	AddChapterHook(boil.AfterUpsertHook, chapterAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	chapterAfterUpsertHooks = []ChapterHook{}
}

func testChaptersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChaptersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(chapterColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChapterToManyNovels(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
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

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into \"ste\".\"chapter_of_novel\" (\"chapter_id\", \"novel_id\") values ($1, $2)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into \"ste\".\"chapter_of_novel\" (\"chapter_id\", \"novel_id\") values ($1, $2)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	check, err := a.Novels().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChapterSlice{&a}
	if err = a.L.LoadNovels(ctx, tx, false, (*[]*Chapter)(&slice), nil); err != nil {
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

func testChapterToManyChapterQueues(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c ChapterQueue

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, chapterQueueDBTypes, false, chapterQueueColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, chapterQueueDBTypes, false, chapterQueueColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ChapterID = a.ID
	c.ChapterID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ChapterQueues().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ChapterID == b.ChapterID {
			bFound = true
		}
		if v.ChapterID == c.ChapterID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := ChapterSlice{&a}
	if err = a.L.LoadChapterQueues(ctx, tx, false, (*[]*Chapter)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChapterQueues); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ChapterQueues = nil
	if err = a.L.LoadChapterQueues(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ChapterQueues); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testChapterToManyAddOpNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
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

		if first.R.Chapters[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Chapters[0] != &a {
			t.Error("relationship was not added properly to the slice")
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

func testChapterToManySetOpNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
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

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Chapters) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Chapters) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Chapters[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Chapters[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Novels[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Novels[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testChapterToManyRemoveOpNovels(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c, d, e Novel

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
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

	if len(b.R.Chapters) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Chapters) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Chapters[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Chapters[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
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

func testChapterToManyAddOpChapterQueues(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c, d, e ChapterQueue

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*ChapterQueue{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, chapterQueueDBTypes, false, strmangle.SetComplement(chapterQueuePrimaryKeyColumns, chapterQueueColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*ChapterQueue{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddChapterQueues(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ChapterID {
			t.Error("foreign key was wrong value", a.ID, first.ChapterID)
		}
		if a.ID != second.ChapterID {
			t.Error("foreign key was wrong value", a.ID, second.ChapterID)
		}

		if first.R.Chapter != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Chapter != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ChapterQueues[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ChapterQueues[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ChapterQueues().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testChapterToOneGroupUsingGroup(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Chapter
	var foreign Group

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.GroupID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Group().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := ChapterSlice{&local}
	if err = local.L.LoadGroup(ctx, tx, false, (*[]*Chapter)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Group == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Group = nil
	if err = local.L.LoadGroup(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Group == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testChapterToOneSetOpGroupUsingGroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b, c Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Group{&b, &c} {
		err = a.SetGroup(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Group != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Chapters[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.GroupID, x.ID) {
			t.Error("foreign key was wrong value", a.GroupID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.GroupID))
		reflect.Indirect(reflect.ValueOf(&a.GroupID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.GroupID, x.ID) {
			t.Error("foreign key was wrong value", a.GroupID, x.ID)
		}
	}
}

func testChapterToOneRemoveOpGroupUsingGroup(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Chapter
	var b Group

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, chapterDBTypes, false, strmangle.SetComplement(chapterPrimaryKeyColumns, chapterColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, groupDBTypes, false, strmangle.SetComplement(groupPrimaryKeyColumns, groupColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetGroup(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveGroup(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Group().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Group != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.GroupID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Chapters) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testChaptersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
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

func testChaptersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ChapterSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testChaptersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Chapters().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	chapterDBTypes = map[string]string{`Downloaded`: `boolean`, `GroupID`: `integer`, `ID`: `integer`, `Idx`: `integer`, `Part`: `integer`, `Title`: `text`, `URL`: `text`}
	_              = bytes.MinRead
)

func testChaptersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(chapterPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(chapterColumns) == len(chapterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testChaptersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(chapterColumns) == len(chapterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Chapter{}
	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, chapterDBTypes, true, chapterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(chapterColumns, chapterPrimaryKeyColumns) {
		fields = chapterColumns
	} else {
		fields = strmangle.SetComplement(
			chapterColumns,
			chapterPrimaryKeyColumns,
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

	slice := ChapterSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testChaptersUpsert(t *testing.T) {
	t.Parallel()

	if len(chapterColumns) == len(chapterPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Chapter{}
	if err = randomize.Struct(seed, &o, chapterDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Chapter: %s", err)
	}

	count, err := Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, chapterDBTypes, false, chapterPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Chapter struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Chapter: %s", err)
	}

	count, err = Chapters().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
