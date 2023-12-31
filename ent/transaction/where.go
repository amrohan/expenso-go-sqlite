// Code generated by ent, DO NOT EDIT.

package transaction

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/amrohan/expenso-go/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldID, id))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTitle, v))
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDescription, v))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDate, v))
}

// CategoryId applies equality check predicate on the "categoryId" field. It's identical to CategoryIdEQ.
func CategoryId(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCategoryId, v))
}

// AccountId applies equality check predicate on the "accountId" field. It's identical to AccountIdEQ.
func AccountId(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAccountId, v))
}

// UserId applies equality check predicate on the "userId" field. It's identical to UserIdEQ.
func UserId(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUserId, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldType, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldStatus, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCurrency, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldTitle, v))
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAmount, v))
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldAmount, v))
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldAmount, vs...))
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldAmount, vs...))
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldAmount, v))
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldAmount, v))
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldAmount, v))
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v float64) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldAmount, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldDescription, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldDate, v))
}

// CategoryIdEQ applies the EQ predicate on the "categoryId" field.
func CategoryIdEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCategoryId, v))
}

// CategoryIdNEQ applies the NEQ predicate on the "categoryId" field.
func CategoryIdNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCategoryId, v))
}

// CategoryIdIn applies the In predicate on the "categoryId" field.
func CategoryIdIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCategoryId, vs...))
}

// CategoryIdNotIn applies the NotIn predicate on the "categoryId" field.
func CategoryIdNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCategoryId, vs...))
}

// CategoryIdGT applies the GT predicate on the "categoryId" field.
func CategoryIdGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldCategoryId, v))
}

// CategoryIdGTE applies the GTE predicate on the "categoryId" field.
func CategoryIdGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldCategoryId, v))
}

// CategoryIdLT applies the LT predicate on the "categoryId" field.
func CategoryIdLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldCategoryId, v))
}

// CategoryIdLTE applies the LTE predicate on the "categoryId" field.
func CategoryIdLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldCategoryId, v))
}

// AccountIdEQ applies the EQ predicate on the "accountId" field.
func AccountIdEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldAccountId, v))
}

// AccountIdNEQ applies the NEQ predicate on the "accountId" field.
func AccountIdNEQ(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldAccountId, v))
}

// AccountIdIn applies the In predicate on the "accountId" field.
func AccountIdIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldAccountId, vs...))
}

// AccountIdNotIn applies the NotIn predicate on the "accountId" field.
func AccountIdNotIn(vs ...int) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldAccountId, vs...))
}

// AccountIdGT applies the GT predicate on the "accountId" field.
func AccountIdGT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldAccountId, v))
}

// AccountIdGTE applies the GTE predicate on the "accountId" field.
func AccountIdGTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldAccountId, v))
}

// AccountIdLT applies the LT predicate on the "accountId" field.
func AccountIdLT(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldAccountId, v))
}

// AccountIdLTE applies the LTE predicate on the "accountId" field.
func AccountIdLTE(v int) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldAccountId, v))
}

// UserIdEQ applies the EQ predicate on the "userId" field.
func UserIdEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldUserId, v))
}

// UserIdNEQ applies the NEQ predicate on the "userId" field.
func UserIdNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldUserId, v))
}

// UserIdIn applies the In predicate on the "userId" field.
func UserIdIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldUserId, vs...))
}

// UserIdNotIn applies the NotIn predicate on the "userId" field.
func UserIdNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldUserId, vs...))
}

// UserIdGT applies the GT predicate on the "userId" field.
func UserIdGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldUserId, v))
}

// UserIdGTE applies the GTE predicate on the "userId" field.
func UserIdGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldUserId, v))
}

// UserIdLT applies the LT predicate on the "userId" field.
func UserIdLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldUserId, v))
}

// UserIdLTE applies the LTE predicate on the "userId" field.
func UserIdLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldUserId, v))
}

// UserIdContains applies the Contains predicate on the "userId" field.
func UserIdContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldUserId, v))
}

// UserIdHasPrefix applies the HasPrefix predicate on the "userId" field.
func UserIdHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldUserId, v))
}

// UserIdHasSuffix applies the HasSuffix predicate on the "userId" field.
func UserIdHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldUserId, v))
}

// UserIdEqualFold applies the EqualFold predicate on the "userId" field.
func UserIdEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldUserId, v))
}

// UserIdContainsFold applies the ContainsFold predicate on the "userId" field.
func UserIdContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldUserId, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldType, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldStatus, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...string) predicate.Transaction {
	return predicate.Transaction(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyContains applies the Contains predicate on the "currency" field.
func CurrencyContains(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContains(FieldCurrency, v))
}

// CurrencyHasPrefix applies the HasPrefix predicate on the "currency" field.
func CurrencyHasPrefix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasPrefix(FieldCurrency, v))
}

// CurrencyHasSuffix applies the HasSuffix predicate on the "currency" field.
func CurrencyHasSuffix(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldHasSuffix(FieldCurrency, v))
}

// CurrencyIsNil applies the IsNil predicate on the "currency" field.
func CurrencyIsNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldIsNull(FieldCurrency))
}

// CurrencyNotNil applies the NotNil predicate on the "currency" field.
func CurrencyNotNil() predicate.Transaction {
	return predicate.Transaction(sql.FieldNotNull(FieldCurrency))
}

// CurrencyEqualFold applies the EqualFold predicate on the "currency" field.
func CurrencyEqualFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldEqualFold(FieldCurrency, v))
}

// CurrencyContainsFold applies the ContainsFold predicate on the "currency" field.
func CurrencyContainsFold(v string) predicate.Transaction {
	return predicate.Transaction(sql.FieldContainsFold(FieldCurrency, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Transaction) predicate.Transaction {
	return predicate.Transaction(sql.NotPredicates(p))
}
