// Code generated by ent, DO NOT EDIT.

package article

import (
	"blug/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitle, v))
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldDesc, v))
}

// Category applies equality check predicate on the "category" field. It's identical to CategoryEQ.
func Category(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCategory, v))
}

// Tags applies equality check predicate on the "tags" field. It's identical to TagsEQ.
func Tags(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTags, v))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldURL, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreateTime, v))
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v bool) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldIsShow, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldContent, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTitle, v))
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldDesc, v))
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldDesc, v))
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldDesc, vs...))
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldDesc, vs...))
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldDesc, v))
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldDesc, v))
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldDesc, v))
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldDesc, v))
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldDesc, v))
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldDesc, v))
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldDesc, v))
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldDesc, v))
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldDesc, v))
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCategory, v))
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCategory, v))
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCategory, vs...))
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCategory, vs...))
}

// CategoryGT applies the GT predicate on the "category" field.
func CategoryGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCategory, v))
}

// CategoryGTE applies the GTE predicate on the "category" field.
func CategoryGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCategory, v))
}

// CategoryLT applies the LT predicate on the "category" field.
func CategoryLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCategory, v))
}

// CategoryLTE applies the LTE predicate on the "category" field.
func CategoryLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCategory, v))
}

// CategoryContains applies the Contains predicate on the "category" field.
func CategoryContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldCategory, v))
}

// CategoryHasPrefix applies the HasPrefix predicate on the "category" field.
func CategoryHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldCategory, v))
}

// CategoryHasSuffix applies the HasSuffix predicate on the "category" field.
func CategoryHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldCategory, v))
}

// CategoryEqualFold applies the EqualFold predicate on the "category" field.
func CategoryEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldCategory, v))
}

// CategoryContainsFold applies the ContainsFold predicate on the "category" field.
func CategoryContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldCategory, v))
}

// TagsEQ applies the EQ predicate on the "tags" field.
func TagsEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldTags, v))
}

// TagsNEQ applies the NEQ predicate on the "tags" field.
func TagsNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldTags, v))
}

// TagsIn applies the In predicate on the "tags" field.
func TagsIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldTags, vs...))
}

// TagsNotIn applies the NotIn predicate on the "tags" field.
func TagsNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldTags, vs...))
}

// TagsGT applies the GT predicate on the "tags" field.
func TagsGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldTags, v))
}

// TagsGTE applies the GTE predicate on the "tags" field.
func TagsGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldTags, v))
}

// TagsLT applies the LT predicate on the "tags" field.
func TagsLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldTags, v))
}

// TagsLTE applies the LTE predicate on the "tags" field.
func TagsLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldTags, v))
}

// TagsContains applies the Contains predicate on the "tags" field.
func TagsContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldTags, v))
}

// TagsHasPrefix applies the HasPrefix predicate on the "tags" field.
func TagsHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldTags, v))
}

// TagsHasSuffix applies the HasSuffix predicate on the "tags" field.
func TagsHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldTags, v))
}

// TagsEqualFold applies the EqualFold predicate on the "tags" field.
func TagsEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldTags, v))
}

// TagsContainsFold applies the ContainsFold predicate on the "tags" field.
func TagsContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldTags, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldURL, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeContains applies the Contains predicate on the "create_time" field.
func CreateTimeContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldCreateTime, v))
}

// CreateTimeHasPrefix applies the HasPrefix predicate on the "create_time" field.
func CreateTimeHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldCreateTime, v))
}

// CreateTimeHasSuffix applies the HasSuffix predicate on the "create_time" field.
func CreateTimeHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldCreateTime, v))
}

// CreateTimeEqualFold applies the EqualFold predicate on the "create_time" field.
func CreateTimeEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldCreateTime, v))
}

// CreateTimeContainsFold applies the ContainsFold predicate on the "create_time" field.
func CreateTimeContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldCreateTime, v))
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v bool) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldIsShow, v))
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v bool) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldIsShow, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Article {
	return predicate.Article(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Article {
	return predicate.Article(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Article {
	return predicate.Article(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Article {
	return predicate.Article(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Article {
	return predicate.Article(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Article {
	return predicate.Article(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Article {
	return predicate.Article(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Article {
	return predicate.Article(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Article {
	return predicate.Article(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Article {
	return predicate.Article(sql.FieldContainsFold(FieldContent, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Article) predicate.Article {
	return predicate.Article(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Article) predicate.Article {
	return predicate.Article(sql.NotPredicates(p))
}
