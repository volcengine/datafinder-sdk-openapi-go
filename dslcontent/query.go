package dslcontent

type Query struct {
	SamplePercent  int                    `json:"sample_percent,omitempty"`
	ShowName       string                 `json:"show_name,omitempty"`
	ShowLabel      string                 `json:"show_label,omitempty"`
	EventID        int                    `json:"event_id,omitempty"`
	EventType      string                 `json:"event_type,omitempty"`
	EventName      string                 `json:"event_name,omitempty"`
	EventIndicator string                 `json:"event_indicator,omitempty"`
	MeasureInfo    map[string]interface{} `json:"measure_info"`
	Filters        []Filter               `json:"filters"`
	Groups         []string               `json:"groups, omitempty"`
	GroupsV2       []GroupV2              `json:"groups_v2,omitempty"`
}

func NewQuery() Query {
	measureInfo := make(map[string]interface{})
	filters := []Filter{}
	groups := []string{}
	return Query{SamplePercent: 100, MeasureInfo: measureInfo, Filters: filters, Groups: groups}
}

func (query *Query) AddFilter(filter Filter) {
	query.Filters = append(query.Filters, filter)
}

func (query *Query) AddGroup(group string) {
	query.Groups = append(query.Groups, group)
}

func (query *Query) AddGroups(groups []string) {
	for _, group := range groups {
		query.AddGroup(group)
	}
}

func (query *Query) AddGroupsV2(groupV2 GroupV2) {
	query.GroupsV2 = append(query.GroupsV2, groupV2)
}

func (query *Query) SetSamplePercent(samplePercent int) {
	query.SamplePercent = samplePercent
}

func (query *Query) SetShowName(showName string) {
	query.ShowName = showName
}

func (query *Query) SetShowLabel(showLabel string) {
	query.ShowLabel = showLabel
}

func (query *Query) SetEventID(eventID int) {
	query.EventID = eventID
}

func (query *Query) SetEventType(eventType string) {
	query.EventType = eventType
}

func (query *Query) SetEventName(eventName string) {
	query.EventName = eventName
}

func (query *Query) SetEventIndicator(eventIndicator string) {
	query.EventIndicator = eventIndicator
}

func (query *Query) SetMeasureInfo(measureInfo map[string]interface{}) {
	query.MeasureInfo = measureInfo
}
func (query *Query) SetFilters(filters []Filter) {
	query.Filters = filters
}
func (query *Query) SetGroups(groups []string) {
	query.Groups = groups
}

func BuildQueryBuilder() QueryBuilder {
	return QueryBuilder{NewQuery()}
}
