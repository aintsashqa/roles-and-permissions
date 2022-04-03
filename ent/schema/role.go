package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()).Default(uuid.New).Immutable().Unique(),
		field.String("name").NotEmpty().MaxLen(255),
		field.Time("creation_date").Default(time.Now).Immutable(),
		field.Time("last_update_date").Default(time.Now).UpdateDefault(time.Now),
		//field.Time("mark_as_delete_date").Optional().Nillable(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
