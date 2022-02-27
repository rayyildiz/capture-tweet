package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ContactUs holds the schema definition for the ContactUs entity.
type ContactUs struct {
	ent.Schema
}

// Fields of the ContactUs.
func (ContactUs) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(36).NotEmpty().Unique().Immutable(),
		field.String("email"),
		field.String("full_name").Optional().MaxLen(512),
		field.String("message").Optional().MaxLen(4096),
	}
}

// Edges of the ContactUs.
func (ContactUs) Edges() []ent.Edge {
	return nil
}

func (ContactUs) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
