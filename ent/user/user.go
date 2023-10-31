// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldProvider holds the string denoting the provider field in the database.
	FieldProvider = "provider"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldPhotoURL holds the string denoting the photo_url field in the database.
	FieldPhotoURL = "photo_url"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldDisabledReason holds the string denoting the disabled_reason field in the database.
	FieldDisabledReason = "disabled_reason"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeKeys holds the string denoting the keys edge name in mutations.
	EdgeKeys = "keys"
	// EdgeAudit holds the string denoting the audit edge name in mutations.
	EdgeAudit = "audit"
	// Table holds the table name of the user in the database.
	Table = "users"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "users"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_users"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "devices"
	// DevicesInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DevicesInverseTable = "devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "user_devices"
	// KeysTable is the table that holds the keys relation/edge.
	KeysTable = "api_keys"
	// KeysInverseTable is the table name for the ApiKey entity.
	// It exists in this package in order to avoid circular dependency with the "apikey" package.
	KeysInverseTable = "api_keys"
	// KeysColumn is the table column denoting the keys relation/edge.
	KeysColumn = "user_keys"
	// AuditTable is the table that holds the audit relation/edge.
	AuditTable = "audits"
	// AuditInverseTable is the table name for the Audit entity.
	// It exists in this package in order to avoid circular dependency with the "audit" package.
	AuditInverseTable = "audits"
	// AuditColumn is the table column denoting the audit relation/edge.
	AuditColumn = "user_audit"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldFirstname,
	FieldLastname,
	FieldProvider,
	FieldPassword,
	FieldSalt,
	FieldPhotoURL,
	FieldDisabled,
	FieldDisabledReason,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_users",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	FirstnameValidator func(string) error
	// LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	LastnameValidator func(string) error
	// ProviderValidator is a validator for the "provider" field. It is called by the builders before save.
	ProviderValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// SaltValidator is a validator for the "salt" field. It is called by the builders before save.
	SaltValidator func(string) error
	// DefaultDisabled holds the default value on creation for the "disabled" field.
	DefaultDisabled bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByFirstname orders the results by the firstname field.
func ByFirstname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstname, opts...).ToFunc()
}

// ByLastname orders the results by the lastname field.
func ByLastname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastname, opts...).ToFunc()
}

// ByProvider orders the results by the provider field.
func ByProvider(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProvider, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// BySalt orders the results by the salt field.
func BySalt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSalt, opts...).ToFunc()
}

// ByPhotoURL orders the results by the photo_url field.
func ByPhotoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhotoURL, opts...).ToFunc()
}

// ByDisabled orders the results by the disabled field.
func ByDisabled(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabled, opts...).ToFunc()
}

// ByDisabledReason orders the results by the disabled_reason field.
func ByDisabledReason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisabledReason, opts...).ToFunc()
}

// ByGroupField orders the results by group field.
func ByGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), sql.OrderByField(field, opts...))
	}
}

// ByDevicesCount orders the results by devices count.
func ByDevicesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDevicesStep(), opts...)
	}
}

// ByDevices orders the results by devices terms.
func ByDevices(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDevicesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByKeysCount orders the results by keys count.
func ByKeysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newKeysStep(), opts...)
	}
}

// ByKeys orders the results by keys terms.
func ByKeys(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newKeysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAuditCount orders the results by audit count.
func ByAuditCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAuditStep(), opts...)
	}
}

// ByAudit orders the results by audit terms.
func ByAudit(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAuditStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
	)
}
func newDevicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DevicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
	)
}
func newKeysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(KeysInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, KeysTable, KeysColumn),
	)
}
func newAuditStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AuditInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AuditTable, AuditColumn),
	)
}
