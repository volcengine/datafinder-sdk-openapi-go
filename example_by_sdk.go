package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/volcengine/datafinder-sdk-openapi-go/dslcontent"
)

var (
	ak = "xxx"
	sk = "xxx"
)

func main() {
	client := dslcontent.NewRangersClient(ak, sk)
	appid := "xxx"
	event_id := "xxx"
	TestGetDataExportUrl(&client, appid)
	TestFinderAll(&client, appid)
	TestFinderDashboard(&client, appid)
	TestFinderDashboardReports(&client, appid)
	TestFinderDashboardReport(&client, "xxx", appid)
	TestFinderCohorts(&client, appid)
	TestFinderCohort(&client, appid)
	TestFinderCohortsSample(&client, appid)
	TestEventDSLQuery(&client, appid, event_id, 1603929600, 1603929600)
}

func TestGetDataExportUrl(client *dslcontent.RangersClient, appid string) {
	//如果需要指定url: client:= dslcontent.NewRangersClientWithUrl(ak,sk,url)
	// client := dslcontent.NewRangersClient(ak, sk)
	res, err := client.DataRangers("/openapi/v1/"+appid+"/date/2020-05-03/2020-05-09/downloads", "get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderAll(client *dslcontent.RangersClient, appid string) {
	res, err := client.DataFinder("/openapi/v1/"+appid+"/dashboards/all", "get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderDashboard(client *dslcontent.RangersClient, appid string) {
	res, err := client.DataFinder("/openapi/v1/"+appid+"/dashboards/xxx", "get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderDashboardReports(client *dslcontent.RangersClient, appid string) {
	res, err := client.DataFinder("/openapi/v1/"+appid+"/dashboards/xxx/reports", "get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderDashboardReport(client *dslcontent.RangersClient, id string, appid string) {
	headers := map[string]string{}
	params := map[string]string{"count": "10"}
	res, err := client.DataFinderFull("/openapi/v1/"+appid+"/reports/"+id, "get", headers, params, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderCohorts(client *dslcontent.RangersClient, appid string) {
	res, err := client.DataFinder("/openapi/v1/"+appid+"/cohorts", "get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderCohort(client *dslcontent.RangersClient, appid string) {
	res, err := client.DataFinder("/openapi/v1/"+appid+"/cohorts/4627", "get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestFinderCohortsSample(client *dslcontent.RangersClient, appid string) {
	headers := map[string]string{}
	params := map[string]string{"count": "10"}
	res, err := client.DataFinderFull("/openapi/v1/"+appid+"/cohorts/4453/sample", "get", headers, params, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func TestEventDSLQuery(client *dslcontent.RangersClient, appid string, event_id string, period_start int, period_end int) {
	appid_int, _ := strconv.Atoi(appid)
	dsl := GetEventDSL(appid_int, event_id, period_start, period_end)
	data, _ := json.Marshal(dsl)
	resp, _ := client.DataAnalyzerFull("/openapi/v3/analysis", "post", nil, nil, string(data))
	fmt.Print(resp)
}

func GetEventDSL(appid int, event_id string, start, end int) dslcontent.DSL {
	builder := dslcontent.EventBuilder()
	return builder.
		AppID(appid). //AppIDs()支持传入多个AppID
		RangePeriod("day", start, end, false).
		RangePeriod("hour", start, end, false).
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
