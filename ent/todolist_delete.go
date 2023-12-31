// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"todo/ent/predicate"
	"todo/ent/todolist"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoListDelete is the builder for deleting a TodoList entity.
type TodoListDelete struct {
	config
	hooks    []Hook
	mutation *TodoListMutation
}

// Where appends a list predicates to the TodoListDelete builder.
func (tld *TodoListDelete) Where(ps ...predicate.TodoList) *TodoListDelete {
	tld.mutation.Where(ps...)
	return tld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tld *TodoListDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tld.sqlExec, tld.mutation, tld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tld *TodoListDelete) ExecX(ctx context.Context) int {
	n, err := tld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tld *TodoListDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(todolist.Table, sqlgraph.NewFieldSpec(todolist.FieldID, field.TypeInt))
	if ps := tld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tld.mutation.done = true
	return affected, err
}

// TodoListDeleteOne is the builder for deleting a single TodoList entity.
type TodoListDeleteOne struct {
	tld *TodoListDelete
}

// Where appends a list predicates to the TodoListDelete builder.
func (tldo *TodoListDeleteOne) Where(ps ...predicate.TodoList) *TodoListDeleteOne {
	tldo.tld.mutation.Where(ps...)
	return tldo
}

// Exec executes the deletion query.
func (tldo *TodoListDeleteOne) Exec(ctx context.Context) error {
	n, err := tldo.tld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{todolist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tldo *TodoListDeleteOne) ExecX(ctx context.Context) {
	if err := tldo.Exec(ctx); err != nil {
		panic(err)
	}
}
