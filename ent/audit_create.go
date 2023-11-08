// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tiagoposse/connect/ent/audit"
	"github.com/tiagoposse/connect/ent/user"
)

// AuditCreate is the builder for creating a Audit entity.
type AuditCreate struct {
	config
	mutation *AuditMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAction sets the "action" field.
func (ac *AuditCreate) SetAction(s string) *AuditCreate {
	ac.mutation.SetAction(s)
	return ac
}

// SetAuthor sets the "author" field.
func (ac *AuditCreate) SetAuthor(s string) *AuditCreate {
	ac.mutation.SetAuthor(s)
	return ac
}

// SetID sets the "id" field.
func (ac *AuditCreate) SetID(s string) *AuditCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AuditCreate) SetNillableID(s *string) *AuditCreate {
	if s != nil {
		ac.SetID(*s)
	}
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AuditCreate) SetUserID(id string) *AuditCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AuditCreate) SetUser(u *User) *AuditCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AuditMutation object of the builder.
func (ac *AuditCreate) Mutation() *AuditMutation {
	return ac.mutation
}

// Save creates the Audit in the database.
func (ac *AuditCreate) Save(ctx context.Context) (*Audit, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AuditCreate) SaveX(ctx context.Context) *Audit {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AuditCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AuditCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AuditCreate) defaults() {
	if _, ok := ac.mutation.ID(); !ok {
		v := audit.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AuditCreate) check() error {
	if _, ok := ac.mutation.Action(); !ok {
		return &ValidationError{Name: "action", err: errors.New(`ent: missing required field "Audit.action"`)}
	}
	if v, ok := ac.mutation.Action(); ok {
		if err := audit.ActionValidator(v); err != nil {
			return &ValidationError{Name: "action", err: fmt.Errorf(`ent: validator failed for field "Audit.action": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Author(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required field "Audit.author"`)}
	}
	if v, ok := ac.mutation.Author(); ok {
		if err := audit.AuthorValidator(v); err != nil {
			return &ValidationError{Name: "author", err: fmt.Errorf(`ent: validator failed for field "Audit.author": %w`, err)}
		}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Audit.user"`)}
	}
	return nil
}

func (ac *AuditCreate) sqlSave(ctx context.Context) (*Audit, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Audit.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AuditCreate) createSpec() (*Audit, *sqlgraph.CreateSpec) {
	var (
		_node = &Audit{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(audit.Table, sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.Action(); ok {
		_spec.SetField(audit.FieldAction, field.TypeString, value)
		_node.Action = value
	}
	if value, ok := ac.mutation.Author(); ok {
		_spec.SetField(audit.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   audit.UserTable,
			Columns: []string{audit.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_audit = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Audit.Create().
//		SetAction(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuditUpsert) {
//			SetAction(v+v).
//		}).
//		Exec(ctx)
func (ac *AuditCreate) OnConflict(opts ...sql.ConflictOption) *AuditUpsertOne {
	ac.conflict = opts
	return &AuditUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Audit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AuditCreate) OnConflictColumns(columns ...string) *AuditUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AuditUpsertOne{
		create: ac,
	}
}

type (
	// AuditUpsertOne is the builder for "upsert"-ing
	//  one Audit node.
	AuditUpsertOne struct {
		create *AuditCreate
	}

	// AuditUpsert is the "OnConflict" setter.
	AuditUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Audit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(audit.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuditUpsertOne) UpdateNewValues() *AuditUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(audit.FieldID)
		}
		if _, exists := u.create.mutation.Action(); exists {
			s.SetIgnore(audit.FieldAction)
		}
		if _, exists := u.create.mutation.Author(); exists {
			s.SetIgnore(audit.FieldAuthor)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Audit.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AuditUpsertOne) Ignore() *AuditUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuditUpsertOne) DoNothing() *AuditUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuditCreate.OnConflict
// documentation for more info.
func (u *AuditUpsertOne) Update(set func(*AuditUpsert)) *AuditUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuditUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *AuditUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuditCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuditUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuditUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AuditUpsertOne.ID is not supported by MySQL driver. Use AuditUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuditUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuditCreateBulk is the builder for creating many Audit entities in bulk.
type AuditCreateBulk struct {
	config
	err      error
	builders []*AuditCreate
	conflict []sql.ConflictOption
}

// Save creates the Audit entities in the database.
func (acb *AuditCreateBulk) Save(ctx context.Context) ([]*Audit, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Audit, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuditMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AuditCreateBulk) SaveX(ctx context.Context) []*Audit {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AuditCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AuditCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Audit.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuditUpsert) {
//			SetAction(v+v).
//		}).
//		Exec(ctx)
func (acb *AuditCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuditUpsertBulk {
	acb.conflict = opts
	return &AuditUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Audit.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AuditCreateBulk) OnConflictColumns(columns ...string) *AuditUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AuditUpsertBulk{
		create: acb,
	}
}

// AuditUpsertBulk is the builder for "upsert"-ing
// a bulk of Audit nodes.
type AuditUpsertBulk struct {
	create *AuditCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Audit.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(audit.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AuditUpsertBulk) UpdateNewValues() *AuditUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(audit.FieldID)
			}
			if _, exists := b.mutation.Action(); exists {
				s.SetIgnore(audit.FieldAction)
			}
			if _, exists := b.mutation.Author(); exists {
				s.SetIgnore(audit.FieldAuthor)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Audit.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AuditUpsertBulk) Ignore() *AuditUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuditUpsertBulk) DoNothing() *AuditUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuditCreateBulk.OnConflict
// documentation for more info.
func (u *AuditUpsertBulk) Update(set func(*AuditUpsert)) *AuditUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuditUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *AuditUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuditCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuditCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuditUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}