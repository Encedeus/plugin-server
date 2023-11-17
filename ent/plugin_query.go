// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/pluginServer/ent/plugin"
	"github.com/Encedeus/pluginServer/ent/predicate"
	"github.com/Encedeus/pluginServer/ent/source"
	"github.com/Encedeus/pluginServer/ent/user"
	"github.com/google/uuid"
)

// PluginQuery is the builder for querying Plugin entities.
type PluginQuery struct {
	config
	ctx        *QueryContext
	order      []plugin.OrderOption
	inters     []Interceptor
	predicates []predicate.Plugin
	withOwner  *UserQuery
	withSource *SourceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PluginQuery builder.
func (pq *PluginQuery) Where(ps ...predicate.Plugin) *PluginQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PluginQuery) Limit(limit int) *PluginQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PluginQuery) Offset(offset int) *PluginQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PluginQuery) Unique(unique bool) *PluginQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PluginQuery) Order(o ...plugin.OrderOption) *PluginQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryOwner chains the current query on the "owner" edge.
func (pq *PluginQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(plugin.Table, plugin.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, plugin.OwnerTable, plugin.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySource chains the current query on the "source" edge.
func (pq *PluginQuery) QuerySource() *SourceQuery {
	query := (&SourceClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(plugin.Table, plugin.FieldID, selector),
			sqlgraph.To(source.Table, source.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, plugin.SourceTable, plugin.SourceColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Plugin entity from the query.
// Returns a *NotFoundError when no Plugin was found.
func (pq *PluginQuery) First(ctx context.Context) (*Plugin, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{plugin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PluginQuery) FirstX(ctx context.Context) *Plugin {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Plugin ID from the query.
// Returns a *NotFoundError when no Plugin ID was found.
func (pq *PluginQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{plugin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PluginQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Plugin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Plugin entity is found.
// Returns a *NotFoundError when no Plugin entities are found.
func (pq *PluginQuery) Only(ctx context.Context) (*Plugin, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{plugin.Label}
	default:
		return nil, &NotSingularError{plugin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PluginQuery) OnlyX(ctx context.Context) *Plugin {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Plugin ID in the query.
// Returns a *NotSingularError when more than one Plugin ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PluginQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{plugin.Label}
	default:
		err = &NotSingularError{plugin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PluginQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Plugins.
func (pq *PluginQuery) All(ctx context.Context) ([]*Plugin, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Plugin, *PluginQuery]()
	return withInterceptors[[]*Plugin](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PluginQuery) AllX(ctx context.Context) []*Plugin {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Plugin IDs.
func (pq *PluginQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(plugin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PluginQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PluginQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PluginQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PluginQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PluginQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PluginQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PluginQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PluginQuery) Clone() *PluginQuery {
	if pq == nil {
		return nil
	}
	return &PluginQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]plugin.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Plugin{}, pq.predicates...),
		withOwner:  pq.withOwner.Clone(),
		withSource: pq.withSource.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PluginQuery) WithOwner(opts ...func(*UserQuery)) *PluginQuery {
	query := (&UserClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withOwner = query
	return pq
}

// WithSource tells the query-builder to eager-load the nodes that are connected to
// the "source" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PluginQuery) WithSource(opts ...func(*SourceQuery)) *PluginQuery {
	query := (&SourceClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withSource = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Plugin.Query().
//		GroupBy(plugin.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PluginQuery) GroupBy(field string, fields ...string) *PluginGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PluginGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = plugin.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Plugin.Query().
//		Select(plugin.FieldName).
//		Scan(ctx, &v)
func (pq *PluginQuery) Select(fields ...string) *PluginSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PluginSelect{PluginQuery: pq}
	sbuild.label = plugin.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PluginSelect configured with the given aggregations.
func (pq *PluginQuery) Aggregate(fns ...AggregateFunc) *PluginSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PluginQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !plugin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PluginQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Plugin, error) {
	var (
		nodes       = []*Plugin{}
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withOwner != nil,
			pq.withSource != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Plugin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Plugin{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withOwner; query != nil {
		if err := pq.loadOwner(ctx, query, nodes, nil,
			func(n *Plugin, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withSource; query != nil {
		if err := pq.loadSource(ctx, query, nodes, nil,
			func(n *Plugin, e *Source) { n.Edges.Source = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PluginQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*Plugin, init func(*Plugin), assign func(*Plugin, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Plugin)
	for i := range nodes {
		fk := nodes[i].OwnerID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "owner_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *PluginQuery) loadSource(ctx context.Context, query *SourceQuery, nodes []*Plugin, init func(*Plugin), assign func(*Plugin, *Source)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Plugin)
	for i := range nodes {
		fk := nodes[i].SourceID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(source.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "source_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PluginQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PluginQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(plugin.Table, plugin.Columns, sqlgraph.NewFieldSpec(plugin.FieldID, field.TypeUUID))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plugin.FieldID)
		for i := range fields {
			if fields[i] != plugin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withOwner != nil {
			_spec.Node.AddColumnOnce(plugin.FieldOwnerID)
		}
		if pq.withSource != nil {
			_spec.Node.AddColumnOnce(plugin.FieldSourceID)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PluginQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(plugin.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = plugin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PluginGroupBy is the group-by builder for Plugin entities.
type PluginGroupBy struct {
	selector
	build *PluginQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PluginGroupBy) Aggregate(fns ...AggregateFunc) *PluginGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PluginGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PluginQuery, *PluginGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PluginGroupBy) sqlScan(ctx context.Context, root *PluginQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PluginSelect is the builder for selecting fields of Plugin entities.
type PluginSelect struct {
	*PluginQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PluginSelect) Aggregate(fns ...AggregateFunc) *PluginSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PluginSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PluginQuery, *PluginSelect](ctx, ps.PluginQuery, ps, ps.inters, v)
}

func (ps *PluginSelect) sqlScan(ctx context.Context, root *PluginQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}