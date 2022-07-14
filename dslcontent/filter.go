package dslcontent

import "github.com/volcengine/datafinder-sdk-openapi-go/consts"

type Filter struct {
	ShowName   string                 `json:"show_name,omitempty"`
	ShowLabel  string                 `json:"show_label,omitempty"`
	Expression map[string]interface{} `json:"expression"`
}

func NewFilter() Filter {
	expression := map[string]interface{}{consts.CONDITIONS: []Condition{}}
	return Filter{Expression: expression}
}

func (filter *Filter) UpdateExpression(key string, value interface{}) {
	filter.Expression[key] = value
}

func (filter *Filter) UpdateExpressionCondition(conditions []Condition) {
	key := consts.CONDITIONS
	if _, ok := filter.Expression[key]; !ok {
		filter.Expression[key] = conditions
		return
	}

	if v, ok := filter.Expression[key].([]Condition); ok {
		for _, condition := range conditions {
			v = append(v, condition)
		}
		filter.Expression[key] = v
	}
}

func (filter *Filter) SetShowName(showName string) {
	filter.ShowName = showName
}

func (filter *Filter) SetShowLabel(showLabel string) {
	filter.ShowLabel = showLabel
}

func (filter *Filter) SetExpression(expression map[string]interface{}) {
	filter.Expression = expression
}

func BuildFilterBuilder() FilterBuilder {
	return *NewFilterBuilder(NewFilter())
}
