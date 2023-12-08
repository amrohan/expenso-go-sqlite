// Code generated by ent, DO NOT EDIT.

package category

import (
	"entgo.io/ent/dialect/sql"
	"github.com/amrohan/expenso-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldIcon, v))
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldUserId, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldName, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldIcon, v))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldIcon, v))
}

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldEQ(FieldUserId, v))
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v string) predicate.Category {
	return predicate.Category(sql.FieldNEQ(FieldUserId, v))
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldIn(FieldUserId, vs...))
}

// UserIdNotIn applies the NotIn predicate on the "userId" field.
func UserIdNotIn(vs ...string) predicate.Category {
	return predicate.Category(sql.FieldNotIn(FieldUserId, vs...))
}

// UserIdGT applies the GT predicate on the "userId" field.
func UserIdGT(v string) predicate.Category {
	return predicate.Category(sql.FieldGT(FieldUserId, v))
}

// UserIdGTE applies the GTE predicate on the "userId" field.
func UserIdGTE(v string) predicate.Category {
	return predicate.Category(sql.FieldGTE(FieldUserId, v))
}

// UserIdLT applies the LT predicate on the "userId" field.
func UserIdLT(v string) predicate.Category {
	return predicate.Category(sql.FieldLT(FieldUserId, v))
}

// UserIdLTE applies the LTE predicate on the "userId" field.
func UserIdLTE(v string) predicate.Category {
	return predicate.Category(sql.FieldLTE(FieldUserId, v))
}

// UserIdContains applies the Contains predicate on the "userId" field.
func UserIdContains(v string) predicate.Category {
	return predicate.Category(sql.FieldContains(FieldUserId, v))
}

// UserIdHasPrefix applies the HasPrefix predicate on the "userId" field.
func UserIdHasPrefix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasPrefix(FieldUserId, v))
}

// UserIdHasSuffix applies the HasSuffix predicate on the "userId" field.
func UserIdHasSuffix(v string) predicate.Category {
	return predicate.Category(sql.FieldHasSuffix(FieldUserId, v))
}

// UserIdEqualFold applies the EqualFold predicate on the "userId" field.
func UserIdEqualFold(v string) predicate.Category {
	return predicate.Category(sql.FieldEqualFold(FieldUserId, v))
}

// UserIdContainsFold applies the ContainsFold predicate on the "userId" field.
func UserIdContainsFold(v string) predicate.Category {
	return predicate.Category(sql.FieldContainsFold(FieldUserId, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Category) predicate.Category {
	return predicate.Category(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Category) predicate.Category {
	return predicate.Category(sql.NotPredicates(p))
}
