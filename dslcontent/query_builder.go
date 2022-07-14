package dslcontent

import "github.com/volcengine/datafinder-sdk-openapi-go/consts"

type QueryBuilder struct {
	Query Query
}

func (queryBuilder *QueryBuilder) GroupV2(propertyName, propertyType string) *QueryBuilder {
	groupV2 := NewGroupV2(propertyName, propertyType)
	queryBuilder.Query.AddGroupsV2(groupV2)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) Sample(samplePercent int) *QueryBuilder {
	queryBuilder.Query.SetSamplePercent(samplePercent)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) ShowName(name string) *QueryBuilder {
	queryBuilder.Query.SetShowName(name)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) ShowLabel(label string) *QueryBuilder {
	queryBuilder.Query.SetShowLabel(label)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) Event(eventType, eventName, eventIndicator string, eventId int) *QueryBuilder {
	queryBuilder.Query.SetEventID(eventId)
	queryBuilder.Query.SetEventName(eventName)
	queryBuilder.Query.SetEventType(eventType)
	queryBuilder.Query.SetEventIndicator(eventIndicator)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) EventWithoutEventId(eventType, eventName, eventIndicator string) *QueryBuilder {
	queryBuilder.Query.SetEventName(eventName)
	queryBuilder.Query.SetEventType(eventType)
	queryBuilder.Query.SetEventIndicator(eventIndicator)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) EventWithTypeAndName(eventType, eventName string) *QueryBuilder {
	queryBuilder.Query.SetEventName(eventName)
	queryBuilder.Query.SetEventType(eventType)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) MeasureInfo(measureType, measureName string, measureValue int) *QueryBuilder {
	queryBuilder.Query.SetMeasureInfo(map[string]interface{}{
		consts.MEASURE_TYPE:  measureType,
		consts.PROPERTY_NAME: measureName,
		consts.MEASURE_VALUE: measureValue,
	})
	return queryBuilder
}

func (queryBuilder *QueryBuilder) AndFilter(fb *FilterBuilder) *QueryBuilder {
	queryBuilder.Query.AddFilter(fb.Logic(consts.LOGIC_AND).BuildFilter())
	return queryBuilder
}

func (queryBuilder *QueryBuilder) OrFilter(fb *FilterBuilder) *QueryBuilder {
	queryBuilder.Query.AddFilter(fb.Logic(consts.LOGIC_OR).BuildFilter())
	return queryBuilder
}

func (queryBuilder *QueryBuilder) Group(group string) *QueryBuilder {
	queryBuilder.Query.AddGroup(group)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) Groups(groups []string) *QueryBuilder {
	queryBuilder.Query.AddGroups(groups)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) GroupsV2(groupv2 GroupV2) *QueryBuilder {
	queryBuilder.Query.AddGroupsV2(groupv2)
	return queryBuilder
}

func (queryBuilder *QueryBuilder) BuildQuery() Query {
	return queryBuilder.Query
}
