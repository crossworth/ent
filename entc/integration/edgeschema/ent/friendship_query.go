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
	"entgo.io/ent/entc/integration/edgeschema/ent/friendship"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// FriendshipQuery is the builder for querying Friendship entities.
type FriendshipQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	inters     []Interceptor
	predicates []predicate.Friendship
	withUser   *UserQuery
	withFriend *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FriendshipQuery builder.
func (fq *FriendshipQuery) Where(ps ...predicate.Friendship) *FriendshipQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FriendshipQuery) Limit(limit int) *FriendshipQuery {
	fq.limit = &limit
	return fq
}

// Offset to start from.
func (fq *FriendshipQuery) Offset(offset int) *FriendshipQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FriendshipQuery) Unique(unique bool) *FriendshipQuery {
	fq.unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FriendshipQuery) Order(o ...OrderFunc) *FriendshipQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryUser chains the current query on the "user" edge.
func (fq *FriendshipQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.UserTable, friendship.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFriend chains the current query on the "friend" edge.
func (fq *FriendshipQuery) QueryFriend() *UserQuery {
	query := (&UserClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(friendship.Table, friendship.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, friendship.FriendTable, friendship.FriendColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Friendship entity from the query.
// Returns a *NotFoundError when no Friendship was found.
func (fq *FriendshipQuery) First(ctx context.Context) (*Friendship, error) {
	nodes, err := fq.Limit(1).All(newQueryContext(ctx, TypeFriendship, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{friendship.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FriendshipQuery) FirstX(ctx context.Context) *Friendship {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Friendship ID from the query.
// Returns a *NotFoundError when no Friendship ID was found.
func (fq *FriendshipQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(newQueryContext(ctx, TypeFriendship, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{friendship.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FriendshipQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Friendship entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Friendship entity is found.
// Returns a *NotFoundError when no Friendship entities are found.
func (fq *FriendshipQuery) Only(ctx context.Context) (*Friendship, error) {
	nodes, err := fq.Limit(2).All(newQueryContext(ctx, TypeFriendship, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{friendship.Label}
	default:
		return nil, &NotSingularError{friendship.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FriendshipQuery) OnlyX(ctx context.Context) *Friendship {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Friendship ID in the query.
// Returns a *NotSingularError when more than one Friendship ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FriendshipQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(newQueryContext(ctx, TypeFriendship, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{friendship.Label}
	default:
		err = &NotSingularError{friendship.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FriendshipQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Friendships.
func (fq *FriendshipQuery) All(ctx context.Context) ([]*Friendship, error) {
	ctx = newQueryContext(ctx, TypeFriendship, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Friendship, *FriendshipQuery]()
	return withInterceptors[[]*Friendship](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FriendshipQuery) AllX(ctx context.Context) []*Friendship {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Friendship IDs.
func (fq *FriendshipQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = newQueryContext(ctx, TypeFriendship, "IDs")
	if err := fq.Select(friendship.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FriendshipQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FriendshipQuery) Count(ctx context.Context) (int, error) {
	ctx = newQueryContext(ctx, TypeFriendship, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FriendshipQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FriendshipQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FriendshipQuery) Exist(ctx context.Context) (bool, error) {
	ctx = newQueryContext(ctx, TypeFriendship, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FriendshipQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FriendshipQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FriendshipQuery) Clone() *FriendshipQuery {
	if fq == nil {
		return nil
	}
	return &FriendshipQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Friendship{}, fq.predicates...),
		withUser:   fq.withUser.Clone(),
		withFriend: fq.withFriend.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FriendshipQuery) WithUser(opts ...func(*UserQuery)) *FriendshipQuery {
	query := (&UserClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withUser = query
	return fq
}

// WithFriend tells the query-builder to eager-load the nodes that are connected to
// the "friend" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FriendshipQuery) WithFriend(opts ...func(*UserQuery)) *FriendshipQuery {
	query := (&UserClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withFriend = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Friendship.Query().
//		GroupBy(friendship.FieldWeight).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FriendshipQuery) GroupBy(field string, fields ...string) *FriendshipGroupBy {
	fq.fields = append([]string{field}, fields...)
	grbuild := &FriendshipGroupBy{build: fq}
	grbuild.flds = &fq.fields
	grbuild.label = friendship.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Weight int `json:"weight,omitempty"`
//	}
//
//	client.Friendship.Query().
//		Select(friendship.FieldWeight).
//		Scan(ctx, &v)
func (fq *FriendshipQuery) Select(fields ...string) *FriendshipSelect {
	fq.fields = append(fq.fields, fields...)
	sbuild := &FriendshipSelect{FriendshipQuery: fq}
	sbuild.label = friendship.Label
	sbuild.flds, sbuild.scan = &fq.fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FriendshipSelect configured with the given aggregations.
func (fq *FriendshipQuery) Aggregate(fns ...AggregateFunc) *FriendshipSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FriendshipQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.fields {
		if !friendship.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FriendshipQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Friendship, error) {
	var (
		nodes       = []*Friendship{}
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withUser != nil,
			fq.withFriend != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Friendship).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Friendship{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withUser; query != nil {
		if err := fq.loadUser(ctx, query, nodes, nil,
			func(n *Friendship, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withFriend; query != nil {
		if err := fq.loadFriend(ctx, query, nodes, nil,
			func(n *Friendship, e *User) { n.Edges.Friend = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FriendshipQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Friendship, init func(*Friendship), assign func(*Friendship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Friendship)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *FriendshipQuery) loadFriend(ctx context.Context, query *UserQuery, nodes []*Friendship, init func(*Friendship), assign func(*Friendship, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Friendship)
	for i := range nodes {
		fk := nodes[i].FriendID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "friend_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (fq *FriendshipQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FriendshipQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   friendship.Table,
			Columns: friendship.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: friendship.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, friendship.FieldID)
		for i := range fields {
			if fields[i] != friendship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FriendshipQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(friendship.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = friendship.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FriendshipGroupBy is the group-by builder for Friendship entities.
type FriendshipGroupBy struct {
	selector
	build *FriendshipQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FriendshipGroupBy) Aggregate(fns ...AggregateFunc) *FriendshipGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FriendshipGroupBy) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeFriendship, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendshipQuery, *FriendshipGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FriendshipGroupBy) sqlScan(ctx context.Context, root *FriendshipQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FriendshipSelect is the builder for selecting fields of Friendship entities.
type FriendshipSelect struct {
	*FriendshipQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FriendshipSelect) Aggregate(fns ...AggregateFunc) *FriendshipSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FriendshipSelect) Scan(ctx context.Context, v any) error {
	ctx = newQueryContext(ctx, TypeFriendship, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FriendshipQuery, *FriendshipSelect](ctx, fs.FriendshipQuery, fs, fs.inters, v)
}

func (fs *FriendshipSelect) sqlScan(ctx context.Context, root *FriendshipQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
