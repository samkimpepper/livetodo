// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeTodoLists holds the string denoting the todo_lists edge name in mutations.
	EdgeTodoLists = "todo_lists"
	// Table holds the table name of the user in the database.
	Table = "users"
	// TodoListsTable is the table that holds the todo_lists relation/edge. The primary key declared below.
	TodoListsTable = "user_todo_lists"
	// TodoListsInverseTable is the table name for the TodoList entity.
	// It exists in this package in order to avoid circular dependency with the "todolist" package.
	TodoListsInverseTable = "todo_lists"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPassword,
	FieldUsername,
	FieldCreatedAt,
}

var (
	// TodoListsPrimaryKey and TodoListsColumn2 are the table columns denoting the
	// primary key for the todo_lists relation (M2M).
	TodoListsPrimaryKey = []string{"user_id", "todo_list_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByTodoListsCount orders the results by todo_lists count.
func ByTodoListsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodoListsStep(), opts...)
	}
}

// ByTodoLists orders the results by todo_lists terms.
func ByTodoLists(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodoListsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodoListsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodoListsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TodoListsTable, TodoListsPrimaryKey...),
	)
}
