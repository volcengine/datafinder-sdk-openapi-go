package dslcontent

import "gosdk/consts"

type DSLBuilder struct {
	dsl       DSL
	queryType string
}

func NewDSLBuilder(dsl DSL, queryType string) *DSLBuilder {
	return &DSLBuilder{dsl, queryType}
}

func (dslBuilder *DSLBuilder) AppID(appId int) *DSLBuilder {
	dslBuilder.dsl.AddAppId(appId)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) AppIDs(appIds []int) *DSLBuilder {
	dslBuilder.dsl.AddAppIds(appIds)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Optimized(opt bool) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Option(consts.OPTIMIZED, opt)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) QueryType(queryType string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.QueryType(queryType)
	if consts.FUNNEL == queryType {
		dslBuilder.Optimized(true)
	} else {
		dslBuilder.Optimized(false)
	}
	return dslBuilder
}

func (dslBuilder *DSLBuilder) RangePeriod(granularity string, start int, end int, realtime bool) *DSLBuilder {
	period := map[string]interface{}{
		consts.TYPE:        consts.RANGE,
		consts.GRANULARITY: granularity,
		consts.RANGE:       []int{start, end},
	}
	if realtime {
		period[consts.REALTIME] = realtime
	}
	dslBuilder.dsl.AddPeriods(period)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) LastPeriod(granularity string, amount int, unit string, realtime bool) *DSLBuilder {
	period := map[string]interface{}{
		consts.TYPE:        consts.LAST,
		consts.GRANULARITY: granularity,
		consts.LAST:        map[string]interface{}{consts.AMOUNT: amount, consts.UNIT: unit},
	}
	if realtime {
		period[consts.REALTIME] = realtime
	}
	dslBuilder.dsl.AddPeriods(period)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) TodayPeriod(granularity string, realtime bool) *DSLBuilder {
	period := map[string]interface{}{
		consts.TYPE:        consts.TODAY,
		consts.GRANULARITY: granularity,
	}
	if realtime {
		period[consts.REALTIME] = realtime
	}
	dslBuilder.dsl.AddPeriods(period)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) ProfileGroup(pg string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.ProfileGroup(pg)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) ProfileGroups(pgs []string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.ProfileGroups(pgs)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) ProfileGroupV2(propertyName, propertyType string) *DSLBuilder {
	groupV2 := NewGroupV2(propertyName, propertyType)
	dslBuilder.dsl.ContentBuilder.ProfileGroupsV2(groupV2)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) OrderAsc(order string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Order(order, consts.ASC)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Order(order string, direction string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Order(order, direction)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) OrderMap(orderMap map[string]string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.OrderMap(orderMap)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) OrderList(orderList []map[string]string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.OrderMaps(orderList)
	return dslBuilder
}
func (dslBuilder *DSLBuilder) Page(limit int, offset int) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Page(limit, offset)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Limit(limit int) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Limit(limit)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Offset(offset int) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Offset(offset)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) SkipCache(sc bool) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Option(consts.SKIP_CACHE, sc)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) IsStack(stack bool) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Option(consts.IS_STACK, stack)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) LifeCycle(granularity string, interval int, lifeCycleQueryType string) *DSLBuilder {
	builder := dslBuilder.dsl.ContentBuilder
	builder.Option(consts.LIFECYCLE_QUERY_TYPE, lifeCycleQueryType)
	builder.Option(consts.LIFECYCLE_PERIOD, map[string]interface{}{
		consts.GRANULARITY: granularity,
		consts.PERIOD:      interval,
	})
	return dslBuilder
}

func (dslBuilder *DSLBuilder) LifeCycleStickinessType(granularity string, interval int) *DSLBuilder {
	return dslBuilder.LifeCycle(granularity, interval, consts.STICKINESS)
}

func (dslBuilder *DSLBuilder) Retention(granularity string, interval int) *DSLBuilder {
	builder := dslBuilder.dsl.ContentBuilder
	builder.Option(consts.RETENTION_TYPE, granularity)
	builder.Option(consts.RETENTION_N_DAYS, interval)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Web(paramType string, timeout int) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Option(consts.WeB_SESSION_PARAMS, map[string]interface{}{
		consts.SESSION_PARAMS_TYPE: paramType,
		consts.SESSION_TIMEOUT:     timeout,
	})
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Product(product string) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Option(consts.PRODUCT, product)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Advertise(adp map[string]interface{}) *DSLBuilder {
	return dslBuilder.Option(adp)
}

func (dslBuilder *DSLBuilder) Option(options map[string]interface{}) *DSLBuilder {
	builder := dslBuilder.dsl.ContentBuilder
	for op := range options {
		builder.Option(op, options[op])
	}
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Tags(tags map[string]interface{}) *DSLBuilder {
	builder := dslBuilder.dsl.ContentBuilder
	for tag := range tags {
		builder.ShowOption(tag, tags[tag])
	}
	return dslBuilder
}

func (dslBuilder *DSLBuilder) AndProfileFilter(fb *FilterBuilder) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.ProfileFilter(fb.Logic(consts.LOGIC_AND).BuildFilter())
	return dslBuilder
}
func (dslBuilder *DSLBuilder) OrProfileFilter(fb *FilterBuilder) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.ProfileFilter(fb.Logic(consts.LOGIC_OR).BuildFilter())
	return dslBuilder
}

func (dslBuilder *DSLBuilder) QueryList(qbs []*QueryBuilder) *DSLBuilder {
	var queries []Query
	for _, qb := range qbs {
		queries = append(queries, qb.BuildQuery())
	}
	dslBuilder.dsl.ContentBuilder.QueryList(queries)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Query(qb *QueryBuilder) *DSLBuilder {
	dslBuilder.dsl.ContentBuilder.Query(qb.BuildQuery())
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Periods(periods []map[string]interface{}) *DSLBuilder {
	dslBuilder.dsl.SetPeriods(periods)
	return dslBuilder
}

func (dslBuilder *DSLBuilder) Window(granularity string, interval int) *DSLBuilder {
	if consts.LIFECYCLE == dslBuilder.queryType {
		dslBuilder.LifeCycleStickinessType(granularity, interval)
	} else if consts.RETENTION == dslBuilder.queryType {
		dslBuilder.Retention(granularity, interval)
	} else {
		builder := dslBuilder.dsl.ContentBuilder
		builder.Option(consts.WINDOW_PERIOD_TYPE, granularity)
		builder.Option(consts.WINDOW_PERIOD, interval)
	}
	return dslBuilder
}

func (dslBuilder *DSLBuilder) BuildDSL() DSL {
	dslBuilder.dsl.SetContent(dslBuilder.dsl.ContentBuilder.Build())
	return dslBuilder.dsl
}
