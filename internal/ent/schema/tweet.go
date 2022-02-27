package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

type Resource struct {
	URL       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	MediaType string `json:"media_type"`
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(32).NotEmpty().Unique().Immutable(),
		field.String("full_text").NotEmpty().MaxLen(1024),
		field.String("capture_url").Optional(),
		field.String("capture_thumb_url").Optional(),
		field.String("lang").MaxLen(4).NotEmpty().Immutable(),
		field.Int("favorite_count").Default(0),
		field.Int("retweet_count").Default(0),
		field.String("author_id").Optional().MaxLen(32),
		field.JSON("resources", []Resource{}),
		field.Time("posted_at"),
	}
}

func (Tweet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).Field("author_id").Unique(),
	}
}

type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
