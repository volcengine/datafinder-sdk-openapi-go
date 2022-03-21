package dslcontent

import "gosdk/consts"

type ContentBuilder struct {
	Content Content
}

func (contentBuilder *ContentBuilder) QueryType(queryType string) *ContentBuilder {
	contentBuilder.Content.SetQueryType(queryType)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) ProfileFilter(filter Filter) *ContentBuilder {
	contentBuilder.Content.AddProfileFilter(filter)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) ProfileGroup(pg string) *ContentBuilder {
	contentBuilder.Content.AddProfileGroup(pg)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) ProfileGroups(groups []string) *ContentBuilder {
	contentBuilder.Content.AddProfileGroup(groups)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) ProfileGroupsV2(pg2 GroupV2) *ContentBuilder {
	contentBuilder.Content.AddProfileGroupsV2(pg2)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Order(order, direction string) *ContentBuilder {
	contentBuilder.Content.AddOrder(order, direction)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) OrderAsc(order string) *ContentBuilder {
	contentBuilder.Content.AddOrderAsc(order)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) OrderMap(order map[string]string) *ContentBuilder {
	var t []map[string]string
	t = append(t, order)
	contentBuilder.Content.AddOrderMaps(t)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) OrderMaps(orders []map[string]string) *ContentBuilder {
	contentBuilder.Content.AddOrderMaps(orders)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Page(limit, offset int) *ContentBuilder {
	contentBuilder.Content.UpdatePage(consts.LIMIt, limit)
	contentBuilder.Content.UpdatePage(consts.OFFSET, offset)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Limit(limit int) *ContentBuilder {
	contentBuilder.Content.UpdatePage(consts.LIMIt, limit)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Offset(offset int) *ContentBuilder {
	contentBuilder.Content.UpdatePage(consts.OFFSET, offset)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Option(key string, value interface{}) *ContentBuilder {
	contentBuilder.Content.UpdateOption(key, value)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) ShowOption(key string, value interface{}) *ContentBuilder {
	contentBuilder.Content.UpdateShowOption(key, value)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Query(query Query) *ContentBuilder {
	q := []Query{query}
	contentBuilder.Content.AddQuery(q)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) QueryList(queries []Query) *ContentBuilder {
	contentBuilder.Content.AddQuery(queries)
	return contentBuilder
}

func (contentBuilder *ContentBuilder) Build() Content {
	return contentBuilder.Content
}
