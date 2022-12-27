// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/schema/field"
)

// RelationshipInfoQuery is the builder for querying RelationshipInfo entities.
type RelationshipInfoQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.RelationshipInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RelationshipInfoQuery builder.
func (riq *RelationshipInfoQuery) Where(ps ...predicate.RelationshipInfo) *RelationshipInfoQuery {
	riq.predicates = append(riq.predicates, ps...)
	return riq
}

// Limit the number of records to be returned by this query.
func (riq *RelationshipInfoQuery) Limit(limit int) *RelationshipInfoQuery {
	riq.limit = &limit
	return riq
}

// Offset to start from.
func (riq *RelationshipInfoQuery) Offset(offset int) *RelationshipInfoQuery {
	riq.offset = &offset
	return riq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (riq *RelationshipInfoQuery) Unique(unique bool) *RelationshipInfoQuery {
	riq.unique = &unique
	return riq
}

// Order specifies how the records should be ordered.
func (riq *RelationshipInfoQuery) Order(o ...OrderFunc) *RelationshipInfoQuery {
	riq.order = append(riq.order, o...)
	return riq
}

// First returns the first RelationshipInfo entity from the query.
// Returns a *NotFoundError when no RelationshipInfo was found.
func (riq *RelationshipInfoQuery) First(ctx context.Context) (*RelationshipInfo, error) {
	nodes, err := riq.Limit(1).All(newQueryContext(ctx, TypeRelationshipInfo, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{relationshipinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (riq *RelationshipInfoQuery) FirstX(ctx context.Context) *RelationshipInfo {
	node, err := riq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RelationshipInfo ID from the query.
// Returns a *NotFoundError when no RelationshipInfo ID was found.
func (riq *RelationshipInfoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = riq.Limit(1).IDs(newQueryContext(ctx, TypeRelationshipInfo, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{relationshipinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (riq *RelationshipInfoQuery) FirstIDX(ctx context.Context) int {
	id, err := riq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RelationshipInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RelationshipInfo entity is found.
// Returns a *NotFoundError when no RelationshipInfo entities are found.
func (riq *RelationshipInfoQuery) Only(ctx context.Context) (*RelationshipInfo, error) {
	nodes, err := riq.Limit(2).All(newQueryContext(ctx, TypeRelationshipInfo, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{relationshipinfo.Label}
	default:
		return nil, &NotSingularError{relationshipinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (riq *RelationshipInfoQuery) OnlyX(ctx context.Context) *RelationshipInfo {
	node, err := riq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RelationshipInfo ID in the query.
// Returns a *NotSingularError when more than one RelationshipInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (riq *RelationshipInfoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = riq.Limit(2).IDs(newQueryContext(ctx, TypeRelationshipInfo, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{relationshipinfo.Label}
	default:
		err = &NotSingularError{relationshipinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (riq *RelationshipInfoQuery) OnlyIDX(ctx context.Context) int {
	id, err := riq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RelationshipInfos.
func (riq *RelationshipInfoQuery) All(ctx context.Context) ([]*RelationshipInfo, error) {
	ctx = newQueryContext(ctx, TypeRelationshipInfo, "All")
	if err := riq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RelationshipInfo, *RelationshipInfoQuery]()
	return withInterceptors[[]*RelationshipInfo](ctx, riq, qr, riq.inters)
}

// AllX is like All, but panics if an error occurs.
func (riq *RelationshipInfoQuery) AllX(ctx context.Context) []*RelationshipInfo {
	nodes, err := riq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RelationshipInfo IDs.
func (riq *RelationshipInfoQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeRelationshipInfo, "IDs")
	if err := riq.Select(relationshipinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (riq *RelationshipInfoQuery) IDsX(ctx context.Context) []int {
	ids, err := riq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (riq *RelationshipInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeRelationshipInfo, "Count")
	if err := riq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, riq, querierCount[*RelationshipInfoQuery](), riq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (riq *RelationshipInfoQuery) CountX(ctx context.Context) int {
	count, err := riq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (riq *RelationshipInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeRelationshipInfo, "Exist")
	switch _, err := riq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (riq *RelationshipInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := riq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RelationshipInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (riq *RelationshipInfoQuery) Clone() *RelationshipInfoQuery {
	if riq == nil {
		return nil
	}
	return &RelationshipInfoQuery{
		config:     riq.config,
		limit:      riq.limit,
		offset:     riq.offset,
		order:      append([]OrderFunc{}, riq.order...),
		inters:     append([]Interceptor{}, riq.inters...),
		predicates: append([]predicate.RelationshipInfo{}, riq.predicates...),
		// clone intermediate query.
		sql:    riq.sql.Clone(),
		path:   riq.path,
		unique: riq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RelationshipInfo.Query().
//		GroupBy(relationshipinfo.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (riq *RelationshipInfoQuery) GroupBy(field string, fields ...string) *RelationshipInfoGroupBy {
	riq.fields = append([]string{field}, fields...)
	grbuild := &RelationshipInfoGroupBy{build: riq}
	grbuild.flds = &riq.fields
	grbuild.label = relationshipinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.RelationshipInfo.Query().
//		Select(relationshipinfo.FieldText).
//		Scan(ctx, &v)
func (riq *RelationshipInfoQuery) Select(fields ...string) *RelationshipInfoSelect {
	riq.fields = append(riq.fields, fields...)
	sbuild := &RelationshipInfoSelect{RelationshipInfoQuery: riq}
	sbuild.label = relationshipinfo.Label
	sbuild.flds, sbuild.scan = &riq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RelationshipInfoSelect configured with the given aggregations.
func (riq *RelationshipInfoQuery) Aggregate(fns ...AggregateFunc) *RelationshipInfoSelect {
	return riq.Select().Aggregate(fns...)
}

func (riq *RelationshipInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range riq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, riq); err != nil {
				return err
			}
		}
	}
	for _, f := range riq.fields {
		if !relationshipinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if riq.path != nil {
		prev, err := riq.path(ctx)
		if err != nil {
			return err
		}
		riq.sql = prev
	}
	return nil
}

func (riq *RelationshipInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RelationshipInfo, error) {
	var (
		nodes = []*RelationshipInfo{}
		_spec = riq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RelationshipInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RelationshipInfo{config: riq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, riq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (riq *RelationshipInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := riq.querySpec()
	_spec.Node.Columns = riq.fields
	if len(riq.fields) > 0 {
		_spec.Unique = riq.unique != nil && *riq.unique
	}
	return sqlgraph.CountNodes(ctx, riq.driver, _spec)
}

func (riq *RelationshipInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relationshipinfo.Table,
			Columns: relationshipinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: relationshipinfo.FieldID,
			},
		},
		From:   riq.sql,
		Unique: true,
	}
	if unique := riq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := riq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, relationshipinfo.FieldID)
		for i := range fields {
			if fields[i] != relationshipinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := riq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := riq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := riq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := riq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (riq *RelationshipInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(riq.driver.Dialect())
	t1 := builder.Table(relationshipinfo.Table)
	columns := riq.fields
	if len(columns) == 0 {
		columns = relationshipinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if riq.sql != nil {
		selector = riq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if riq.unique != nil && *riq.unique {
		selector.Distinct()
	}
	for _, p := range riq.predicates {
		p(selector)
	}
	for _, p := range riq.order {
		p(selector)
	}
	if offset := riq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := riq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RelationshipInfoGroupBy is the group-by builder for RelationshipInfo entities.
type RelationshipInfoGroupBy struct {
	selector
	build *RelationshipInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rigb *RelationshipInfoGroupBy) Aggregate(fns ...AggregateFunc) *RelationshipInfoGroupBy {
	rigb.fns = append(rigb.fns, fns...)
	return rigb
}

// Scan applies the selector query and scans the result into the given value.
func (rigb *RelationshipInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeRelationshipInfo, "GroupBy")
	if err := rigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelationshipInfoQuery, *RelationshipInfoGroupBy](ctx, rigb.build, rigb, rigb.build.inters, v)
}

func (rigb *RelationshipInfoGroupBy) sqlScan(ctx context.Context, root *RelationshipInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rigb.fns))
	for _, fn := range rigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rigb.flds)+len(rigb.fns))
		for _, f := range *rigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RelationshipInfoSelect is the builder for selecting fields of RelationshipInfo entities.
type RelationshipInfoSelect struct {
	*RelationshipInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ris *RelationshipInfoSelect) Aggregate(fns ...AggregateFunc) *RelationshipInfoSelect {
	ris.fns = append(ris.fns, fns...)
	return ris
}

// Scan applies the selector query and scans the result into the given value.
func (ris *RelationshipInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeRelationshipInfo, "Select")
	if err := ris.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RelationshipInfoQuery, *RelationshipInfoSelect](ctx, ris.RelationshipInfoQuery, ris, ris.inters, v)
}

func (ris *RelationshipInfoSelect) sqlScan(ctx context.Context, root *RelationshipInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ris.fns))
	for _, fn := range ris.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ris.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ris.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
