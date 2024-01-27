// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Encedeus/pluginServer/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Encedeus/pluginServer/ent/plugin"
	"github.com/Encedeus/pluginServer/ent/publication"
	"github.com/Encedeus/pluginServer/ent/source"
	"github.com/Encedeus/pluginServer/ent/user"
	"github.com/Encedeus/pluginServer/ent/verificationsession"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Plugin is the client for interacting with the Plugin builders.
	Plugin *PluginClient
	// Publication is the client for interacting with the Publication builders.
	Publication *PublicationClient
	// Source is the client for interacting with the Source builders.
	Source *SourceClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// VerificationSession is the client for interacting with the VerificationSession builders.
	VerificationSession *VerificationSessionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Plugin = NewPluginClient(c.config)
	c.Publication = NewPublicationClient(c.config)
	c.Source = NewSourceClient(c.config)
	c.User = NewUserClient(c.config)
	c.VerificationSession = NewVerificationSessionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Plugin:              NewPluginClient(cfg),
		Publication:         NewPublicationClient(cfg),
		Source:              NewSourceClient(cfg),
		User:                NewUserClient(cfg),
		VerificationSession: NewVerificationSessionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Plugin:              NewPluginClient(cfg),
		Publication:         NewPublicationClient(cfg),
		Source:              NewSourceClient(cfg),
		User:                NewUserClient(cfg),
		VerificationSession: NewVerificationSessionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Plugin.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Plugin.Use(hooks...)
	c.Publication.Use(hooks...)
	c.Source.Use(hooks...)
	c.User.Use(hooks...)
	c.VerificationSession.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Plugin.Intercept(interceptors...)
	c.Publication.Intercept(interceptors...)
	c.Source.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.VerificationSession.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PluginMutation:
		return c.Plugin.mutate(ctx, m)
	case *PublicationMutation:
		return c.Publication.mutate(ctx, m)
	case *SourceMutation:
		return c.Source.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *VerificationSessionMutation:
		return c.VerificationSession.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PluginClient is a client for the Plugin schema.
type PluginClient struct {
	config
}

// NewPluginClient returns a client for the Plugin from the given config.
func NewPluginClient(c config) *PluginClient {
	return &PluginClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `plugin.Hooks(f(g(h())))`.
func (c *PluginClient) Use(hooks ...Hook) {
	c.hooks.Plugin = append(c.hooks.Plugin, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `plugin.Intercept(f(g(h())))`.
func (c *PluginClient) Intercept(interceptors ...Interceptor) {
	c.inters.Plugin = append(c.inters.Plugin, interceptors...)
}

// Create returns a builder for creating a Plugin entity.
func (c *PluginClient) Create() *PluginCreate {
	mutation := newPluginMutation(c.config, OpCreate)
	return &PluginCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Plugin entities.
func (c *PluginClient) CreateBulk(builders ...*PluginCreate) *PluginCreateBulk {
	return &PluginCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PluginClient) MapCreateBulk(slice any, setFunc func(*PluginCreate, int)) *PluginCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PluginCreateBulk{err: fmt.Errorf("calling to PluginClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PluginCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PluginCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Plugin.
func (c *PluginClient) Update() *PluginUpdate {
	mutation := newPluginMutation(c.config, OpUpdate)
	return &PluginUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PluginClient) UpdateOne(pl *Plugin) *PluginUpdateOne {
	mutation := newPluginMutation(c.config, OpUpdateOne, withPlugin(pl))
	return &PluginUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PluginClient) UpdateOneID(id uuid.UUID) *PluginUpdateOne {
	mutation := newPluginMutation(c.config, OpUpdateOne, withPluginID(id))
	return &PluginUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Plugin.
func (c *PluginClient) Delete() *PluginDelete {
	mutation := newPluginMutation(c.config, OpDelete)
	return &PluginDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PluginClient) DeleteOne(pl *Plugin) *PluginDeleteOne {
	return c.DeleteOneID(pl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PluginClient) DeleteOneID(id uuid.UUID) *PluginDeleteOne {
	builder := c.Delete().Where(plugin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PluginDeleteOne{builder}
}

// Query returns a query builder for Plugin.
func (c *PluginClient) Query() *PluginQuery {
	return &PluginQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePlugin},
		inters: c.Interceptors(),
	}
}

// Get returns a Plugin entity by its id.
func (c *PluginClient) Get(ctx context.Context, id uuid.UUID) (*Plugin, error) {
	return c.Query().Where(plugin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PluginClient) GetX(ctx context.Context, id uuid.UUID) *Plugin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Plugin.
func (c *PluginClient) QueryOwner(pl *Plugin) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plugin.Table, plugin.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, plugin.OwnerTable, plugin.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySource queries the source edge of a Plugin.
func (c *PluginClient) QuerySource(pl *Plugin) *SourceQuery {
	query := (&SourceClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plugin.Table, plugin.FieldID, id),
			sqlgraph.To(source.Table, source.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, plugin.SourceTable, plugin.SourceColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPublications queries the publications edge of a Plugin.
func (c *PluginClient) QueryPublications(pl *Plugin) *PublicationQuery {
	query := (&PublicationClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(plugin.Table, plugin.FieldID, id),
			sqlgraph.To(publication.Table, publication.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, plugin.PublicationsTable, plugin.PublicationsColumn),
		)
		fromV = sqlgraph.Neighbors(pl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PluginClient) Hooks() []Hook {
	return c.hooks.Plugin
}

// Interceptors returns the client interceptors.
func (c *PluginClient) Interceptors() []Interceptor {
	return c.inters.Plugin
}

func (c *PluginClient) mutate(ctx context.Context, m *PluginMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PluginCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PluginUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PluginUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PluginDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Plugin mutation op: %q", m.Op())
	}
}

// PublicationClient is a client for the Publication schema.
type PublicationClient struct {
	config
}

// NewPublicationClient returns a client for the Publication from the given config.
func NewPublicationClient(c config) *PublicationClient {
	return &PublicationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `publication.Hooks(f(g(h())))`.
func (c *PublicationClient) Use(hooks ...Hook) {
	c.hooks.Publication = append(c.hooks.Publication, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `publication.Intercept(f(g(h())))`.
func (c *PublicationClient) Intercept(interceptors ...Interceptor) {
	c.inters.Publication = append(c.inters.Publication, interceptors...)
}

// Create returns a builder for creating a Publication entity.
func (c *PublicationClient) Create() *PublicationCreate {
	mutation := newPublicationMutation(c.config, OpCreate)
	return &PublicationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Publication entities.
func (c *PublicationClient) CreateBulk(builders ...*PublicationCreate) *PublicationCreateBulk {
	return &PublicationCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PublicationClient) MapCreateBulk(slice any, setFunc func(*PublicationCreate, int)) *PublicationCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PublicationCreateBulk{err: fmt.Errorf("calling to PublicationClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PublicationCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PublicationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Publication.
func (c *PublicationClient) Update() *PublicationUpdate {
	mutation := newPublicationMutation(c.config, OpUpdate)
	return &PublicationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PublicationClient) UpdateOne(pu *Publication) *PublicationUpdateOne {
	mutation := newPublicationMutation(c.config, OpUpdateOne, withPublication(pu))
	return &PublicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PublicationClient) UpdateOneID(id int) *PublicationUpdateOne {
	mutation := newPublicationMutation(c.config, OpUpdateOne, withPublicationID(id))
	return &PublicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Publication.
func (c *PublicationClient) Delete() *PublicationDelete {
	mutation := newPublicationMutation(c.config, OpDelete)
	return &PublicationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PublicationClient) DeleteOne(pu *Publication) *PublicationDeleteOne {
	return c.DeleteOneID(pu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PublicationClient) DeleteOneID(id int) *PublicationDeleteOne {
	builder := c.Delete().Where(publication.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PublicationDeleteOne{builder}
}

// Query returns a query builder for Publication.
func (c *PublicationClient) Query() *PublicationQuery {
	return &PublicationQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePublication},
		inters: c.Interceptors(),
	}
}

// Get returns a Publication entity by its id.
func (c *PublicationClient) Get(ctx context.Context, id int) (*Publication, error) {
	return c.Query().Where(publication.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PublicationClient) GetX(ctx context.Context, id int) *Publication {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlugin queries the plugin edge of a Publication.
func (c *PublicationClient) QueryPlugin(pu *Publication) *PluginQuery {
	query := (&PluginClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(publication.Table, publication.FieldID, id),
			sqlgraph.To(plugin.Table, plugin.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, publication.PluginTable, publication.PluginColumn),
		)
		fromV = sqlgraph.Neighbors(pu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PublicationClient) Hooks() []Hook {
	return c.hooks.Publication
}

// Interceptors returns the client interceptors.
func (c *PublicationClient) Interceptors() []Interceptor {
	return c.inters.Publication
}

func (c *PublicationClient) mutate(ctx context.Context, m *PublicationMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PublicationCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PublicationUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PublicationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PublicationDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Publication mutation op: %q", m.Op())
	}
}

// SourceClient is a client for the Source schema.
type SourceClient struct {
	config
}

// NewSourceClient returns a client for the Source from the given config.
func NewSourceClient(c config) *SourceClient {
	return &SourceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `source.Hooks(f(g(h())))`.
func (c *SourceClient) Use(hooks ...Hook) {
	c.hooks.Source = append(c.hooks.Source, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `source.Intercept(f(g(h())))`.
func (c *SourceClient) Intercept(interceptors ...Interceptor) {
	c.inters.Source = append(c.inters.Source, interceptors...)
}

// Create returns a builder for creating a Source entity.
func (c *SourceClient) Create() *SourceCreate {
	mutation := newSourceMutation(c.config, OpCreate)
	return &SourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Source entities.
func (c *SourceClient) CreateBulk(builders ...*SourceCreate) *SourceCreateBulk {
	return &SourceCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SourceClient) MapCreateBulk(slice any, setFunc func(*SourceCreate, int)) *SourceCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SourceCreateBulk{err: fmt.Errorf("calling to SourceClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SourceCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SourceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Source.
func (c *SourceClient) Update() *SourceUpdate {
	mutation := newSourceMutation(c.config, OpUpdate)
	return &SourceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SourceClient) UpdateOne(s *Source) *SourceUpdateOne {
	mutation := newSourceMutation(c.config, OpUpdateOne, withSource(s))
	return &SourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SourceClient) UpdateOneID(id int) *SourceUpdateOne {
	mutation := newSourceMutation(c.config, OpUpdateOne, withSourceID(id))
	return &SourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Source.
func (c *SourceClient) Delete() *SourceDelete {
	mutation := newSourceMutation(c.config, OpDelete)
	return &SourceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SourceClient) DeleteOne(s *Source) *SourceDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SourceClient) DeleteOneID(id int) *SourceDeleteOne {
	builder := c.Delete().Where(source.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SourceDeleteOne{builder}
}

// Query returns a query builder for Source.
func (c *SourceClient) Query() *SourceQuery {
	return &SourceQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSource},
		inters: c.Interceptors(),
	}
}

// Get returns a Source entity by its id.
func (c *SourceClient) Get(ctx context.Context, id int) (*Source, error) {
	return c.Query().Where(source.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SourceClient) GetX(ctx context.Context, id int) *Source {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlugin queries the plugin edge of a Source.
func (c *SourceClient) QueryPlugin(s *Source) *PluginQuery {
	query := (&PluginClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(source.Table, source.FieldID, id),
			sqlgraph.To(plugin.Table, plugin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, source.PluginTable, source.PluginColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SourceClient) Hooks() []Hook {
	return c.hooks.Source
}

// Interceptors returns the client interceptors.
func (c *SourceClient) Interceptors() []Interceptor {
	return c.inters.Source
}

func (c *SourceClient) mutate(ctx context.Context, m *SourceMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SourceCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SourceUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SourceDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Source mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlugins queries the plugins edge of a User.
func (c *UserClient) QueryPlugins(u *User) *PluginQuery {
	query := (&PluginClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(plugin.Table, plugin.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PluginsTable, user.PluginsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVerificationSession queries the verification_session edge of a User.
func (c *UserClient) QueryVerificationSession(u *User) *VerificationSessionQuery {
	query := (&VerificationSessionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(verificationsession.Table, verificationsession.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.VerificationSessionTable, user.VerificationSessionColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// VerificationSessionClient is a client for the VerificationSession schema.
type VerificationSessionClient struct {
	config
}

// NewVerificationSessionClient returns a client for the VerificationSession from the given config.
func NewVerificationSessionClient(c config) *VerificationSessionClient {
	return &VerificationSessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `verificationsession.Hooks(f(g(h())))`.
func (c *VerificationSessionClient) Use(hooks ...Hook) {
	c.hooks.VerificationSession = append(c.hooks.VerificationSession, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `verificationsession.Intercept(f(g(h())))`.
func (c *VerificationSessionClient) Intercept(interceptors ...Interceptor) {
	c.inters.VerificationSession = append(c.inters.VerificationSession, interceptors...)
}

// Create returns a builder for creating a VerificationSession entity.
func (c *VerificationSessionClient) Create() *VerificationSessionCreate {
	mutation := newVerificationSessionMutation(c.config, OpCreate)
	return &VerificationSessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of VerificationSession entities.
func (c *VerificationSessionClient) CreateBulk(builders ...*VerificationSessionCreate) *VerificationSessionCreateBulk {
	return &VerificationSessionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *VerificationSessionClient) MapCreateBulk(slice any, setFunc func(*VerificationSessionCreate, int)) *VerificationSessionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &VerificationSessionCreateBulk{err: fmt.Errorf("calling to VerificationSessionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*VerificationSessionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &VerificationSessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for VerificationSession.
func (c *VerificationSessionClient) Update() *VerificationSessionUpdate {
	mutation := newVerificationSessionMutation(c.config, OpUpdate)
	return &VerificationSessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VerificationSessionClient) UpdateOne(vs *VerificationSession) *VerificationSessionUpdateOne {
	mutation := newVerificationSessionMutation(c.config, OpUpdateOne, withVerificationSession(vs))
	return &VerificationSessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VerificationSessionClient) UpdateOneID(id string) *VerificationSessionUpdateOne {
	mutation := newVerificationSessionMutation(c.config, OpUpdateOne, withVerificationSessionID(id))
	return &VerificationSessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for VerificationSession.
func (c *VerificationSessionClient) Delete() *VerificationSessionDelete {
	mutation := newVerificationSessionMutation(c.config, OpDelete)
	return &VerificationSessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VerificationSessionClient) DeleteOne(vs *VerificationSession) *VerificationSessionDeleteOne {
	return c.DeleteOneID(vs.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VerificationSessionClient) DeleteOneID(id string) *VerificationSessionDeleteOne {
	builder := c.Delete().Where(verificationsession.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VerificationSessionDeleteOne{builder}
}

// Query returns a query builder for VerificationSession.
func (c *VerificationSessionClient) Query() *VerificationSessionQuery {
	return &VerificationSessionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeVerificationSession},
		inters: c.Interceptors(),
	}
}

// Get returns a VerificationSession entity by its id.
func (c *VerificationSessionClient) Get(ctx context.Context, id string) (*VerificationSession, error) {
	return c.Query().Where(verificationsession.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VerificationSessionClient) GetX(ctx context.Context, id string) *VerificationSession {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySession queries the session edge of a VerificationSession.
func (c *VerificationSessionClient) QuerySession(vs *VerificationSession) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := vs.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(verificationsession.Table, verificationsession.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, verificationsession.SessionTable, verificationsession.SessionColumn),
		)
		fromV = sqlgraph.Neighbors(vs.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VerificationSessionClient) Hooks() []Hook {
	return c.hooks.VerificationSession
}

// Interceptors returns the client interceptors.
func (c *VerificationSessionClient) Interceptors() []Interceptor {
	return c.inters.VerificationSession
}

func (c *VerificationSessionClient) mutate(ctx context.Context, m *VerificationSessionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&VerificationSessionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&VerificationSessionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&VerificationSessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&VerificationSessionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown VerificationSession mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Plugin, Publication, Source, User, VerificationSession []ent.Hook
	}
	inters struct {
		Plugin, Publication, Source, User, VerificationSession []ent.Interceptor
	}
)
