// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Firstname holds the value of the "firstname" field.
	Firstname string `json:"firstname,omitempty"`
	// Lastname holds the value of the "lastname" field.
	Lastname string `json:"lastname,omitempty"`
	// Provider holds the value of the "provider" field.
	Provider string `json:"provider,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// Salt holds the value of the "salt" field.
	Salt string `json:"-"`
	// PhotoURL holds the value of the "photo_url" field.
	PhotoURL string `json:"photo_url,omitempty"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty"`
	// DisabledReason holds the value of the "disabled_reason" field.
	DisabledReason string `json:"disabled_reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	group_users  *string
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Devices holds the value of the devices edge.
	Devices []*Device `json:"devices,omitempty"`
	// Keys holds the value of the keys edge.
	Keys []*ApiKey `json:"keys,omitempty"`
	// Audit holds the value of the audit edge.
	Audit []*Audit `json:"audit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes  [4]bool
	namedDevices map[string][]*Device
	namedKeys    map[string][]*ApiKey
	namedAudit   map[string][]*Audit
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// DevicesOrErr returns the Devices value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) DevicesOrErr() ([]*Device, error) {
	if e.loadedTypes[1] {
		return e.Devices, nil
	}
	return nil, &NotLoadedError{edge: "devices"}
}

// KeysOrErr returns the Keys value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) KeysOrErr() ([]*ApiKey, error) {
	if e.loadedTypes[2] {
		return e.Keys, nil
	}
	return nil, &NotLoadedError{edge: "keys"}
}

// AuditOrErr returns the Audit value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AuditOrErr() ([]*Audit, error) {
	if e.loadedTypes[3] {
		return e.Audit, nil
	}
	return nil, &NotLoadedError{edge: "audit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldDisabled:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldEmail, user.FieldFirstname, user.FieldLastname, user.FieldProvider, user.FieldPassword, user.FieldSalt, user.FieldPhotoURL, user.FieldDisabledReason:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // group_users
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				u.ID = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldFirstname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field firstname", values[i])
			} else if value.Valid {
				u.Firstname = value.String
			}
		case user.FieldLastname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lastname", values[i])
			} else if value.Valid {
				u.Lastname = value.String
			}
		case user.FieldProvider:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field provider", values[i])
			} else if value.Valid {
				u.Provider = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldSalt:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field salt", values[i])
			} else if value.Valid {
				u.Salt = value.String
			}
		case user.FieldPhotoURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field photo_url", values[i])
			} else if value.Valid {
				u.PhotoURL = value.String
			}
		case user.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				u.Disabled = value.Bool
			}
		case user.FieldDisabledReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field disabled_reason", values[i])
			} else if value.Valid {
				u.DisabledReason = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_users", values[i])
			} else if value.Valid {
				u.group_users = new(string)
				*u.group_users = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the User entity.
func (u *User) QueryGroup() *GroupQuery {
	return NewUserClient(u.config).QueryGroup(u)
}

// QueryDevices queries the "devices" edge of the User entity.
func (u *User) QueryDevices() *DeviceQuery {
	return NewUserClient(u.config).QueryDevices(u)
}

// QueryKeys queries the "keys" edge of the User entity.
func (u *User) QueryKeys() *ApiKeyQuery {
	return NewUserClient(u.config).QueryKeys(u)
}

// QueryAudit queries the "audit" edge of the User entity.
func (u *User) QueryAudit() *AuditQuery {
	return NewUserClient(u.config).QueryAudit(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("firstname=")
	builder.WriteString(u.Firstname)
	builder.WriteString(", ")
	builder.WriteString("lastname=")
	builder.WriteString(u.Lastname)
	builder.WriteString(", ")
	builder.WriteString("provider=")
	builder.WriteString(u.Provider)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("salt=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("photo_url=")
	builder.WriteString(u.PhotoURL)
	builder.WriteString(", ")
	builder.WriteString("disabled=")
	builder.WriteString(fmt.Sprintf("%v", u.Disabled))
	builder.WriteString(", ")
	builder.WriteString("disabled_reason=")
	builder.WriteString(u.DisabledReason)
	builder.WriteByte(')')
	return builder.String()
}

// NamedDevices returns the Devices named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedDevices(name string) ([]*Device, error) {
	if u.Edges.namedDevices == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedDevices[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedDevices(name string, edges ...*Device) {
	if u.Edges.namedDevices == nil {
		u.Edges.namedDevices = make(map[string][]*Device)
	}
	if len(edges) == 0 {
		u.Edges.namedDevices[name] = []*Device{}
	} else {
		u.Edges.namedDevices[name] = append(u.Edges.namedDevices[name], edges...)
	}
}

// NamedKeys returns the Keys named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedKeys(name string) ([]*ApiKey, error) {
	if u.Edges.namedKeys == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedKeys[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedKeys(name string, edges ...*ApiKey) {
	if u.Edges.namedKeys == nil {
		u.Edges.namedKeys = make(map[string][]*ApiKey)
	}
	if len(edges) == 0 {
		u.Edges.namedKeys[name] = []*ApiKey{}
	} else {
		u.Edges.namedKeys[name] = append(u.Edges.namedKeys[name], edges...)
	}
}

// NamedAudit returns the Audit named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedAudit(name string) ([]*Audit, error) {
	if u.Edges.namedAudit == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedAudit[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedAudit(name string, edges ...*Audit) {
	if u.Edges.namedAudit == nil {
		u.Edges.namedAudit = make(map[string][]*Audit)
	}
	if len(edges) == 0 {
		u.Edges.namedAudit[name] = []*Audit{}
	} else {
		u.Edges.namedAudit[name] = append(u.Edges.namedAudit[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User