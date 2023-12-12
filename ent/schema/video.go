package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Video holds the schema definition for the Video entity.
type Video struct {
	ent.Schema
}

// Fields of the Video.
func (Video) Fields() []ent.Field {
	return []ent.Field{
		field.String("origin_url").
			Unique(),
		field.String("convert_url").
			Optional(),
	}
}

// Edges of the Video.
func (Video) Edges() []ent.Edge {
	return nil
}
