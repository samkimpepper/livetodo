package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TodoList holds the schema definition for the TodoList entity.
type TodoList struct {
	ent.Schema
}

// Fields of the TodoList.
func (TodoList) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the TodoList.
func (TodoList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("items", TodoItem.Type),
		edge.From("users", User.Type).Ref("todo_lists"),
	}
}
