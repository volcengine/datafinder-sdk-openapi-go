package test

import (
	"fmt"
	"github.com/volcengine/datafinder-sdk-openapi-go/dslcontent"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

//获取所有DSL
func TestGetDsl(t *testing.T) {
	fmt.Println("[")
	appid := 0
	dsl := GetEventDSL(appid)
	data, _ := jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetFunnelDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetLifeCycleDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetPathFinderDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetRetentionDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetWebDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetTopKDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data), ",")

	dsl = GetTracerTableDSL(appid)
	data, _ = jsoniter.Marshal(dsl)
	fmt.Println(string(data))

	fmt.Println("]")
}

func GetEventDSL(appid int) dslcontent.DSL {
	builder := dslcontent.EventBuilder()
	return builder.
		AppID(appid). //AppIDs()支持传入多个AppID
		RangePeriod("day", 1562688000, 1563206400, false).
		RangePeriod("hour", 1562688000, 1563206400, false).
		//ProfileGroup("app_channel").
		ProfileGroupV2("os_name", "common_param").
		ProfileGroupV2("custom_app_version", "user_profile").
		SkipCache(false).
		Tags(map[string]interface{}{
			"contains_today": 0,
			"show_yesterday": 0,
			"series_type":    "line",
			"show_map":       map[string]string{},
		}).
		AndProfileFilter(
			dslcontent.IntExprProfile("user_is_new", "=", []int{0}).
				//IntExprProfile() 默认的operationType="profile",value支持int string slice，或者基本数组类型值
				Show("老用户", "1")).
		AndProfileFilter(
			dslcontent.StringExprProfile("language", "=", []string{"zj_CN", "zh_cn"}).
				StringExprProfile("age", "!=", []int{20}).
				Show("zh_CN, zh_cn; not(20)", "2")).
		Query( //查询指标名称, 每个Query表示的是并列关系，Query()方法里面的表示顺序关系  //dslcontent.Show()返回*QueryBuilder
			dslcontent.Show("A", "A"). //dslcontent.Show()返回*QueryBuilder
							Group("app_name").
							GroupV2("os_name", "common_param").
							GroupV2("session_duration", "event_param").
							EventWithoutEventId("origin", "predefine_pageview", "pv"). //Event()可以输入EventId
							MeasureInfo("pct", "event_index", 100).
							AndFilter(
					dslcontent.StringExprProfile("os_name", "=", []string{"windows"}).
						StringExprProfile("network_type", "!=", []string{"wifi"}).
						Show("referer", "referrer_label"))).
		Query(
			dslcontent.Show("B", "B").
				Group("app_name"). //ProfileGroups()支持传入string列表
				EventWithoutEventId("origin", "page_open", "pv").
				AndFilter(
					dslcontent.EmptyExpr().
						Show("app_id", "app_id_label"))).
		BuildDSL()
}

func GetFunnelDSL(appid int) dslcontent.DSL {
	builder := dslcontent.FunnelBuilder()
	return builder.AppID(appid).RangePeriod("day", 1560268800, 1562774400, false).
		ProfileGroup("os_name").
		ProfileGroupV2("custom_app_version", "user_profile").
		Page(10, 2).
		Window("day", 10).
		SkipCache(false).
		AndProfileFilter(dslcontent.IntExprProfile("user_is_new", "=", []int{0}).
			StringExprProfile("network_type", "!=", []string{"4g", "3g"}).
			Show("1", "老用户; not(4g, 3g)")).
		QueryList([]*dslcontent.QueryBuilder{
			dslcontent.Show("1", "查询1").
				Sample(100).
				EventWithoutEventId("origin", "play_time", "pv").
				AndFilter(dslcontent.StringExprProfile("os_name", "=", []string{"windows"}).
					Show("referer_label", "referrer")),
			dslcontent.Show("2", "查询2").
				Sample(100).
				EventWithoutEventId("origin", "app_launch", "pv")}).
		BuildDSL()
}

func GetLifeCycleDSL(appid int) dslcontent.DSL {
	builder := dslcontent.LifeCycleBuilder()
	return builder.AppID(appid).
		RangePeriod("day", 1561910400, 1562428800, false).
		Page(10, 2).
		Window("day", 1).SkipCache(false).
		Tags(map[string]interface{}{
			"series_type":           "line",
			"contains_today":        0,
			"metrics_type":          "number",
			"disabled_in_dashboard": true,
		}).
		AndProfileFilter(
			dslcontent.StringExprProfile("custom_mp_platform", "=", []int{2}).
				StringExprProfile("app_channel", "in", []string{"alibaba", "baidu"}).
				Show("1", "全体用户")).
		Query(dslcontent.Show("active_user", "active_user").
			Sample(100).EventWithoutEventId("origin", "app_launch", "pv")).
		BuildDSL()
}

func GetPathFinderDSL(appid int) dslcontent.DSL {
	builder := dslcontent.PathFindBuilder()
	return builder.AppID(appid).
		RangePeriod("day", 1563120000, 1563638400, false).
		Page(10, 2).
		Window("minute", 10).
		SkipCache(false).
		IsStack(false).
		AndProfileFilter(
							dslcontent.StringExprProfile("os_name", "in", []string{"android", "ios"}).
								StringExprProfile("network_type", "in", []string{"wifi", "4g"}).
								Show("1", "android, ios; wifi, 4g")).
		QueryList([]*dslcontent.QueryBuilder{ //QueryList(qbs []*QueryBuilder)接收指针,定义QueryBuilder指针数组
			dslcontent.Show("1", "查询1").
				Sample(100).
				EventWithTypeAndName("origin", "app_launch").
				AndFilter(dslcontent.EmptyExpr().Show("1", "全体用户")),
			dslcontent.Show("2", "查询2").
				Sample(100).
				EventWithTypeAndName("origin", "register").
				AndFilter(dslcontent.EmptyExpr().Show("1", "全体用户")),
			dslcontent.Show("3", "查询3").
				Sample(100).
				EventWithTypeAndName("origin", "register").
				AndFilter(dslcontent.EmptyExpr().Show("1", "全体用户")),
		}).BuildDSL()
}

func GetRetentionDSL(appid int) dslcontent.DSL {
	builder := dslcontent.RetentionBuilder()
	return builder.AppID(appid).
		RangePeriod("day", 1561910400, 1563033600, false).
		Page(10, 2).
		ProfileGroup("network_type").
		Window("day", 30).
		SkipCache(false).
		IsStack(false).
		Tags(map[string]interface{}{
			"retention_from": "custom",
			"series_type":    "table",
		}).
		AndProfileFilter(dslcontent.IntExprProfile("user_is_new", "=", []int{0}).
			Show("1", "老用户")).
		QueryList(
			[]*dslcontent.QueryBuilder{
				dslcontent.Show("first", "起始事件").
					EventWithoutEventId("origin", "page_open", "pv").
					AndFilter(
						dslcontent.StringExprProfile("os_name", "=", []string{"windows", "mac", "ios"}).
							StringExprProfile("network_type", "!=", []string{"4g"}).
							Show("os_name_label", "os_name,network_type")),
				dslcontent.Show("return", "回访事件").
					EventWithTypeAndName("origin", "any_event").
					AndFilter(dslcontent.StringExprProfile("os_name", "=", []string{"windows", "mac"}).
						StringExprProfile("os_name", "=", []string{"Chrome", "Internet Explore"}).
						Show("1", "全体用户"))}).
		BuildDSL()

}

func GetWebDSL(appid int) dslcontent.DSL {
	builder := dslcontent.WebBuilder()
	return builder.AppID(appid).
		RangePeriod("day", 1562774400, 1563292800, false).
		Page(10, 2).
		ProfileGroup("browser").
		Web("first", 1200).
		SkipCache(false).
		IsStack(false).
		Tags(map[string]interface{}{
			"contains_today": 0,
			"series_type":    "line",
		}).AndProfileFilter(
		dslcontent.StringExprProfile("os_name", "=", []string{"windows", "android"}).
			Show("1", "操作系统")).
		QueryList([]*dslcontent.QueryBuilder{
			dslcontent.Show("session_count", "会话数").
				Sample(100).
				EventWithoutEventId("origin", "predefine_pageview", "session_count").
				AndFilter(dslcontent.EmptyExpr().Show("1", "source")),
			dslcontent.Show("average_session_duration", "平均会话时长").
				EventWithoutEventId("origin", "predefine_pageview", "average_session_duration").
				AndFilter(dslcontent.EmptyExpr().Show("1", "source")),
			dslcontent.Show("bounce_rate", "跳出率").
				EventWithoutEventId("origin", "predefine_pageview", "bounce_rate").
				AndFilter(dslcontent.EmptyExpr().Show("1", "source")),
			dslcontent.Show("average_session_depth", "平均会话深度").
				EventWithoutEventId("origin", "predefine_pageview", "average_session_depth").
				AndFilter(dslcontent.EmptyExpr().Show("1", "source")),
		}).BuildDSL()
}

func GetTopKDSL(appid int) dslcontent.DSL {
	builder := dslcontent.TopKBuilder()
	return builder.AppID(appid).
		RangePeriod("day", 1563379200, 1563897600, false).
		OrderAsc("app_version").
		Page(10, 2).
		SkipCache(true).
		Tags(map[string]interface{}{
			"contains_today": 0,
			"show_yesterday": 0,
			"series_type":    "line",
			"show_map":       map[string]string{},
		}).
		AndProfileFilter(dslcontent.IntExprProfile("ab_version", "=", []int{1}).
			IntExprProfile("user_is_new", "=", []int{0}).
			Show("B", "新用户")).
		Query(dslcontent.Show("A", "查询A").
			Sample(100).
			EventWithoutEventId("origin", "predefine_pageview", "pv").
			MeasureInfo("pct", "event_index", 100).
			AndFilter(dslcontent.StringExpr("referrer", "=", "event_param", []string{"http://www.baidu.com", "http://www.bytedance.com"}).
				Show("referer_label", "referer"))).
		BuildDSL()
}

func GetTracerTableDSL(appid int) dslcontent.DSL {
	builder := dslcontent.AdvertiseBuilder()
	dsl_adv := builder.Advertise(map[string]interface{}{
		"timeout":       1000,
		"alias_convert": false,
		"blend_params": map[string]string{
			"group_by": "date",
		},
	}).
		Product("bytetracer").
		AppID(appid).
		LastPeriod("day", 7, "day", false).
		TodayPeriod("day", true).
		Limit(1000).
		Offset(0).
		AndProfileFilter(
			dslcontent.EmptyExpr().Show("1", "channel_1, traceing_1, group_id_1")).
		Query(dslcontent.Show("impression_count", "impression_count").
			EventWithoutEventId("customed", "impression", "impression_count")).
		Query(dslcontent.Show("click_count", "click_count").
			EventWithoutEventId("customed", "click", "click_count")).
		Query(dslcontent.Show("promotion_activation_count", "promotion_activation_count").
			EventWithoutEventId("customed", "activation", "promotion_activation_count")).
		BuildDSL()

	builder = dslcontent.RetentionBuilder()
	dsl_rate := builder.
		Product("bytefinder").
		AppID(appid).
		LastPeriod("day", 7, "day", false).
		TodayPeriod("day", true).
		Page(1000, 0).
		AndProfileFilter(
			dslcontent.EmptyExpr().Show("1", "channel_1, traceing_1, group_id_1")).
		QueryList([]*dslcontent.QueryBuilder{
			dslcontent.Show("1", "查询1").
				Sample(100).
				EventWithTypeAndName("origin", "app_launch").
				AndFilter(
					dslcontent.IntExprProfile("user_is_new", "=", []int{1}).
						Show("new_user", "new_user")),
			dslcontent.Show("2", "查询2").
				Sample(100).
				EventWithTypeAndName("origin", "app_launch"),
		}).
		BuildDSL()

	return dslcontent.BlendDSLs(0, []dslcontent.DSL{dsl_adv, dsl_rate})
}
