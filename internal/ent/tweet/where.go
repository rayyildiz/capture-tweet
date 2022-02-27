// Code generated by entc, DO NOT EDIT.

package tweet

import (
	"time"

	"capturetweet.com/internal/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// FullText applies equality check predicate on the "full_text" field. It's identical to FullTextEQ.
func FullText(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullText), v))
	})
}

// CaptureURL applies equality check predicate on the "capture_url" field. It's identical to CaptureURLEQ.
func CaptureURL(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCaptureURL), v))
	})
}

// CaptureThumbURL applies equality check predicate on the "capture_thumb_url" field. It's identical to CaptureThumbURLEQ.
func CaptureThumbURL(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCaptureThumbURL), v))
	})
}

// Lang applies equality check predicate on the "lang" field. It's identical to LangEQ.
func Lang(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLang), v))
	})
}

// FavoriteCount applies equality check predicate on the "favorite_count" field. It's identical to FavoriteCountEQ.
func FavoriteCount(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFavoriteCount), v))
	})
}

// RetweetCount applies equality check predicate on the "retweet_count" field. It's identical to RetweetCountEQ.
func RetweetCount(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetweetCount), v))
	})
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorID), v))
	})
}

// PostedAt applies equality check predicate on the "posted_at" field. It's identical to PostedAtEQ.
func PostedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostedAt), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// FullTextEQ applies the EQ predicate on the "full_text" field.
func FullTextEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFullText), v))
	})
}

// FullTextNEQ applies the NEQ predicate on the "full_text" field.
func FullTextNEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFullText), v))
	})
}

// FullTextIn applies the In predicate on the "full_text" field.
func FullTextIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFullText), v...))
	})
}

// FullTextNotIn applies the NotIn predicate on the "full_text" field.
func FullTextNotIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFullText), v...))
	})
}

// FullTextGT applies the GT predicate on the "full_text" field.
func FullTextGT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFullText), v))
	})
}

// FullTextGTE applies the GTE predicate on the "full_text" field.
func FullTextGTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFullText), v))
	})
}

// FullTextLT applies the LT predicate on the "full_text" field.
func FullTextLT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFullText), v))
	})
}

// FullTextLTE applies the LTE predicate on the "full_text" field.
func FullTextLTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFullText), v))
	})
}

// FullTextContains applies the Contains predicate on the "full_text" field.
func FullTextContains(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFullText), v))
	})
}

// FullTextHasPrefix applies the HasPrefix predicate on the "full_text" field.
func FullTextHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFullText), v))
	})
}

// FullTextHasSuffix applies the HasSuffix predicate on the "full_text" field.
func FullTextHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFullText), v))
	})
}

// FullTextEqualFold applies the EqualFold predicate on the "full_text" field.
func FullTextEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFullText), v))
	})
}

// FullTextContainsFold applies the ContainsFold predicate on the "full_text" field.
func FullTextContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFullText), v))
	})
}

// CaptureURLEQ applies the EQ predicate on the "capture_url" field.
func CaptureURLEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLNEQ applies the NEQ predicate on the "capture_url" field.
func CaptureURLNEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLIn applies the In predicate on the "capture_url" field.
func CaptureURLIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCaptureURL), v...))
	})
}

// CaptureURLNotIn applies the NotIn predicate on the "capture_url" field.
func CaptureURLNotIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCaptureURL), v...))
	})
}

// CaptureURLGT applies the GT predicate on the "capture_url" field.
func CaptureURLGT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLGTE applies the GTE predicate on the "capture_url" field.
func CaptureURLGTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLLT applies the LT predicate on the "capture_url" field.
func CaptureURLLT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLLTE applies the LTE predicate on the "capture_url" field.
func CaptureURLLTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLContains applies the Contains predicate on the "capture_url" field.
func CaptureURLContains(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLHasPrefix applies the HasPrefix predicate on the "capture_url" field.
func CaptureURLHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLHasSuffix applies the HasSuffix predicate on the "capture_url" field.
func CaptureURLHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLIsNil applies the IsNil predicate on the "capture_url" field.
func CaptureURLIsNil() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCaptureURL)))
	})
}

// CaptureURLNotNil applies the NotNil predicate on the "capture_url" field.
func CaptureURLNotNil() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCaptureURL)))
	})
}

// CaptureURLEqualFold applies the EqualFold predicate on the "capture_url" field.
func CaptureURLEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCaptureURL), v))
	})
}

// CaptureURLContainsFold applies the ContainsFold predicate on the "capture_url" field.
func CaptureURLContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCaptureURL), v))
	})
}

// CaptureThumbURLEQ applies the EQ predicate on the "capture_thumb_url" field.
func CaptureThumbURLEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLNEQ applies the NEQ predicate on the "capture_thumb_url" field.
func CaptureThumbURLNEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLIn applies the In predicate on the "capture_thumb_url" field.
func CaptureThumbURLIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCaptureThumbURL), v...))
	})
}

// CaptureThumbURLNotIn applies the NotIn predicate on the "capture_thumb_url" field.
func CaptureThumbURLNotIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCaptureThumbURL), v...))
	})
}

// CaptureThumbURLGT applies the GT predicate on the "capture_thumb_url" field.
func CaptureThumbURLGT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLGTE applies the GTE predicate on the "capture_thumb_url" field.
func CaptureThumbURLGTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLLT applies the LT predicate on the "capture_thumb_url" field.
func CaptureThumbURLLT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLLTE applies the LTE predicate on the "capture_thumb_url" field.
func CaptureThumbURLLTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLContains applies the Contains predicate on the "capture_thumb_url" field.
func CaptureThumbURLContains(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLHasPrefix applies the HasPrefix predicate on the "capture_thumb_url" field.
func CaptureThumbURLHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLHasSuffix applies the HasSuffix predicate on the "capture_thumb_url" field.
func CaptureThumbURLHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLIsNil applies the IsNil predicate on the "capture_thumb_url" field.
func CaptureThumbURLIsNil() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCaptureThumbURL)))
	})
}

// CaptureThumbURLNotNil applies the NotNil predicate on the "capture_thumb_url" field.
func CaptureThumbURLNotNil() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCaptureThumbURL)))
	})
}

// CaptureThumbURLEqualFold applies the EqualFold predicate on the "capture_thumb_url" field.
func CaptureThumbURLEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCaptureThumbURL), v))
	})
}

// CaptureThumbURLContainsFold applies the ContainsFold predicate on the "capture_thumb_url" field.
func CaptureThumbURLContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCaptureThumbURL), v))
	})
}

// LangEQ applies the EQ predicate on the "lang" field.
func LangEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLang), v))
	})
}

// LangNEQ applies the NEQ predicate on the "lang" field.
func LangNEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLang), v))
	})
}

// LangIn applies the In predicate on the "lang" field.
func LangIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLang), v...))
	})
}

// LangNotIn applies the NotIn predicate on the "lang" field.
func LangNotIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLang), v...))
	})
}

// LangGT applies the GT predicate on the "lang" field.
func LangGT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLang), v))
	})
}

// LangGTE applies the GTE predicate on the "lang" field.
func LangGTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLang), v))
	})
}

// LangLT applies the LT predicate on the "lang" field.
func LangLT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLang), v))
	})
}

// LangLTE applies the LTE predicate on the "lang" field.
func LangLTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLang), v))
	})
}

// LangContains applies the Contains predicate on the "lang" field.
func LangContains(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLang), v))
	})
}

// LangHasPrefix applies the HasPrefix predicate on the "lang" field.
func LangHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLang), v))
	})
}

// LangHasSuffix applies the HasSuffix predicate on the "lang" field.
func LangHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLang), v))
	})
}

// LangEqualFold applies the EqualFold predicate on the "lang" field.
func LangEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLang), v))
	})
}

// LangContainsFold applies the ContainsFold predicate on the "lang" field.
func LangContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLang), v))
	})
}

// FavoriteCountEQ applies the EQ predicate on the "favorite_count" field.
func FavoriteCountEQ(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFavoriteCount), v))
	})
}

// FavoriteCountNEQ applies the NEQ predicate on the "favorite_count" field.
func FavoriteCountNEQ(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFavoriteCount), v))
	})
}

// FavoriteCountIn applies the In predicate on the "favorite_count" field.
func FavoriteCountIn(vs ...int) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFavoriteCount), v...))
	})
}

// FavoriteCountNotIn applies the NotIn predicate on the "favorite_count" field.
func FavoriteCountNotIn(vs ...int) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFavoriteCount), v...))
	})
}

// FavoriteCountGT applies the GT predicate on the "favorite_count" field.
func FavoriteCountGT(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFavoriteCount), v))
	})
}

// FavoriteCountGTE applies the GTE predicate on the "favorite_count" field.
func FavoriteCountGTE(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFavoriteCount), v))
	})
}

// FavoriteCountLT applies the LT predicate on the "favorite_count" field.
func FavoriteCountLT(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFavoriteCount), v))
	})
}

// FavoriteCountLTE applies the LTE predicate on the "favorite_count" field.
func FavoriteCountLTE(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFavoriteCount), v))
	})
}

// RetweetCountEQ applies the EQ predicate on the "retweet_count" field.
func RetweetCountEQ(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRetweetCount), v))
	})
}

// RetweetCountNEQ applies the NEQ predicate on the "retweet_count" field.
func RetweetCountNEQ(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRetweetCount), v))
	})
}

// RetweetCountIn applies the In predicate on the "retweet_count" field.
func RetweetCountIn(vs ...int) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRetweetCount), v...))
	})
}

// RetweetCountNotIn applies the NotIn predicate on the "retweet_count" field.
func RetweetCountNotIn(vs ...int) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRetweetCount), v...))
	})
}

// RetweetCountGT applies the GT predicate on the "retweet_count" field.
func RetweetCountGT(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRetweetCount), v))
	})
}

// RetweetCountGTE applies the GTE predicate on the "retweet_count" field.
func RetweetCountGTE(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRetweetCount), v))
	})
}

// RetweetCountLT applies the LT predicate on the "retweet_count" field.
func RetweetCountLT(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRetweetCount), v))
	})
}

// RetweetCountLTE applies the LTE predicate on the "retweet_count" field.
func RetweetCountLTE(v int) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRetweetCount), v))
	})
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthorID), v))
	})
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthorID), v))
	})
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthorID), v...))
	})
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...string) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthorID), v...))
	})
}

// AuthorIDGT applies the GT predicate on the "author_id" field.
func AuthorIDGT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthorID), v))
	})
}

// AuthorIDGTE applies the GTE predicate on the "author_id" field.
func AuthorIDGTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthorID), v))
	})
}

// AuthorIDLT applies the LT predicate on the "author_id" field.
func AuthorIDLT(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthorID), v))
	})
}

// AuthorIDLTE applies the LTE predicate on the "author_id" field.
func AuthorIDLTE(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthorID), v))
	})
}

// AuthorIDContains applies the Contains predicate on the "author_id" field.
func AuthorIDContains(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthorID), v))
	})
}

// AuthorIDHasPrefix applies the HasPrefix predicate on the "author_id" field.
func AuthorIDHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthorID), v))
	})
}

// AuthorIDHasSuffix applies the HasSuffix predicate on the "author_id" field.
func AuthorIDHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthorID), v))
	})
}

// AuthorIDIsNil applies the IsNil predicate on the "author_id" field.
func AuthorIDIsNil() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAuthorID)))
	})
}

// AuthorIDNotNil applies the NotNil predicate on the "author_id" field.
func AuthorIDNotNil() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAuthorID)))
	})
}

// AuthorIDEqualFold applies the EqualFold predicate on the "author_id" field.
func AuthorIDEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthorID), v))
	})
}

// AuthorIDContainsFold applies the ContainsFold predicate on the "author_id" field.
func AuthorIDContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthorID), v))
	})
}

// PostedAtEQ applies the EQ predicate on the "posted_at" field.
func PostedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPostedAt), v))
	})
}

// PostedAtNEQ applies the NEQ predicate on the "posted_at" field.
func PostedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPostedAt), v))
	})
}

// PostedAtIn applies the In predicate on the "posted_at" field.
func PostedAtIn(vs ...time.Time) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPostedAt), v...))
	})
}

// PostedAtNotIn applies the NotIn predicate on the "posted_at" field.
func PostedAtNotIn(vs ...time.Time) predicate.Tweet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tweet(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPostedAt), v...))
	})
}

// PostedAtGT applies the GT predicate on the "posted_at" field.
func PostedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPostedAt), v))
	})
}

// PostedAtGTE applies the GTE predicate on the "posted_at" field.
func PostedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPostedAt), v))
	})
}

// PostedAtLT applies the LT predicate on the "posted_at" field.
func PostedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPostedAt), v))
	})
}

// PostedAtLTE applies the LTE predicate on the "posted_at" field.
func PostedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPostedAt), v))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		p(s.Not())
	})
}
