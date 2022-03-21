package dslcontent

import "gosdk/consts"

func NewExpr(valueType, name, operation, operationType string, value interface{}) *FilterBuilder {
	builder := BuildFilterBuilder()
	return builder.Condition(NewCondition(valueType, name, operation, operationType, value))
}

func EmptyExpr() *FilterBuilder {
	FilterBuilder := BuildFilterBuilder()
	return &FilterBuilder
}

func StringExpr(name, operation, operationType string, value interface{}) *FilterBuilder {
	return NewExpr(consts.STRING, name, operation, operationType, value)
}

func StringExprProfile(name, operation string, value interface{}) *FilterBuilder {
	return NewExpr(consts.STRING, name, operation, consts.PROFILE, value)
}

func IntExpr(name, operation, operationType string, value interface{}) *FilterBuilder {
	return NewExpr(consts.INt, name, operation, operationType, value)
}

func IntExprProfile(name, operation string, value interface{}) *FilterBuilder {
	return NewExpr(consts.INt, name, operation, consts.PROFILE, value)
}

func Show(label, name string) *QueryBuilder {
	builder := BuildQueryBuilder()
	return builder.ShowLabel(label).ShowName(name)
}
