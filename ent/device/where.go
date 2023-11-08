// Code generated by ent, DO NOT EDIT.

package device

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/tiagoposse/connect/ent/predicate"
	"github.com/tiagoposse/connect/internal/types"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDescription, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldType, v))
}

// DNS applies equality check predicate on the "dns" field. It's identical to DNSEQ.
func DNS(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDNS, v))
}

// PublicKey applies equality check predicate on the "public_key" field. It's identical to PublicKeyEQ.
func PublicKey(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldPublicKey, v))
}

// PresharedKey applies equality check predicate on the "preshared_key" field. It's identical to PresharedKeyEQ.
func PresharedKey(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldPresharedKey, v))
}

// KeepAlive applies equality check predicate on the "keep_alive" field. It's identical to KeepAliveEQ.
func KeepAlive(v bool) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldKeepAlive, v))
}

// Endpoint applies equality check predicate on the "endpoint" field. It's identical to EndpointEQ.
func Endpoint(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldEndpoint, v))
}

// AllowedIps applies equality check predicate on the "allowed_ips" field. It's identical to AllowedIpsEQ.
func AllowedIps(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldAllowedIps, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Device {
	return predicate.Device(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Device {
	return predicate.Device(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldDescription, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldType, v))
}

// DNSEQ applies the EQ predicate on the "dns" field.
func DNSEQ(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldDNS, v))
}

// DNSNEQ applies the NEQ predicate on the "dns" field.
func DNSNEQ(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldDNS, v))
}

// DNSIn applies the In predicate on the "dns" field.
func DNSIn(vs ...types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldDNS, vs...))
}

// DNSNotIn applies the NotIn predicate on the "dns" field.
func DNSNotIn(vs ...types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldDNS, vs...))
}

// DNSGT applies the GT predicate on the "dns" field.
func DNSGT(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldDNS, v))
}

// DNSGTE applies the GTE predicate on the "dns" field.
func DNSGTE(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldDNS, v))
}

// DNSLT applies the LT predicate on the "dns" field.
func DNSLT(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldDNS, v))
}

// DNSLTE applies the LTE predicate on the "dns" field.
func DNSLTE(v types.InetSlice) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldDNS, v))
}

// PublicKeyEQ applies the EQ predicate on the "public_key" field.
func PublicKeyEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldPublicKey, v))
}

// PublicKeyNEQ applies the NEQ predicate on the "public_key" field.
func PublicKeyNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldPublicKey, v))
}

// PublicKeyIn applies the In predicate on the "public_key" field.
func PublicKeyIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldPublicKey, vs...))
}

// PublicKeyNotIn applies the NotIn predicate on the "public_key" field.
func PublicKeyNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldPublicKey, vs...))
}

// PublicKeyGT applies the GT predicate on the "public_key" field.
func PublicKeyGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldPublicKey, v))
}

// PublicKeyGTE applies the GTE predicate on the "public_key" field.
func PublicKeyGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldPublicKey, v))
}

// PublicKeyLT applies the LT predicate on the "public_key" field.
func PublicKeyLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldPublicKey, v))
}

// PublicKeyLTE applies the LTE predicate on the "public_key" field.
func PublicKeyLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldPublicKey, v))
}

// PublicKeyContains applies the Contains predicate on the "public_key" field.
func PublicKeyContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldPublicKey, v))
}

// PublicKeyHasPrefix applies the HasPrefix predicate on the "public_key" field.
func PublicKeyHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldPublicKey, v))
}

// PublicKeyHasSuffix applies the HasSuffix predicate on the "public_key" field.
func PublicKeyHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldPublicKey, v))
}

// PublicKeyEqualFold applies the EqualFold predicate on the "public_key" field.
func PublicKeyEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldPublicKey, v))
}

// PublicKeyContainsFold applies the ContainsFold predicate on the "public_key" field.
func PublicKeyContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldPublicKey, v))
}

// PresharedKeyEQ applies the EQ predicate on the "preshared_key" field.
func PresharedKeyEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldPresharedKey, v))
}

// PresharedKeyNEQ applies the NEQ predicate on the "preshared_key" field.
func PresharedKeyNEQ(v string) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldPresharedKey, v))
}

// PresharedKeyIn applies the In predicate on the "preshared_key" field.
func PresharedKeyIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldPresharedKey, vs...))
}

// PresharedKeyNotIn applies the NotIn predicate on the "preshared_key" field.
func PresharedKeyNotIn(vs ...string) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldPresharedKey, vs...))
}

// PresharedKeyGT applies the GT predicate on the "preshared_key" field.
func PresharedKeyGT(v string) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldPresharedKey, v))
}

// PresharedKeyGTE applies the GTE predicate on the "preshared_key" field.
func PresharedKeyGTE(v string) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldPresharedKey, v))
}

// PresharedKeyLT applies the LT predicate on the "preshared_key" field.
func PresharedKeyLT(v string) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldPresharedKey, v))
}

// PresharedKeyLTE applies the LTE predicate on the "preshared_key" field.
func PresharedKeyLTE(v string) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldPresharedKey, v))
}

// PresharedKeyContains applies the Contains predicate on the "preshared_key" field.
func PresharedKeyContains(v string) predicate.Device {
	return predicate.Device(sql.FieldContains(FieldPresharedKey, v))
}

// PresharedKeyHasPrefix applies the HasPrefix predicate on the "preshared_key" field.
func PresharedKeyHasPrefix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasPrefix(FieldPresharedKey, v))
}

// PresharedKeyHasSuffix applies the HasSuffix predicate on the "preshared_key" field.
func PresharedKeyHasSuffix(v string) predicate.Device {
	return predicate.Device(sql.FieldHasSuffix(FieldPresharedKey, v))
}

// PresharedKeyEqualFold applies the EqualFold predicate on the "preshared_key" field.
func PresharedKeyEqualFold(v string) predicate.Device {
	return predicate.Device(sql.FieldEqualFold(FieldPresharedKey, v))
}

// PresharedKeyContainsFold applies the ContainsFold predicate on the "preshared_key" field.
func PresharedKeyContainsFold(v string) predicate.Device {
	return predicate.Device(sql.FieldContainsFold(FieldPresharedKey, v))
}

// KeepAliveEQ applies the EQ predicate on the "keep_alive" field.
func KeepAliveEQ(v bool) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldKeepAlive, v))
}

// KeepAliveNEQ applies the NEQ predicate on the "keep_alive" field.
func KeepAliveNEQ(v bool) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldKeepAlive, v))
}

// EndpointEQ applies the EQ predicate on the "endpoint" field.
func EndpointEQ(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldEndpoint, v))
}

// EndpointNEQ applies the NEQ predicate on the "endpoint" field.
func EndpointNEQ(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldEndpoint, v))
}

// EndpointIn applies the In predicate on the "endpoint" field.
func EndpointIn(vs ...types.Inet) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldEndpoint, vs...))
}

// EndpointNotIn applies the NotIn predicate on the "endpoint" field.
func EndpointNotIn(vs ...types.Inet) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldEndpoint, vs...))
}

// EndpointGT applies the GT predicate on the "endpoint" field.
func EndpointGT(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldEndpoint, v))
}

// EndpointGTE applies the GTE predicate on the "endpoint" field.
func EndpointGTE(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldEndpoint, v))
}

// EndpointLT applies the LT predicate on the "endpoint" field.
func EndpointLT(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldEndpoint, v))
}

// EndpointLTE applies the LTE predicate on the "endpoint" field.
func EndpointLTE(v types.Inet) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldEndpoint, v))
}

// EndpointContains applies the Contains predicate on the "endpoint" field.
func EndpointContains(v types.Inet) predicate.Device {
	vc := v.String()
	return predicate.Device(sql.FieldContains(FieldEndpoint, vc))
}

// EndpointHasPrefix applies the HasPrefix predicate on the "endpoint" field.
func EndpointHasPrefix(v types.Inet) predicate.Device {
	vc := v.String()
	return predicate.Device(sql.FieldHasPrefix(FieldEndpoint, vc))
}

// EndpointHasSuffix applies the HasSuffix predicate on the "endpoint" field.
func EndpointHasSuffix(v types.Inet) predicate.Device {
	vc := v.String()
	return predicate.Device(sql.FieldHasSuffix(FieldEndpoint, vc))
}

// EndpointEqualFold applies the EqualFold predicate on the "endpoint" field.
func EndpointEqualFold(v types.Inet) predicate.Device {
	vc := v.String()
	return predicate.Device(sql.FieldEqualFold(FieldEndpoint, vc))
}

// EndpointContainsFold applies the ContainsFold predicate on the "endpoint" field.
func EndpointContainsFold(v types.Inet) predicate.Device {
	vc := v.String()
	return predicate.Device(sql.FieldContainsFold(FieldEndpoint, vc))
}

// AllowedIpsEQ applies the EQ predicate on the "allowed_ips" field.
func AllowedIpsEQ(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldEQ(FieldAllowedIps, v))
}

// AllowedIpsNEQ applies the NEQ predicate on the "allowed_ips" field.
func AllowedIpsNEQ(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldNEQ(FieldAllowedIps, v))
}

// AllowedIpsIn applies the In predicate on the "allowed_ips" field.
func AllowedIpsIn(vs ...types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldIn(FieldAllowedIps, vs...))
}

// AllowedIpsNotIn applies the NotIn predicate on the "allowed_ips" field.
func AllowedIpsNotIn(vs ...types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldNotIn(FieldAllowedIps, vs...))
}

// AllowedIpsGT applies the GT predicate on the "allowed_ips" field.
func AllowedIpsGT(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldGT(FieldAllowedIps, v))
}

// AllowedIpsGTE applies the GTE predicate on the "allowed_ips" field.
func AllowedIpsGTE(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldGTE(FieldAllowedIps, v))
}

// AllowedIpsLT applies the LT predicate on the "allowed_ips" field.
func AllowedIpsLT(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldLT(FieldAllowedIps, v))
}

// AllowedIpsLTE applies the LTE predicate on the "allowed_ips" field.
func AllowedIpsLTE(v types.CidrSlice) predicate.Device {
	return predicate.Device(sql.FieldLTE(FieldAllowedIps, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Device {
	return predicate.Device(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Device) predicate.Device {
	return predicate.Device(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Device) predicate.Device {
	return predicate.Device(sql.NotPredicates(p))
}