// Code generated by ent, DO NOT EDIT.

package task

import (
	"blug/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskID, v))
}

// TaskName applies equality check predicate on the "task_name" field. It's identical to TaskNameEQ.
func TaskName(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskName, v))
}

// TaskType applies equality check predicate on the "task_type" field. It's identical to TaskTypeEQ.
func TaskType(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskType, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// Reason applies equality check predicate on the "reason" field. It's identical to ReasonEQ.
func Reason(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldReason, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateTime, v))
}

// FinishTime applies equality check predicate on the "finish_time" field. It's identical to FinishTimeEQ.
func FinishTime(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldFinishTime, v))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTaskID, vs...))
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTaskID, v))
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTaskID, v))
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTaskID, v))
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTaskID, v))
}

// TaskIDContains applies the Contains predicate on the "task_id" field.
func TaskIDContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTaskID, v))
}

// TaskIDHasPrefix applies the HasPrefix predicate on the "task_id" field.
func TaskIDHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTaskID, v))
}

// TaskIDHasSuffix applies the HasSuffix predicate on the "task_id" field.
func TaskIDHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTaskID, v))
}

// TaskIDEqualFold applies the EqualFold predicate on the "task_id" field.
func TaskIDEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTaskID, v))
}

// TaskIDContainsFold applies the ContainsFold predicate on the "task_id" field.
func TaskIDContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTaskID, v))
}

// TaskNameEQ applies the EQ predicate on the "task_name" field.
func TaskNameEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskName, v))
}

// TaskNameNEQ applies the NEQ predicate on the "task_name" field.
func TaskNameNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTaskName, v))
}

// TaskNameIn applies the In predicate on the "task_name" field.
func TaskNameIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTaskName, vs...))
}

// TaskNameNotIn applies the NotIn predicate on the "task_name" field.
func TaskNameNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTaskName, vs...))
}

// TaskNameGT applies the GT predicate on the "task_name" field.
func TaskNameGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTaskName, v))
}

// TaskNameGTE applies the GTE predicate on the "task_name" field.
func TaskNameGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTaskName, v))
}

// TaskNameLT applies the LT predicate on the "task_name" field.
func TaskNameLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTaskName, v))
}

// TaskNameLTE applies the LTE predicate on the "task_name" field.
func TaskNameLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTaskName, v))
}

// TaskNameContains applies the Contains predicate on the "task_name" field.
func TaskNameContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTaskName, v))
}

// TaskNameHasPrefix applies the HasPrefix predicate on the "task_name" field.
func TaskNameHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTaskName, v))
}

// TaskNameHasSuffix applies the HasSuffix predicate on the "task_name" field.
func TaskNameHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTaskName, v))
}

// TaskNameEqualFold applies the EqualFold predicate on the "task_name" field.
func TaskNameEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTaskName, v))
}

// TaskNameContainsFold applies the ContainsFold predicate on the "task_name" field.
func TaskNameContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTaskName, v))
}

// TaskTypeEQ applies the EQ predicate on the "task_type" field.
func TaskTypeEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldTaskType, v))
}

// TaskTypeNEQ applies the NEQ predicate on the "task_type" field.
func TaskTypeNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldTaskType, v))
}

// TaskTypeIn applies the In predicate on the "task_type" field.
func TaskTypeIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldTaskType, vs...))
}

// TaskTypeNotIn applies the NotIn predicate on the "task_type" field.
func TaskTypeNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldTaskType, vs...))
}

// TaskTypeGT applies the GT predicate on the "task_type" field.
func TaskTypeGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldTaskType, v))
}

// TaskTypeGTE applies the GTE predicate on the "task_type" field.
func TaskTypeGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldTaskType, v))
}

// TaskTypeLT applies the LT predicate on the "task_type" field.
func TaskTypeLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldTaskType, v))
}

// TaskTypeLTE applies the LTE predicate on the "task_type" field.
func TaskTypeLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldTaskType, v))
}

// TaskTypeContains applies the Contains predicate on the "task_type" field.
func TaskTypeContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldTaskType, v))
}

// TaskTypeHasPrefix applies the HasPrefix predicate on the "task_type" field.
func TaskTypeHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldTaskType, v))
}

// TaskTypeHasSuffix applies the HasSuffix predicate on the "task_type" field.
func TaskTypeHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldTaskType, v))
}

// TaskTypeEqualFold applies the EqualFold predicate on the "task_type" field.
func TaskTypeEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldTaskType, v))
}

// TaskTypeContainsFold applies the ContainsFold predicate on the "task_type" field.
func TaskTypeContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldTaskType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldStatus, v))
}

// ReasonEQ applies the EQ predicate on the "reason" field.
func ReasonEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldReason, v))
}

// ReasonNEQ applies the NEQ predicate on the "reason" field.
func ReasonNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldReason, v))
}

// ReasonIn applies the In predicate on the "reason" field.
func ReasonIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldReason, vs...))
}

// ReasonNotIn applies the NotIn predicate on the "reason" field.
func ReasonNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldReason, vs...))
}

// ReasonGT applies the GT predicate on the "reason" field.
func ReasonGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldReason, v))
}

// ReasonGTE applies the GTE predicate on the "reason" field.
func ReasonGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldReason, v))
}

// ReasonLT applies the LT predicate on the "reason" field.
func ReasonLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldReason, v))
}

// ReasonLTE applies the LTE predicate on the "reason" field.
func ReasonLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldReason, v))
}

// ReasonContains applies the Contains predicate on the "reason" field.
func ReasonContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldReason, v))
}

// ReasonHasPrefix applies the HasPrefix predicate on the "reason" field.
func ReasonHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldReason, v))
}

// ReasonHasSuffix applies the HasSuffix predicate on the "reason" field.
func ReasonHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldReason, v))
}

// ReasonEqualFold applies the EqualFold predicate on the "reason" field.
func ReasonEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldReason, v))
}

// ReasonContainsFold applies the ContainsFold predicate on the "reason" field.
func ReasonContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldReason, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeContains applies the Contains predicate on the "create_time" field.
func CreateTimeContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldCreateTime, v))
}

// CreateTimeHasPrefix applies the HasPrefix predicate on the "create_time" field.
func CreateTimeHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldCreateTime, v))
}

// CreateTimeHasSuffix applies the HasSuffix predicate on the "create_time" field.
func CreateTimeHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldCreateTime, v))
}

// CreateTimeEqualFold applies the EqualFold predicate on the "create_time" field.
func CreateTimeEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldCreateTime, v))
}

// CreateTimeContainsFold applies the ContainsFold predicate on the "create_time" field.
func CreateTimeContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldCreateTime, v))
}

// FinishTimeEQ applies the EQ predicate on the "finish_time" field.
func FinishTimeEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldFinishTime, v))
}

// FinishTimeNEQ applies the NEQ predicate on the "finish_time" field.
func FinishTimeNEQ(v string) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldFinishTime, v))
}

// FinishTimeIn applies the In predicate on the "finish_time" field.
func FinishTimeIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldFinishTime, vs...))
}

// FinishTimeNotIn applies the NotIn predicate on the "finish_time" field.
func FinishTimeNotIn(vs ...string) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldFinishTime, vs...))
}

// FinishTimeGT applies the GT predicate on the "finish_time" field.
func FinishTimeGT(v string) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldFinishTime, v))
}

// FinishTimeGTE applies the GTE predicate on the "finish_time" field.
func FinishTimeGTE(v string) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldFinishTime, v))
}

// FinishTimeLT applies the LT predicate on the "finish_time" field.
func FinishTimeLT(v string) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldFinishTime, v))
}

// FinishTimeLTE applies the LTE predicate on the "finish_time" field.
func FinishTimeLTE(v string) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldFinishTime, v))
}

// FinishTimeContains applies the Contains predicate on the "finish_time" field.
func FinishTimeContains(v string) predicate.Task {
	return predicate.Task(sql.FieldContains(FieldFinishTime, v))
}

// FinishTimeHasPrefix applies the HasPrefix predicate on the "finish_time" field.
func FinishTimeHasPrefix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasPrefix(FieldFinishTime, v))
}

// FinishTimeHasSuffix applies the HasSuffix predicate on the "finish_time" field.
func FinishTimeHasSuffix(v string) predicate.Task {
	return predicate.Task(sql.FieldHasSuffix(FieldFinishTime, v))
}

// FinishTimeEqualFold applies the EqualFold predicate on the "finish_time" field.
func FinishTimeEqualFold(v string) predicate.Task {
	return predicate.Task(sql.FieldEqualFold(FieldFinishTime, v))
}

// FinishTimeContainsFold applies the ContainsFold predicate on the "finish_time" field.
func FinishTimeContainsFold(v string) predicate.Task {
	return predicate.Task(sql.FieldContainsFold(FieldFinishTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(sql.NotPredicates(p))
}