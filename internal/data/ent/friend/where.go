// Code generated by ent, DO NOT EDIT.

package friend

import (
	"blug/internal/data/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "Title" field. It's identical to TitleEQ.
func Title(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldTitle, v))
}

// Desc applies equality check predicate on the "Desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldDesc, v))
}

// Link applies equality check predicate on the "Link" field. It's identical to LinkEQ.
func Link(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldLink, v))
}

// Avatar applies equality check predicate on the "Avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldAvatar, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldCreateTime, v))
}

// TitleEQ applies the EQ predicate on the "Title" field.
func TitleEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "Title" field.
func TitleNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "Title" field.
func TitleIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "Title" field.
func TitleNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "Title" field.
func TitleGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "Title" field.
func TitleGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "Title" field.
func TitleLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "Title" field.
func TitleLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "Title" field.
func TitleContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "Title" field.
func TitleHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "Title" field.
func TitleHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "Title" field.
func TitleEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "Title" field.
func TitleContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldTitle, v))
}

// DescEQ applies the EQ predicate on the "Desc" field.
func DescEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "Desc" field.
func DescNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "Desc" field.
func DescIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "Desc" field.
func DescNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "Desc" field.
func DescGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "Desc" field.
func DescGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "Desc" field.
func DescLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "Desc" field.
func DescLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "Desc" field.
func DescContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "Desc" field.
func DescHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "Desc" field.
func DescHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "Desc" field.
func DescEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "Desc" field.
func DescContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldDesc, v))
}

// LinkEQ applies the EQ predicate on the "Link" field.
func LinkEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "Link" field.
func LinkNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "Link" field.
func LinkIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "Link" field.
func LinkNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "Link" field.
func LinkGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "Link" field.
func LinkGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "Link" field.
func LinkLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "Link" field.
func LinkLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "Link" field.
func LinkContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "Link" field.
func LinkHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "Link" field.
func LinkHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "Link" field.
func LinkEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "Link" field.
func LinkContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldLink, v))
}

// AvatarEQ applies the EQ predicate on the "Avatar" field.
func AvatarEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "Avatar" field.
func AvatarNEQ(v string) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "Avatar" field.
func AvatarIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "Avatar" field.
func AvatarNotIn(vs ...string) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "Avatar" field.
func AvatarGT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "Avatar" field.
func AvatarGTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "Avatar" field.
func AvatarLT(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "Avatar" field.
func AvatarLTE(v string) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "Avatar" field.
func AvatarContains(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "Avatar" field.
func AvatarHasPrefix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "Avatar" field.
func AvatarHasSuffix(v string) predicate.Friend {
	return predicate.Friend(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarEqualFold applies the EqualFold predicate on the "Avatar" field.
func AvatarEqualFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "Avatar" field.
func AvatarContainsFold(v string) predicate.Friend {
	return predicate.Friend(sql.FieldContainsFold(FieldAvatar, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Friend {
	return predicate.Friend(sql.FieldLTE(FieldCreateTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Friend) predicate.Friend {
	return predicate.Friend(sql.NotPredicates(p))
}
