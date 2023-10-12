package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TodoItem holds the schema definition for the TodoItem entity.
type TodoItem struct {
	ent.Schema
}

// Fields of the TodoItem.
func (TodoItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Bool("done").Default(false),
		field.Time("due_date").Optional(),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the TodoItem.
func (TodoItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("list", TodoList.Type).Ref("items").Unique(),
	}
}
