// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo/ent/predicate"
	"todo/ent/todoitem"
	"todo/ent/todolist"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoItemUpdate is the builder for updating TodoItem entities.
type TodoItemUpdate struct {
	config
	hooks    []Hook
	mutation *TodoItemMutation
}

// Where appends a list predicates to the TodoItemUpdate builder.
func (tiu *TodoItemUpdate) Where(ps ...predicate.TodoItem) *TodoItemUpdate {
	tiu.mutation.Where(ps...)
	return tiu
}

// SetTitle sets the "title" field.
func (tiu *TodoItemUpdate) SetTitle(s string) *TodoItemUpdate {
	tiu.mutation.SetTitle(s)
	return tiu
}

// SetDone sets the "done" field.
func (tiu *TodoItemUpdate) SetDone(b bool) *TodoItemUpdate {
	tiu.mutation.SetDone(b)
	return tiu
}

// SetNillableDone sets the "done" field if the given value is not nil.
func (tiu *TodoItemUpdate) SetNillableDone(b *bool) *TodoItemUpdate {
	if b != nil {
		tiu.SetDone(*b)
	}
	return tiu
}

// SetDueDate sets the "due_date" field.
func (tiu *TodoItemUpdate) SetDueDate(t time.Time) *TodoItemUpdate {
	tiu.mutation.SetDueDate(t)
	return tiu
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (tiu *TodoItemUpdate) SetNillableDueDate(t *time.Time) *TodoItemUpdate {
	if t != nil {
		tiu.SetDueDate(*t)
	}
	return tiu
}

// ClearDueDate clears the value of the "due_date" field.
func (tiu *TodoItemUpdate) ClearDueDate() *TodoItemUpdate {
	tiu.mutation.ClearDueDate()
	return tiu
}

// SetCreatedAt sets the "created_at" field.
func (tiu *TodoItemUpdate) SetCreatedAt(t time.Time) *TodoItemUpdate {
	tiu.mutation.SetCreatedAt(t)
	return tiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tiu *TodoItemUpdate) SetNillableCreatedAt(t *time.Time) *TodoItemUpdate {
	if t != nil {
		tiu.SetCreatedAt(*t)
	}
	return tiu
}

// SetListID sets the "list" edge to the TodoList entity by ID.
func (tiu *TodoItemUpdate) SetListID(id int) *TodoItemUpdate {
	tiu.mutation.SetListID(id)
	return tiu
}

// SetNillableListID sets the "list" edge to the TodoList entity by ID if the given value is not nil.
func (tiu *TodoItemUpdate) SetNillableListID(id *int) *TodoItemUpdate {
	if id != nil {
		tiu = tiu.SetListID(*id)
	}
	return tiu
}

// SetList sets the "list" edge to the TodoList entity.
func (tiu *TodoItemUpdate) SetList(t *TodoList) *TodoItemUpdate {
	return tiu.SetListID(t.ID)
}

// Mutation returns the TodoItemMutation object of the builder.
func (tiu *TodoItemUpdate) Mutation() *TodoItemMutation {
	return tiu.mutation
}

// ClearList clears the "list" edge to the TodoList entity.
func (tiu *TodoItemUpdate) ClearList() *TodoItemUpdate {
	tiu.mutation.ClearList()
	return tiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tiu *TodoItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tiu.sqlSave, tiu.mutation, tiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tiu *TodoItemUpdate) SaveX(ctx context.Context) int {
	affected, err := tiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tiu *TodoItemUpdate) Exec(ctx context.Context) error {
	_, err := tiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiu *TodoItemUpdate) ExecX(ctx context.Context) {
	if err := tiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tiu *TodoItemUpdate) check() error {
	if v, ok := tiu.mutation.Title(); ok {
		if err := todoitem.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "TodoItem.title": %w`, err)}
		}
	}
	return nil
}

func (tiu *TodoItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todoitem.Table, todoitem.Columns, sqlgraph.NewFieldSpec(todoitem.FieldID, field.TypeInt))
	if ps := tiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiu.mutation.Title(); ok {
		_spec.SetField(todoitem.FieldTitle, field.TypeString, value)
	}
	if value, ok := tiu.mutation.Done(); ok {
		_spec.SetField(todoitem.FieldDone, field.TypeBool, value)
	}
	if value, ok := tiu.mutation.DueDate(); ok {
		_spec.SetField(todoitem.FieldDueDate, field.TypeTime, value)
	}
	if tiu.mutation.DueDateCleared() {
		_spec.ClearField(todoitem.FieldDueDate, field.TypeTime)
	}
	if value, ok := tiu.mutation.CreatedAt(); ok {
		_spec.SetField(todoitem.FieldCreatedAt, field.TypeTime, value)
	}
	if tiu.mutation.ListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todoitem.ListTable,
			Columns: []string{todoitem.ListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todolist.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tiu.mutation.ListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todoitem.ListTable,
			Columns: []string{todoitem.ListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todolist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todoitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tiu.mutation.done = true
	return n, nil
}

// TodoItemUpdateOne is the builder for updating a single TodoItem entity.
type TodoItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoItemMutation
}

// SetTitle sets the "title" field.
func (tiuo *TodoItemUpdateOne) SetTitle(s string) *TodoItemUpdateOne {
	tiuo.mutation.SetTitle(s)
	return tiuo
}

// SetDone sets the "done" field.
func (tiuo *TodoItemUpdateOne) SetDone(b bool) *TodoItemUpdateOne {
	tiuo.mutation.SetDone(b)
	return tiuo
}

// SetNillableDone sets the "done" field if the given value is not nil.
func (tiuo *TodoItemUpdateOne) SetNillableDone(b *bool) *TodoItemUpdateOne {
	if b != nil {
		tiuo.SetDone(*b)
	}
	return tiuo
}

// SetDueDate sets the "due_date" field.
func (tiuo *TodoItemUpdateOne) SetDueDate(t time.Time) *TodoItemUpdateOne {
	tiuo.mutation.SetDueDate(t)
	return tiuo
}

// SetNillableDueDate sets the "due_date" field if the given value is not nil.
func (tiuo *TodoItemUpdateOne) SetNillableDueDate(t *time.Time) *TodoItemUpdateOne {
	if t != nil {
		tiuo.SetDueDate(*t)
	}
	return tiuo
}

// ClearDueDate clears the value of the "due_date" field.
func (tiuo *TodoItemUpdateOne) ClearDueDate() *TodoItemUpdateOne {
	tiuo.mutation.ClearDueDate()
	return tiuo
}

// SetCreatedAt sets the "created_at" field.
func (tiuo *TodoItemUpdateOne) SetCreatedAt(t time.Time) *TodoItemUpdateOne {
	tiuo.mutation.SetCreatedAt(t)
	return tiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tiuo *TodoItemUpdateOne) SetNillableCreatedAt(t *time.Time) *TodoItemUpdateOne {
	if t != nil {
		tiuo.SetCreatedAt(*t)
	}
	return tiuo
}

// SetListID sets the "list" edge to the TodoList entity by ID.
func (tiuo *TodoItemUpdateOne) SetListID(id int) *TodoItemUpdateOne {
	tiuo.mutation.SetListID(id)
	return tiuo
}

// SetNillableListID sets the "list" edge to the TodoList entity by ID if the given value is not nil.
func (tiuo *TodoItemUpdateOne) SetNillableListID(id *int) *TodoItemUpdateOne {
	if id != nil {
		tiuo = tiuo.SetListID(*id)
	}
	return tiuo
}

// SetList sets the "list" edge to the TodoList entity.
func (tiuo *TodoItemUpdateOne) SetList(t *TodoList) *TodoItemUpdateOne {
	return tiuo.SetListID(t.ID)
}

// Mutation returns the TodoItemMutation object of the builder.
func (tiuo *TodoItemUpdateOne) Mutation() *TodoItemMutation {
	return tiuo.mutation
}

// ClearList clears the "list" edge to the TodoList entity.
func (tiuo *TodoItemUpdateOne) ClearList() *TodoItemUpdateOne {
	tiuo.mutation.ClearList()
	return tiuo
}

// Where appends a list predicates to the TodoItemUpdate builder.
func (tiuo *TodoItemUpdateOne) Where(ps ...predicate.TodoItem) *TodoItemUpdateOne {
	tiuo.mutation.Where(ps...)
	return tiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tiuo *TodoItemUpdateOne) Select(field string, fields ...string) *TodoItemUpdateOne {
	tiuo.fields = append([]string{field}, fields...)
	return tiuo
}

// Save executes the query and returns the updated TodoItem entity.
func (tiuo *TodoItemUpdateOne) Save(ctx context.Context) (*TodoItem, error) {
	return withHooks(ctx, tiuo.sqlSave, tiuo.mutation, tiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tiuo *TodoItemUpdateOne) SaveX(ctx context.Context) *TodoItem {
	node, err := tiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tiuo *TodoItemUpdateOne) Exec(ctx context.Context) error {
	_, err := tiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tiuo *TodoItemUpdateOne) ExecX(ctx context.Context) {
	if err := tiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tiuo *TodoItemUpdateOne) check() error {
	if v, ok := tiuo.mutation.Title(); ok {
		if err := todoitem.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "TodoItem.title": %w`, err)}
		}
	}
	return nil
}

func (tiuo *TodoItemUpdateOne) sqlSave(ctx context.Context) (_node *TodoItem, err error) {
	if err := tiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todoitem.Table, todoitem.Columns, sqlgraph.NewFieldSpec(todoitem.FieldID, field.TypeInt))
	id, ok := tiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TodoItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todoitem.FieldID)
		for _, f := range fields {
			if !todoitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != todoitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tiuo.mutation.Title(); ok {
		_spec.SetField(todoitem.FieldTitle, field.TypeString, value)
	}
	if value, ok := tiuo.mutation.Done(); ok {
		_spec.SetField(todoitem.FieldDone, field.TypeBool, value)
	}
	if value, ok := tiuo.mutation.DueDate(); ok {
		_spec.SetField(todoitem.FieldDueDate, field.TypeTime, value)
	}
	if tiuo.mutation.DueDateCleared() {
		_spec.ClearField(todoitem.FieldDueDate, field.TypeTime)
	}
	if value, ok := tiuo.mutation.CreatedAt(); ok {
		_spec.SetField(todoitem.FieldCreatedAt, field.TypeTime, value)
	}
	if tiuo.mutation.ListCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todoitem.ListTable,
			Columns: []string{todoitem.ListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todolist.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tiuo.mutation.ListIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   todoitem.ListTable,
			Columns: []string{todoitem.ListColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todolist.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TodoItem{config: tiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todoitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tiuo.mutation.done = true
	return _node, nil
}
