package dslcontent

import "github.com/volcengine/datafinder-sdk-openapi-go/consts"

type Content struct {
	QueryType       string                 `json:"query_type,omitempty"`
	ProfileFilters  []Filter               `json:"profile_filters"`
	ProfileGroup    []string               `json:"profile_groups, omitempty"`
	ProfileGroupsV2 []GroupV2              `json:"profile_groups_v2,omitempty"`
	Orders          []map[string]string    `json:"orders"`
	Page            map[string]int         `json:"page"`
	Option          map[string]interface{} `json:"option,omitempty"`
	ShowOption      map[string]interface{} `json:"show_option"`
	Queries         [][]Query              `json:"queries"`
}

func NewContent() Content {
	profileFilters := []Filter{}
	profileGroup := []string{}
	orders := []map[string]string{}
	page := make(map[string]int)
	option := make(map[string]interface{})
	showOption := make(map[string]interface{})
	qureies := [][]Query{}
	return Content{ProfileFilters: profileFilters, ProfileGroup: profileGroup, QueryType: "event",
		Orders: orders, Page: page, Option: option, ShowOption: showOption, Queries: qureies}
}

func (content *Content) AddProfileGroupsV2(groupv2 GroupV2) {
	content.ProfileGroupsV2 = append(content.ProfileGroupsV2, groupv2)
}

func (content *Content) AddProfileFilter(pf Filter) {
	content.ProfileFilters = append(content.ProfileFilters, pf)
}

func (content *Content) AddProfileGroup(pg interface{}) {
	switch pg.(type) {
	case []string:
		data, _ := pg.([]string)
		for _, group := range data {
			content.ProfileGroup = append(content.ProfileGroup, group)
		}
	default:
		data, _ := pg.(string)
		content.ProfileGroup = append(content.ProfileGroup, data)
	}
}

func (content *Content) AddOrderAsc(order string) {
	content.AddOrder(order, consts.ASC)
}

func (content *Content) AddOrder(order, direction string) {
	t := map[string]string{consts.FIELD: order, consts.DIRECTION: direction}
	content.Orders = append(content.Orders, t)
}

func (content *Content) AddOrderMaps(orders []map[string]string) {
	for _, order := range orders {
		content.Orders = append(content.Orders, order)
	}
}

func (content *Content) AddQuery(queries []Query) {
	content.Queries = append(content.Queries, queries)
}

func (content *Content) UpdatePage(key string, value int) {
	content.Page[key] = value
}

func (content *Content) UpdateOption(key string, option interface{}) {
	content.Option[key] = option
}

func (content *Content) UpdateShowOption(key string, showOpt interface{}) {
	content.ShowOption[key] = showOpt
}

func (content *Content) SetQueryType(queryType string) {
	content.QueryType = queryType
}

func (content *Content) SetProfileFilters(profileFilters []Filter) {
	content.ProfileFilters = profileFilters
}

func (content *Content) SetProfileGroup(profileGroup []string) {
	content.ProfileGroup = profileGroup
}

func (content *Content) SetOrders(orders []map[string]string) {
	content.Orders = orders
}

func (content *Content) SetPage(page map[string]int) {
	content.Page = page
}

func (content *Content) SetOption(option map[string]interface{}) {
	content.Option = option
}

func (content *Content) SetShowOption(showOption map[string]interface{}) {
	content.ShowOption = showOption
}

func (content *Content) SetQueries(queries [][]Query) {
	content.Queries = queries
}

func BuildContentBuilder() ContentBuilder {
	return ContentBuilder{NewContent()}
}
