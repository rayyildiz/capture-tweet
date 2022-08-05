// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"capturetweet.com/internal/ent/predicate"
	"capturetweet.com/internal/ent/schema"
	"capturetweet.com/internal/ent/tweet"
	"capturetweet.com/internal/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (tu *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TweetUpdate) SetUpdatedAt(t time.Time) *TweetUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetFullText sets the "full_text" field.
func (tu *TweetUpdate) SetFullText(s string) *TweetUpdate {
	tu.mutation.SetFullText(s)
	return tu
}

// SetCaptureURL sets the "capture_url" field.
func (tu *TweetUpdate) SetCaptureURL(s string) *TweetUpdate {
	tu.mutation.SetCaptureURL(s)
	return tu
}

// SetNillableCaptureURL sets the "capture_url" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableCaptureURL(s *string) *TweetUpdate {
	if s != nil {
		tu.SetCaptureURL(*s)
	}
	return tu
}

// ClearCaptureURL clears the value of the "capture_url" field.
func (tu *TweetUpdate) ClearCaptureURL() *TweetUpdate {
	tu.mutation.ClearCaptureURL()
	return tu
}

// SetCaptureThumbURL sets the "capture_thumb_url" field.
func (tu *TweetUpdate) SetCaptureThumbURL(s string) *TweetUpdate {
	tu.mutation.SetCaptureThumbURL(s)
	return tu
}

// SetNillableCaptureThumbURL sets the "capture_thumb_url" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableCaptureThumbURL(s *string) *TweetUpdate {
	if s != nil {
		tu.SetCaptureThumbURL(*s)
	}
	return tu
}

// ClearCaptureThumbURL clears the value of the "capture_thumb_url" field.
func (tu *TweetUpdate) ClearCaptureThumbURL() *TweetUpdate {
	tu.mutation.ClearCaptureThumbURL()
	return tu
}

// SetFavoriteCount sets the "favorite_count" field.
func (tu *TweetUpdate) SetFavoriteCount(i int) *TweetUpdate {
	tu.mutation.ResetFavoriteCount()
	tu.mutation.SetFavoriteCount(i)
	return tu
}

// SetNillableFavoriteCount sets the "favorite_count" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableFavoriteCount(i *int) *TweetUpdate {
	if i != nil {
		tu.SetFavoriteCount(*i)
	}
	return tu
}

// AddFavoriteCount adds i to the "favorite_count" field.
func (tu *TweetUpdate) AddFavoriteCount(i int) *TweetUpdate {
	tu.mutation.AddFavoriteCount(i)
	return tu
}

// SetRetweetCount sets the "retweet_count" field.
func (tu *TweetUpdate) SetRetweetCount(i int) *TweetUpdate {
	tu.mutation.ResetRetweetCount()
	tu.mutation.SetRetweetCount(i)
	return tu
}

// SetNillableRetweetCount sets the "retweet_count" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableRetweetCount(i *int) *TweetUpdate {
	if i != nil {
		tu.SetRetweetCount(*i)
	}
	return tu
}

// AddRetweetCount adds i to the "retweet_count" field.
func (tu *TweetUpdate) AddRetweetCount(i int) *TweetUpdate {
	tu.mutation.AddRetweetCount(i)
	return tu
}

// SetAuthorID sets the "author_id" field.
func (tu *TweetUpdate) SetAuthorID(s string) *TweetUpdate {
	tu.mutation.SetAuthorID(s)
	return tu
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (tu *TweetUpdate) SetNillableAuthorID(s *string) *TweetUpdate {
	if s != nil {
		tu.SetAuthorID(*s)
	}
	return tu
}

// ClearAuthorID clears the value of the "author_id" field.
func (tu *TweetUpdate) ClearAuthorID() *TweetUpdate {
	tu.mutation.ClearAuthorID()
	return tu
}

// SetResources sets the "resources" field.
func (tu *TweetUpdate) SetResources(s []schema.Resource) *TweetUpdate {
	tu.mutation.SetResources(s)
	return tu
}

// SetPostedAt sets the "posted_at" field.
func (tu *TweetUpdate) SetPostedAt(t time.Time) *TweetUpdate {
	tu.mutation.SetPostedAt(t)
	return tu
}

// SetAuthor sets the "author" edge to the User entity.
func (tu *TweetUpdate) SetAuthor(u *User) *TweetUpdate {
	return tu.SetAuthorID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tu *TweetUpdate) Mutation() *TweetMutation {
	return tu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (tu *TweetUpdate) ClearAuthor() *TweetUpdate {
	tu.mutation.ClearAuthor()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TweetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TweetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TweetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TweetUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := tweet.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TweetUpdate) check() error {
	if v, ok := tu.mutation.FullText(); ok {
		if err := tweet.FullTextValidator(v); err != nil {
			return &ValidationError{Name: "full_text", err: fmt.Errorf(`ent: validator failed for field "Tweet.full_text": %w`, err)}
		}
	}
	if v, ok := tu.mutation.AuthorID(); ok {
		if err := tweet.AuthorIDValidator(v); err != nil {
			return &ValidationError{Name: "author_id", err: fmt.Errorf(`ent: validator failed for field "Tweet.author_id": %w`, err)}
		}
	}
	return nil
}

func (tu *TweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tweet.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tweet.FieldUpdatedAt,
		})
	}
	if value, ok := tu.mutation.FullText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldFullText,
		})
	}
	if value, ok := tu.mutation.CaptureURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldCaptureURL,
		})
	}
	if tu.mutation.CaptureURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tweet.FieldCaptureURL,
		})
	}
	if value, ok := tu.mutation.CaptureThumbURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldCaptureThumbURL,
		})
	}
	if tu.mutation.CaptureThumbURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tweet.FieldCaptureThumbURL,
		})
	}
	if value, ok := tu.mutation.FavoriteCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldFavoriteCount,
		})
	}
	if value, ok := tu.mutation.AddedFavoriteCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldFavoriteCount,
		})
	}
	if value, ok := tu.mutation.RetweetCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldRetweetCount,
		})
	}
	if value, ok := tu.mutation.AddedRetweetCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldRetweetCount,
		})
	}
	if value, ok := tu.mutation.Resources(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: tweet.FieldResources,
		})
	}
	if value, ok := tu.mutation.PostedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tweet.FieldPostedAt,
		})
	}
	if tu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TweetUpdateOne) SetUpdatedAt(t time.Time) *TweetUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetFullText sets the "full_text" field.
func (tuo *TweetUpdateOne) SetFullText(s string) *TweetUpdateOne {
	tuo.mutation.SetFullText(s)
	return tuo
}

// SetCaptureURL sets the "capture_url" field.
func (tuo *TweetUpdateOne) SetCaptureURL(s string) *TweetUpdateOne {
	tuo.mutation.SetCaptureURL(s)
	return tuo
}

// SetNillableCaptureURL sets the "capture_url" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableCaptureURL(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetCaptureURL(*s)
	}
	return tuo
}

// ClearCaptureURL clears the value of the "capture_url" field.
func (tuo *TweetUpdateOne) ClearCaptureURL() *TweetUpdateOne {
	tuo.mutation.ClearCaptureURL()
	return tuo
}

// SetCaptureThumbURL sets the "capture_thumb_url" field.
func (tuo *TweetUpdateOne) SetCaptureThumbURL(s string) *TweetUpdateOne {
	tuo.mutation.SetCaptureThumbURL(s)
	return tuo
}

// SetNillableCaptureThumbURL sets the "capture_thumb_url" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableCaptureThumbURL(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetCaptureThumbURL(*s)
	}
	return tuo
}

// ClearCaptureThumbURL clears the value of the "capture_thumb_url" field.
func (tuo *TweetUpdateOne) ClearCaptureThumbURL() *TweetUpdateOne {
	tuo.mutation.ClearCaptureThumbURL()
	return tuo
}

// SetFavoriteCount sets the "favorite_count" field.
func (tuo *TweetUpdateOne) SetFavoriteCount(i int) *TweetUpdateOne {
	tuo.mutation.ResetFavoriteCount()
	tuo.mutation.SetFavoriteCount(i)
	return tuo
}

// SetNillableFavoriteCount sets the "favorite_count" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableFavoriteCount(i *int) *TweetUpdateOne {
	if i != nil {
		tuo.SetFavoriteCount(*i)
	}
	return tuo
}

// AddFavoriteCount adds i to the "favorite_count" field.
func (tuo *TweetUpdateOne) AddFavoriteCount(i int) *TweetUpdateOne {
	tuo.mutation.AddFavoriteCount(i)
	return tuo
}

// SetRetweetCount sets the "retweet_count" field.
func (tuo *TweetUpdateOne) SetRetweetCount(i int) *TweetUpdateOne {
	tuo.mutation.ResetRetweetCount()
	tuo.mutation.SetRetweetCount(i)
	return tuo
}

// SetNillableRetweetCount sets the "retweet_count" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableRetweetCount(i *int) *TweetUpdateOne {
	if i != nil {
		tuo.SetRetweetCount(*i)
	}
	return tuo
}

// AddRetweetCount adds i to the "retweet_count" field.
func (tuo *TweetUpdateOne) AddRetweetCount(i int) *TweetUpdateOne {
	tuo.mutation.AddRetweetCount(i)
	return tuo
}

// SetAuthorID sets the "author_id" field.
func (tuo *TweetUpdateOne) SetAuthorID(s string) *TweetUpdateOne {
	tuo.mutation.SetAuthorID(s)
	return tuo
}

// SetNillableAuthorID sets the "author_id" field if the given value is not nil.
func (tuo *TweetUpdateOne) SetNillableAuthorID(s *string) *TweetUpdateOne {
	if s != nil {
		tuo.SetAuthorID(*s)
	}
	return tuo
}

// ClearAuthorID clears the value of the "author_id" field.
func (tuo *TweetUpdateOne) ClearAuthorID() *TweetUpdateOne {
	tuo.mutation.ClearAuthorID()
	return tuo
}

// SetResources sets the "resources" field.
func (tuo *TweetUpdateOne) SetResources(s []schema.Resource) *TweetUpdateOne {
	tuo.mutation.SetResources(s)
	return tuo
}

// SetPostedAt sets the "posted_at" field.
func (tuo *TweetUpdateOne) SetPostedAt(t time.Time) *TweetUpdateOne {
	tuo.mutation.SetPostedAt(t)
	return tuo
}

// SetAuthor sets the "author" edge to the User entity.
func (tuo *TweetUpdateOne) SetAuthor(u *User) *TweetUpdateOne {
	return tuo.SetAuthorID(u.ID)
}

// Mutation returns the TweetMutation object of the builder.
func (tuo *TweetUpdateOne) Mutation() *TweetMutation {
	return tuo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (tuo *TweetUpdateOne) ClearAuthor() *TweetUpdateOne {
	tuo.mutation.ClearAuthor()
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tweet entity.
func (tuo *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	var (
		err  error
		node *Tweet
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Tweet)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TweetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TweetUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := tweet.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TweetUpdateOne) check() error {
	if v, ok := tuo.mutation.FullText(); ok {
		if err := tweet.FullTextValidator(v); err != nil {
			return &ValidationError{Name: "full_text", err: fmt.Errorf(`ent: validator failed for field "Tweet.full_text": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.AuthorID(); ok {
		if err := tweet.AuthorIDValidator(v); err != nil {
			return &ValidationError{Name: "author_id", err: fmt.Errorf(`ent: validator failed for field "Tweet.author_id": %w`, err)}
		}
	}
	return nil
}

func (tuo *TweetUpdateOne) sqlSave(ctx context.Context) (_node *Tweet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tweet.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tweet.FieldUpdatedAt,
		})
	}
	if value, ok := tuo.mutation.FullText(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldFullText,
		})
	}
	if value, ok := tuo.mutation.CaptureURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldCaptureURL,
		})
	}
	if tuo.mutation.CaptureURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tweet.FieldCaptureURL,
		})
	}
	if value, ok := tuo.mutation.CaptureThumbURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tweet.FieldCaptureThumbURL,
		})
	}
	if tuo.mutation.CaptureThumbURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: tweet.FieldCaptureThumbURL,
		})
	}
	if value, ok := tuo.mutation.FavoriteCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldFavoriteCount,
		})
	}
	if value, ok := tuo.mutation.AddedFavoriteCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldFavoriteCount,
		})
	}
	if value, ok := tuo.mutation.RetweetCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldRetweetCount,
		})
	}
	if value, ok := tuo.mutation.AddedRetweetCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tweet.FieldRetweetCount,
		})
	}
	if value, ok := tuo.mutation.Resources(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: tweet.FieldResources,
		})
	}
	if value, ok := tuo.mutation.PostedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tweet.FieldPostedAt,
		})
	}
	if tuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tweet{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
