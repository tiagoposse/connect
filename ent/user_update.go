// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/ent/apikey"
	"github.com/tiagoposse/connect/ent/audit"
	"github.com/tiagoposse/connect/ent/device"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/predicate"
	"github.com/tiagoposse/connect/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetFirstname sets the "firstname" field.
func (uu *UserUpdate) SetFirstname(s string) *UserUpdate {
	uu.mutation.SetFirstname(s)
	return uu
}

// SetLastname sets the "lastname" field.
func (uu *UserUpdate) SetLastname(s string) *UserUpdate {
	uu.mutation.SetLastname(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// SetSalt sets the "salt" field.
func (uu *UserUpdate) SetSalt(s string) *UserUpdate {
	uu.mutation.SetSalt(s)
	return uu
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSalt(s *string) *UserUpdate {
	if s != nil {
		uu.SetSalt(*s)
	}
	return uu
}

// ClearSalt clears the value of the "salt" field.
func (uu *UserUpdate) ClearSalt() *UserUpdate {
	uu.mutation.ClearSalt()
	return uu
}

// SetPhotoURL sets the "photo_url" field.
func (uu *UserUpdate) SetPhotoURL(s string) *UserUpdate {
	uu.mutation.SetPhotoURL(s)
	return uu
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhotoURL(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhotoURL(*s)
	}
	return uu
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (uu *UserUpdate) ClearPhotoURL() *UserUpdate {
	uu.mutation.ClearPhotoURL()
	return uu
}

// SetDisabled sets the "disabled" field.
func (uu *UserUpdate) SetDisabled(b bool) *UserUpdate {
	uu.mutation.SetDisabled(b)
	return uu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDisabled(b *bool) *UserUpdate {
	if b != nil {
		uu.SetDisabled(*b)
	}
	return uu
}

// SetDisabledReason sets the "disabled_reason" field.
func (uu *UserUpdate) SetDisabledReason(s string) *UserUpdate {
	uu.mutation.SetDisabledReason(s)
	return uu
}

// SetNillableDisabledReason sets the "disabled_reason" field if the given value is not nil.
func (uu *UserUpdate) SetNillableDisabledReason(s *string) *UserUpdate {
	if s != nil {
		uu.SetDisabledReason(*s)
	}
	return uu
}

// ClearDisabledReason clears the value of the "disabled_reason" field.
func (uu *UserUpdate) ClearDisabledReason() *UserUpdate {
	uu.mutation.ClearDisabledReason()
	return uu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uu *UserUpdate) SetGroupID(id string) *UserUpdate {
	uu.mutation.SetGroupID(id)
	return uu
}

// SetGroup sets the "group" edge to the Group entity.
func (uu *UserUpdate) SetGroup(g *Group) *UserUpdate {
	return uu.SetGroupID(g.ID)
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uu *UserUpdate) AddDeviceIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.AddDeviceIDs(ids...)
	return uu
}

// AddDevices adds the "devices" edges to the Device entity.
func (uu *UserUpdate) AddDevices(d ...*Device) *UserUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.AddDeviceIDs(ids...)
}

// AddKeyIDs adds the "keys" edge to the ApiKey entity by IDs.
func (uu *UserUpdate) AddKeyIDs(ids ...int) *UserUpdate {
	uu.mutation.AddKeyIDs(ids...)
	return uu
}

// AddKeys adds the "keys" edges to the ApiKey entity.
func (uu *UserUpdate) AddKeys(a ...*ApiKey) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddKeyIDs(ids...)
}

// AddAuditIDs adds the "audit" edge to the Audit entity by IDs.
func (uu *UserUpdate) AddAuditIDs(ids ...string) *UserUpdate {
	uu.mutation.AddAuditIDs(ids...)
	return uu
}

// AddAudit adds the "audit" edges to the Audit entity.
func (uu *UserUpdate) AddAudit(a ...*Audit) *UserUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAuditIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (uu *UserUpdate) ClearGroup() *UserUpdate {
	uu.mutation.ClearGroup()
	return uu
}

// ClearDevices clears all "devices" edges to the Device entity.
func (uu *UserUpdate) ClearDevices() *UserUpdate {
	uu.mutation.ClearDevices()
	return uu
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (uu *UserUpdate) RemoveDeviceIDs(ids ...uuid.UUID) *UserUpdate {
	uu.mutation.RemoveDeviceIDs(ids...)
	return uu
}

// RemoveDevices removes "devices" edges to Device entities.
func (uu *UserUpdate) RemoveDevices(d ...*Device) *UserUpdate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.RemoveDeviceIDs(ids...)
}

// ClearKeys clears all "keys" edges to the ApiKey entity.
func (uu *UserUpdate) ClearKeys() *UserUpdate {
	uu.mutation.ClearKeys()
	return uu
}

// RemoveKeyIDs removes the "keys" edge to ApiKey entities by IDs.
func (uu *UserUpdate) RemoveKeyIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveKeyIDs(ids...)
	return uu
}

// RemoveKeys removes "keys" edges to ApiKey entities.
func (uu *UserUpdate) RemoveKeys(a ...*ApiKey) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveKeyIDs(ids...)
}

// ClearAudit clears all "audit" edges to the Audit entity.
func (uu *UserUpdate) ClearAudit() *UserUpdate {
	uu.mutation.ClearAudit()
	return uu
}

// RemoveAuditIDs removes the "audit" edge to Audit entities by IDs.
func (uu *UserUpdate) RemoveAuditIDs(ids ...string) *UserUpdate {
	uu.mutation.RemoveAuditIDs(ids...)
	return uu
}

// RemoveAudit removes "audit" edges to Audit entities.
func (uu *UserUpdate) RemoveAudit(a ...*Audit) *UserUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAuditIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf(`ent: validator failed for field "User.firstname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf(`ent: validator failed for field "User.lastname": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Salt(); ok {
		if err := user.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User.salt": %w`, err)}
		}
	}
	if _, ok := uu.mutation.GroupID(); uu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.group"`)
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Firstname(); ok {
		_spec.SetField(user.FieldFirstname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Lastname(); ok {
		_spec.SetField(user.FieldLastname, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uu.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
	}
	if uu.mutation.SaltCleared() {
		_spec.ClearField(user.FieldSalt, field.TypeString)
	}
	if value, ok := uu.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
	}
	if uu.mutation.PhotoURLCleared() {
		_spec.ClearField(user.FieldPhotoURL, field.TypeString)
	}
	if value, ok := uu.mutation.Disabled(); ok {
		_spec.SetField(user.FieldDisabled, field.TypeBool, value)
	}
	if value, ok := uu.mutation.DisabledReason(); ok {
		_spec.SetField(user.FieldDisabledReason, field.TypeString, value)
	}
	if uu.mutation.DisabledReasonCleared() {
		_spec.ClearField(user.FieldDisabledReason, field.TypeString)
	}
	if uu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !uu.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedKeysIDs(); len(nodes) > 0 && !uu.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.KeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AuditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditTable,
			Columns: []string{user.AuditColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAuditIDs(); len(nodes) > 0 && !uu.mutation.AuditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditTable,
			Columns: []string{user.AuditColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AuditIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditTable,
			Columns: []string{user.AuditColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetFirstname sets the "firstname" field.
func (uuo *UserUpdateOne) SetFirstname(s string) *UserUpdateOne {
	uuo.mutation.SetFirstname(s)
	return uuo
}

// SetLastname sets the "lastname" field.
func (uuo *UserUpdateOne) SetLastname(s string) *UserUpdateOne {
	uuo.mutation.SetLastname(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// SetSalt sets the "salt" field.
func (uuo *UserUpdateOne) SetSalt(s string) *UserUpdateOne {
	uuo.mutation.SetSalt(s)
	return uuo
}

// SetNillableSalt sets the "salt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSalt(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSalt(*s)
	}
	return uuo
}

// ClearSalt clears the value of the "salt" field.
func (uuo *UserUpdateOne) ClearSalt() *UserUpdateOne {
	uuo.mutation.ClearSalt()
	return uuo
}

// SetPhotoURL sets the "photo_url" field.
func (uuo *UserUpdateOne) SetPhotoURL(s string) *UserUpdateOne {
	uuo.mutation.SetPhotoURL(s)
	return uuo
}

// SetNillablePhotoURL sets the "photo_url" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhotoURL(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhotoURL(*s)
	}
	return uuo
}

// ClearPhotoURL clears the value of the "photo_url" field.
func (uuo *UserUpdateOne) ClearPhotoURL() *UserUpdateOne {
	uuo.mutation.ClearPhotoURL()
	return uuo
}

// SetDisabled sets the "disabled" field.
func (uuo *UserUpdateOne) SetDisabled(b bool) *UserUpdateOne {
	uuo.mutation.SetDisabled(b)
	return uuo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDisabled(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetDisabled(*b)
	}
	return uuo
}

// SetDisabledReason sets the "disabled_reason" field.
func (uuo *UserUpdateOne) SetDisabledReason(s string) *UserUpdateOne {
	uuo.mutation.SetDisabledReason(s)
	return uuo
}

// SetNillableDisabledReason sets the "disabled_reason" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableDisabledReason(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetDisabledReason(*s)
	}
	return uuo
}

// ClearDisabledReason clears the value of the "disabled_reason" field.
func (uuo *UserUpdateOne) ClearDisabledReason() *UserUpdateOne {
	uuo.mutation.ClearDisabledReason()
	return uuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (uuo *UserUpdateOne) SetGroupID(id string) *UserUpdateOne {
	uuo.mutation.SetGroupID(id)
	return uuo
}

// SetGroup sets the "group" edge to the Group entity.
func (uuo *UserUpdateOne) SetGroup(g *Group) *UserUpdateOne {
	return uuo.SetGroupID(g.ID)
}

// AddDeviceIDs adds the "devices" edge to the Device entity by IDs.
func (uuo *UserUpdateOne) AddDeviceIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.AddDeviceIDs(ids...)
	return uuo
}

// AddDevices adds the "devices" edges to the Device entity.
func (uuo *UserUpdateOne) AddDevices(d ...*Device) *UserUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.AddDeviceIDs(ids...)
}

// AddKeyIDs adds the "keys" edge to the ApiKey entity by IDs.
func (uuo *UserUpdateOne) AddKeyIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddKeyIDs(ids...)
	return uuo
}

// AddKeys adds the "keys" edges to the ApiKey entity.
func (uuo *UserUpdateOne) AddKeys(a ...*ApiKey) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddKeyIDs(ids...)
}

// AddAuditIDs adds the "audit" edge to the Audit entity by IDs.
func (uuo *UserUpdateOne) AddAuditIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.AddAuditIDs(ids...)
	return uuo
}

// AddAudit adds the "audit" edges to the Audit entity.
func (uuo *UserUpdateOne) AddAudit(a ...*Audit) *UserUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAuditIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (uuo *UserUpdateOne) ClearGroup() *UserUpdateOne {
	uuo.mutation.ClearGroup()
	return uuo
}

// ClearDevices clears all "devices" edges to the Device entity.
func (uuo *UserUpdateOne) ClearDevices() *UserUpdateOne {
	uuo.mutation.ClearDevices()
	return uuo
}

// RemoveDeviceIDs removes the "devices" edge to Device entities by IDs.
func (uuo *UserUpdateOne) RemoveDeviceIDs(ids ...uuid.UUID) *UserUpdateOne {
	uuo.mutation.RemoveDeviceIDs(ids...)
	return uuo
}

// RemoveDevices removes "devices" edges to Device entities.
func (uuo *UserUpdateOne) RemoveDevices(d ...*Device) *UserUpdateOne {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.RemoveDeviceIDs(ids...)
}

// ClearKeys clears all "keys" edges to the ApiKey entity.
func (uuo *UserUpdateOne) ClearKeys() *UserUpdateOne {
	uuo.mutation.ClearKeys()
	return uuo
}

// RemoveKeyIDs removes the "keys" edge to ApiKey entities by IDs.
func (uuo *UserUpdateOne) RemoveKeyIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveKeyIDs(ids...)
	return uuo
}

// RemoveKeys removes "keys" edges to ApiKey entities.
func (uuo *UserUpdateOne) RemoveKeys(a ...*ApiKey) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveKeyIDs(ids...)
}

// ClearAudit clears all "audit" edges to the Audit entity.
func (uuo *UserUpdateOne) ClearAudit() *UserUpdateOne {
	uuo.mutation.ClearAudit()
	return uuo
}

// RemoveAuditIDs removes the "audit" edge to Audit entities by IDs.
func (uuo *UserUpdateOne) RemoveAuditIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.RemoveAuditIDs(ids...)
	return uuo
}

// RemoveAudit removes "audit" edges to Audit entities.
func (uuo *UserUpdateOne) RemoveAudit(a ...*Audit) *UserUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAuditIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Firstname(); ok {
		if err := user.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "firstname", err: fmt.Errorf(`ent: validator failed for field "User.firstname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Lastname(); ok {
		if err := user.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "lastname", err: fmt.Errorf(`ent: validator failed for field "User.lastname": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Password(); ok {
		if err := user.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "User.password": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Salt(); ok {
		if err := user.SaltValidator(v); err != nil {
			return &ValidationError{Name: "salt", err: fmt.Errorf(`ent: validator failed for field "User.salt": %w`, err)}
		}
	}
	if _, ok := uuo.mutation.GroupID(); uuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "User.group"`)
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Firstname(); ok {
		_spec.SetField(user.FieldFirstname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Lastname(); ok {
		_spec.SetField(user.FieldLastname, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uuo.mutation.Salt(); ok {
		_spec.SetField(user.FieldSalt, field.TypeString, value)
	}
	if uuo.mutation.SaltCleared() {
		_spec.ClearField(user.FieldSalt, field.TypeString)
	}
	if value, ok := uuo.mutation.PhotoURL(); ok {
		_spec.SetField(user.FieldPhotoURL, field.TypeString, value)
	}
	if uuo.mutation.PhotoURLCleared() {
		_spec.ClearField(user.FieldPhotoURL, field.TypeString)
	}
	if value, ok := uuo.mutation.Disabled(); ok {
		_spec.SetField(user.FieldDisabled, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.DisabledReason(); ok {
		_spec.SetField(user.FieldDisabledReason, field.TypeString, value)
	}
	if uuo.mutation.DisabledReasonCleared() {
		_spec.ClearField(user.FieldDisabledReason, field.TypeString)
	}
	if uuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.GroupTable,
			Columns: []string{user.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedDevicesIDs(); len(nodes) > 0 && !uuo.mutation.DevicesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DevicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DevicesTable,
			Columns: []string{user.DevicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedKeysIDs(); len(nodes) > 0 && !uuo.mutation.KeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.KeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.KeysTable,
			Columns: []string{user.KeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AuditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditTable,
			Columns: []string{user.AuditColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAuditIDs(); len(nodes) > 0 && !uuo.mutation.AuditCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditTable,
			Columns: []string{user.AuditColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AuditIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuditTable,
			Columns: []string{user.AuditColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(audit.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
