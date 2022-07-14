package dslcontent

import "github.com/volcengine/datafinder-sdk-openapi-go/consts"

type FilterBuilder struct {
	profileFilter Filter
}

func NewFilterBuilder(profileFilter Filter) *FilterBuilder {
	return &FilterBuilder{profileFilter}
}

func (filterBuilder *FilterBuilder) ShowName(showName string) *FilterBuilder {
	filterBuilder.profileFilter.SetShowName(showName)
	return filterBuilder
}

func (filterBuilder *FilterBuilder) ShowLabel(showLabel string) *FilterBuilder {
	filterBuilder.profileFilter.SetShowLabel(showLabel)
	return filterBuilder
}

func (filterBuilder *FilterBuilder) Show(label, name string) *FilterBuilder {
	filterBuilder.ShowName(name)
	filterBuilder.ShowLabel(label)
	return filterBuilder
}

func (filterBuilder *FilterBuilder) Logic(logic string) *FilterBuilder {
	filterBuilder.profileFilter.UpdateExpression(consts.LOGIC, logic)
	return filterBuilder
}

func (filterBuilder *FilterBuilder) Condition(condition Condition) *FilterBuilder {
	c := []Condition{condition}
	filterBuilder.profileFilter.UpdateExpressionCondition(c)
	return filterBuilder
}
func (filterBuilder *FilterBuilder) Conditions(conditions []Condition) *FilterBuilder {
	filterBuilder.profileFilter.UpdateExpressionCondition(conditions)
	return filterBuilder
}

func (filterBuilder *FilterBuilder) StringExpr(name, operation, propertyType string, values interface{}) *FilterBuilder {
	filterBuilder.Condition(NewCondition(consts.STRING, name, operation, propertyType, values))
	return filterBuilder
}

func (filterBuilder *FilterBuilder) IntExpr(name, operation, propertyType string, values interface{}) *FilterBuilder {
	filterBuilder.Condition(NewCondition(consts.INt, name, operation, propertyType, values))
	return filterBuilder
}

func (filterBuilder *FilterBuilder) StringExprProfile(name, operation string, values interface{}) *FilterBuilder {
	filterBuilder.Condition(NewCondition(consts.STRING, name, operation, consts.PROFILE, values))
	return filterBuilder
}

func (filterBuilder *FilterBuilder) IntExprProfile(name, operation string, values interface{}) *FilterBuilder {
	filterBuilder.Condition(NewCondition(consts.INt, name, operation, consts.PROFILE, values))
	return filterBuilder
}

func (filterBuilder *FilterBuilder) BuildFilter() Filter {
	return filterBuilder.profileFilter
}
