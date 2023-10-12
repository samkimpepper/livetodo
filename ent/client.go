// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"todo/ent/migrate"

	"todo/ent/todoitem"
	"todo/ent/todolist"
	"todo/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// TodoItem is the client for interacting with the TodoItem builders.
	TodoItem *TodoItemClient
	// TodoList is the client for interacting with the TodoList builders.
	TodoList *TodoListClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.TodoItem = NewTodoItemClient(c.config)
	c.TodoList = NewTodoListClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:      ctx,
		config:   cfg,
		TodoItem: NewTodoItemClient(cfg),
		TodoList: NewTodoListClient(cfg),
		User:     NewUserClient(cfg),
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
		ctx:      ctx,
		config:   cfg,
		TodoItem: NewTodoItemClient(cfg),
		TodoList: NewTodoListClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		TodoItem.
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
	c.TodoItem.Use(hooks...)
	c.TodoList.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.TodoItem.Intercept(interceptors...)
	c.TodoList.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *TodoItemMutation:
		return c.TodoItem.mutate(ctx, m)
	case *TodoListMutation:
		return c.TodoList.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// TodoItemClient is a client for the TodoItem schema.
type TodoItemClient struct {
	config
}

// NewTodoItemClient returns a client for the TodoItem from the given config.
func NewTodoItemClient(c config) *TodoItemClient {
	return &TodoItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todoitem.Hooks(f(g(h())))`.
func (c *TodoItemClient) Use(hooks ...Hook) {
	c.hooks.TodoItem = append(c.hooks.TodoItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `todoitem.Intercept(f(g(h())))`.
func (c *TodoItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.TodoItem = append(c.inters.TodoItem, interceptors...)
}

// Create returns a builder for creating a TodoItem entity.
func (c *TodoItemClient) Create() *TodoItemCreate {
	mutation := newTodoItemMutation(c.config, OpCreate)
	return &TodoItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TodoItem entities.
func (c *TodoItemClient) CreateBulk(builders ...*TodoItemCreate) *TodoItemCreateBulk {
	return &TodoItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TodoItemClient) MapCreateBulk(slice any, setFunc func(*TodoItemCreate, int)) *TodoItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TodoItemCreateBulk{err: fmt.Errorf("calling to TodoItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TodoItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TodoItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TodoItem.
func (c *TodoItemClient) Update() *TodoItemUpdate {
	mutation := newTodoItemMutation(c.config, OpUpdate)
	return &TodoItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TodoItemClient) UpdateOne(ti *TodoItem) *TodoItemUpdateOne {
	mutation := newTodoItemMutation(c.config, OpUpdateOne, withTodoItem(ti))
	return &TodoItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TodoItemClient) UpdateOneID(id int) *TodoItemUpdateOne {
	mutation := newTodoItemMutation(c.config, OpUpdateOne, withTodoItemID(id))
	return &TodoItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TodoItem.
func (c *TodoItemClient) Delete() *TodoItemDelete {
	mutation := newTodoItemMutation(c.config, OpDelete)
	return &TodoItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TodoItemClient) DeleteOne(ti *TodoItem) *TodoItemDeleteOne {
	return c.DeleteOneID(ti.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TodoItemClient) DeleteOneID(id int) *TodoItemDeleteOne {
	builder := c.Delete().Where(todoitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TodoItemDeleteOne{builder}
}

// Query returns a query builder for TodoItem.
func (c *TodoItemClient) Query() *TodoItemQuery {
	return &TodoItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTodoItem},
		inters: c.Interceptors(),
	}
}

// Get returns a TodoItem entity by its id.
func (c *TodoItemClient) Get(ctx context.Context, id int) (*TodoItem, error) {
	return c.Query().Where(todoitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TodoItemClient) GetX(ctx context.Context, id int) *TodoItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryList queries the list edge of a TodoItem.
func (c *TodoItemClient) QueryList(ti *TodoItem) *TodoListQuery {
	query := (&TodoListClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ti.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todoitem.Table, todoitem.FieldID, id),
			sqlgraph.To(todolist.Table, todolist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, todoitem.ListTable, todoitem.ListColumn),
		)
		fromV = sqlgraph.Neighbors(ti.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TodoItemClient) Hooks() []Hook {
	return c.hooks.TodoItem
}

// Interceptors returns the client interceptors.
func (c *TodoItemClient) Interceptors() []Interceptor {
	return c.inters.TodoItem
}

func (c *TodoItemClient) mutate(ctx context.Context, m *TodoItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TodoItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TodoItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TodoItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TodoItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TodoItem mutation op: %q", m.Op())
	}
}

// TodoListClient is a client for the TodoList schema.
type TodoListClient struct {
	config
}

// NewTodoListClient returns a client for the TodoList from the given config.
func NewTodoListClient(c config) *TodoListClient {
	return &TodoListClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `todolist.Hooks(f(g(h())))`.
func (c *TodoListClient) Use(hooks ...Hook) {
	c.hooks.TodoList = append(c.hooks.TodoList, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `todolist.Intercept(f(g(h())))`.
func (c *TodoListClient) Intercept(interceptors ...Interceptor) {
	c.inters.TodoList = append(c.inters.TodoList, interceptors...)
}

// Create returns a builder for creating a TodoList entity.
func (c *TodoListClient) Create() *TodoListCreate {
	mutation := newTodoListMutation(c.config, OpCreate)
	return &TodoListCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TodoList entities.
func (c *TodoListClient) CreateBulk(builders ...*TodoListCreate) *TodoListCreateBulk {
	return &TodoListCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TodoListClient) MapCreateBulk(slice any, setFunc func(*TodoListCreate, int)) *TodoListCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TodoListCreateBulk{err: fmt.Errorf("calling to TodoListClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TodoListCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TodoListCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TodoList.
func (c *TodoListClient) Update() *TodoListUpdate {
	mutation := newTodoListMutation(c.config, OpUpdate)
	return &TodoListUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TodoListClient) UpdateOne(tl *TodoList) *TodoListUpdateOne {
	mutation := newTodoListMutation(c.config, OpUpdateOne, withTodoList(tl))
	return &TodoListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TodoListClient) UpdateOneID(id int) *TodoListUpdateOne {
	mutation := newTodoListMutation(c.config, OpUpdateOne, withTodoListID(id))
	return &TodoListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TodoList.
func (c *TodoListClient) Delete() *TodoListDelete {
	mutation := newTodoListMutation(c.config, OpDelete)
	return &TodoListDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TodoListClient) DeleteOne(tl *TodoList) *TodoListDeleteOne {
	return c.DeleteOneID(tl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TodoListClient) DeleteOneID(id int) *TodoListDeleteOne {
	builder := c.Delete().Where(todolist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TodoListDeleteOne{builder}
}

// Query returns a query builder for TodoList.
func (c *TodoListClient) Query() *TodoListQuery {
	return &TodoListQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTodoList},
		inters: c.Interceptors(),
	}
}

// Get returns a TodoList entity by its id.
func (c *TodoListClient) Get(ctx context.Context, id int) (*TodoList, error) {
	return c.Query().Where(todolist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TodoListClient) GetX(ctx context.Context, id int) *TodoList {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryItems queries the items edge of a TodoList.
func (c *TodoListClient) QueryItems(tl *TodoList) *TodoItemQuery {
	query := (&TodoItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todolist.Table, todolist.FieldID, id),
			sqlgraph.To(todoitem.Table, todoitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, todolist.ItemsTable, todolist.ItemsColumn),
		)
		fromV = sqlgraph.Neighbors(tl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a TodoList.
func (c *TodoListClient) QueryUsers(tl *TodoList) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := tl.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(todolist.Table, todolist.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, todolist.UsersTable, todolist.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(tl.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TodoListClient) Hooks() []Hook {
	return c.hooks.TodoList
}

// Interceptors returns the client interceptors.
func (c *TodoListClient) Interceptors() []Interceptor {
	return c.inters.TodoList
}

func (c *TodoListClient) mutate(ctx context.Context, m *TodoListMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TodoListCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TodoListUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TodoListUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TodoListDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TodoList mutation op: %q", m.Op())
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
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
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
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTodoLists queries the todo_lists edge of a User.
func (c *UserClient) QueryTodoLists(u *User) *TodoListQuery {
	query := (&TodoListClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(todolist.Table, todolist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.TodoListsTable, user.TodoListsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
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

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		TodoItem, TodoList, User []ent.Hook
	}
	inters struct {
		TodoItem, TodoList, User []ent.Interceptor
	}
)