// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// NovelType is an object representing the database table.
type NovelType struct {
	ID   int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`

	R *novelTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L novelTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NovelTypeColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}

// Generated where

var NovelTypeWhere = struct {
	ID   whereHelperint
	Name whereHelperstring
}{
	ID:   whereHelperint{field: "\"ste\".\"novel_type\".\"id\""},
	Name: whereHelperstring{field: "\"ste\".\"novel_type\".\"name\""},
}

// NovelTypeRels is where relationship names are stored.
var NovelTypeRels = struct {
	NtypeNovels string
}{
	NtypeNovels: "NtypeNovels",
}

// novelTypeR is where relationships are stored.
type novelTypeR struct {
	NtypeNovels NovelSlice
}

// NewStruct creates a new relationship struct
func (*novelTypeR) NewStruct() *novelTypeR {
	return &novelTypeR{}
}

// novelTypeL is where Load methods for each relationship are stored.
type novelTypeL struct{}

var (
	novelTypeColumns               = []string{"id", "name"}
	novelTypeColumnsWithoutDefault = []string{"id", "name"}
	novelTypeColumnsWithDefault    = []string{}
	novelTypePrimaryKeyColumns     = []string{"id"}
)

type (
	// NovelTypeSlice is an alias for a slice of pointers to NovelType.
	// This should generally be used opposed to []NovelType.
	NovelTypeSlice []*NovelType
	// NovelTypeHook is the signature for custom NovelType hook methods
	NovelTypeHook func(context.Context, boil.ContextExecutor, *NovelType) error

	novelTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	novelTypeType                 = reflect.TypeOf(&NovelType{})
	novelTypeMapping              = queries.MakeStructMapping(novelTypeType)
	novelTypePrimaryKeyMapping, _ = queries.BindMapping(novelTypeType, novelTypeMapping, novelTypePrimaryKeyColumns)
	novelTypeInsertCacheMut       sync.RWMutex
	novelTypeInsertCache          = make(map[string]insertCache)
	novelTypeUpdateCacheMut       sync.RWMutex
	novelTypeUpdateCache          = make(map[string]updateCache)
	novelTypeUpsertCacheMut       sync.RWMutex
	novelTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var novelTypeBeforeInsertHooks []NovelTypeHook
var novelTypeBeforeUpdateHooks []NovelTypeHook
var novelTypeBeforeDeleteHooks []NovelTypeHook
var novelTypeBeforeUpsertHooks []NovelTypeHook

var novelTypeAfterInsertHooks []NovelTypeHook
var novelTypeAfterSelectHooks []NovelTypeHook
var novelTypeAfterUpdateHooks []NovelTypeHook
var novelTypeAfterDeleteHooks []NovelTypeHook
var novelTypeAfterUpsertHooks []NovelTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NovelType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NovelType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NovelType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NovelType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NovelType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NovelType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NovelType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NovelType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NovelType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range novelTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNovelTypeHook registers your hook function for all future operations.
func AddNovelTypeHook(hookPoint boil.HookPoint, novelTypeHook NovelTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		novelTypeBeforeInsertHooks = append(novelTypeBeforeInsertHooks, novelTypeHook)
	case boil.BeforeUpdateHook:
		novelTypeBeforeUpdateHooks = append(novelTypeBeforeUpdateHooks, novelTypeHook)
	case boil.BeforeDeleteHook:
		novelTypeBeforeDeleteHooks = append(novelTypeBeforeDeleteHooks, novelTypeHook)
	case boil.BeforeUpsertHook:
		novelTypeBeforeUpsertHooks = append(novelTypeBeforeUpsertHooks, novelTypeHook)
	case boil.AfterInsertHook:
		novelTypeAfterInsertHooks = append(novelTypeAfterInsertHooks, novelTypeHook)
	case boil.AfterSelectHook:
		novelTypeAfterSelectHooks = append(novelTypeAfterSelectHooks, novelTypeHook)
	case boil.AfterUpdateHook:
		novelTypeAfterUpdateHooks = append(novelTypeAfterUpdateHooks, novelTypeHook)
	case boil.AfterDeleteHook:
		novelTypeAfterDeleteHooks = append(novelTypeAfterDeleteHooks, novelTypeHook)
	case boil.AfterUpsertHook:
		novelTypeAfterUpsertHooks = append(novelTypeAfterUpsertHooks, novelTypeHook)
	}
}

// One returns a single novelType record from the query.
func (q novelTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*NovelType, error) {
	o := &NovelType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for novel_type")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all NovelType records from the query.
func (q novelTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (NovelTypeSlice, error) {
	var o []*NovelType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to NovelType slice")
	}

	if len(novelTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all NovelType records in the query.
func (q novelTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count novel_type rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q novelTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if novel_type exists")
	}

	return count > 0, nil
}

// NtypeNovels retrieves all the novel's Novels with an executor via ntype_id column.
func (o *NovelType) NtypeNovels(mods ...qm.QueryMod) novelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"ste\".\"novel\".\"ntype_id\"=?", o.ID),
	)

	query := Novels(queryMods...)
	queries.SetFrom(query.Query, "\"ste\".\"novel\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"ste\".\"novel\".*"})
	}

	return query
}

// LoadNtypeNovels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (novelTypeL) LoadNtypeNovels(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNovelType interface{}, mods queries.Applicator) error {
	var slice []*NovelType
	var object *NovelType

	if singular {
		object = maybeNovelType.(*NovelType)
	} else {
		slice = *maybeNovelType.(*[]*NovelType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &novelTypeR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &novelTypeR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`ste.novel`), qm.WhereIn(`ntype_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load novel")
	}

	var resultSlice []*Novel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice novel")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on novel")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for novel")
	}

	if len(novelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.NtypeNovels = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &novelR{}
			}
			foreign.R.Ntype = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.NtypeID) {
				local.R.NtypeNovels = append(local.R.NtypeNovels, foreign)
				if foreign.R == nil {
					foreign.R = &novelR{}
				}
				foreign.R.Ntype = local
				break
			}
		}
	}

	return nil
}

// AddNtypeNovels adds the given related objects to the existing relationships
// of the novel_type, optionally inserting them as new records.
// Appends related to o.R.NtypeNovels.
// Sets related.R.Ntype appropriately.
func (o *NovelType) AddNtypeNovels(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Novel) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.NtypeID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"ste\".\"novel\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"ntype_id"}),
				strmangle.WhereClause("\"", "\"", 2, novelPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.NtypeID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &novelTypeR{
			NtypeNovels: related,
		}
	} else {
		o.R.NtypeNovels = append(o.R.NtypeNovels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &novelR{
				Ntype: o,
			}
		} else {
			rel.R.Ntype = o
		}
	}
	return nil
}

// SetNtypeNovels removes all previously related items of the
// novel_type replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Ntype's NtypeNovels accordingly.
// Replaces o.R.NtypeNovels with related.
// Sets related.R.Ntype's NtypeNovels accordingly.
func (o *NovelType) SetNtypeNovels(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Novel) error {
	query := "update \"ste\".\"novel\" set \"ntype_id\" = null where \"ntype_id\" = $1"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.NtypeNovels {
			queries.SetScanner(&rel.NtypeID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Ntype = nil
		}

		o.R.NtypeNovels = nil
	}
	return o.AddNtypeNovels(ctx, exec, insert, related...)
}

// RemoveNtypeNovels relationships from objects passed in.
// Removes related items from R.NtypeNovels (uses pointer comparison, removal does not keep order)
// Sets related.R.Ntype.
func (o *NovelType) RemoveNtypeNovels(ctx context.Context, exec boil.ContextExecutor, related ...*Novel) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.NtypeID, nil)
		if rel.R != nil {
			rel.R.Ntype = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("ntype_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.NtypeNovels {
			if rel != ri {
				continue
			}

			ln := len(o.R.NtypeNovels)
			if ln > 1 && i < ln-1 {
				o.R.NtypeNovels[i] = o.R.NtypeNovels[ln-1]
			}
			o.R.NtypeNovels = o.R.NtypeNovels[:ln-1]
			break
		}
	}

	return nil
}

// NovelTypes retrieves all the records using an executor.
func NovelTypes(mods ...qm.QueryMod) novelTypeQuery {
	mods = append(mods, qm.From("\"ste\".\"novel_type\""))
	return novelTypeQuery{NewQuery(mods...)}
}

// FindNovelType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNovelType(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*NovelType, error) {
	novelTypeObj := &NovelType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ste\".\"novel_type\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, novelTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from novel_type")
	}

	return novelTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *NovelType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no novel_type provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(novelTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	novelTypeInsertCacheMut.RLock()
	cache, cached := novelTypeInsertCache[key]
	novelTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			novelTypeColumns,
			novelTypeColumnsWithDefault,
			novelTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(novelTypeType, novelTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(novelTypeType, novelTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ste\".\"novel_type\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ste\".\"novel_type\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into novel_type")
	}

	if !cached {
		novelTypeInsertCacheMut.Lock()
		novelTypeInsertCache[key] = cache
		novelTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the NovelType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *NovelType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	novelTypeUpdateCacheMut.RLock()
	cache, cached := novelTypeUpdateCache[key]
	novelTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			novelTypeColumns,
			novelTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update novel_type, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ste\".\"novel_type\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, novelTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(novelTypeType, novelTypeMapping, append(wl, novelTypePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update novel_type row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for novel_type")
	}

	if !cached {
		novelTypeUpdateCacheMut.Lock()
		novelTypeUpdateCache[key] = cache
		novelTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q novelTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for novel_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for novel_type")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NovelTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), novelTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ste\".\"novel_type\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, novelTypePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in novelType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all novelType")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *NovelType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no novel_type provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(novelTypeColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	novelTypeUpsertCacheMut.RLock()
	cache, cached := novelTypeUpsertCache[key]
	novelTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			novelTypeColumns,
			novelTypeColumnsWithDefault,
			novelTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			novelTypeColumns,
			novelTypePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert novel_type, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(novelTypePrimaryKeyColumns))
			copy(conflict, novelTypePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ste\".\"novel_type\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(novelTypeType, novelTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(novelTypeType, novelTypeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert novel_type")
	}

	if !cached {
		novelTypeUpsertCacheMut.Lock()
		novelTypeUpsertCache[key] = cache
		novelTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single NovelType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *NovelType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no NovelType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), novelTypePrimaryKeyMapping)
	sql := "DELETE FROM \"ste\".\"novel_type\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from novel_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for novel_type")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q novelTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no novelTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from novel_type")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for novel_type")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NovelTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(novelTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), novelTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ste\".\"novel_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, novelTypePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from novelType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for novel_type")
	}

	if len(novelTypeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *NovelType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNovelType(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NovelTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NovelTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), novelTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ste\".\"novel_type\".* FROM \"ste\".\"novel_type\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, novelTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NovelTypeSlice")
	}

	*o = slice

	return nil
}

// NovelTypeExists checks if the NovelType row exists.
func NovelTypeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ste\".\"novel_type\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if novel_type exists")
	}

	return exists, nil
}
