package dslcontent

import (
	"github.com/volcengine/datafinder-sdk-openapi-go/consts"
	"reflect"
)

type DSL struct {
	Version        int                      `json:"version,omitempty"`
	UseAppCloudId  bool                     `json:"use_app_cloud_id,omitempty"`
	AppIds         []int                    `json:"app_ids"`
	Periods        []map[string]interface{} `json:"periods"`
	Content        *Content                 `json:"content,omitempty"`
	Contents       []Content                `json:"contents,omitempty"`
	Option         map[string]interface{}   `json:"option,omitempty"`
	ContentBuilder ContentBuilder           `json:"-"`
}

func NewDSL() DSL {
	option := make(map[string]interface{})
	content := NewContent()
	return DSL{Version: 3, UseAppCloudId: true, Content: &content,
		Option: option, ContentBuilder: BuildContentBuilder()}
}

func (dsl *DSL) AddAppId(appId int) {
	dsl.AppIds = append(dsl.AppIds, appId)
}

func (dsl *DSL) AddAppIds(appIds []int) {
	for _, id := range appIds {
		dsl.AppIds = append(dsl.AppIds, id)
	}
}
func (dsl *DSL) AddPeriods(period map[string]interface{}) {
	dsl.Periods = append(dsl.Periods, period)
}

func (dsl *DSL) AddContents(content Content) {
	dsl.Contents = append(dsl.Contents, content)
}

func (dsl *DSL) SetContentBuilder(contentBuilder ContentBuilder) {
	dsl.ContentBuilder = contentBuilder
}

func (dsl *DSL) IsUseAppCloudId() bool {
	return dsl.UseAppCloudId
}

func (dsl *DSL) SetVersion(version int) {
	dsl.Version = version
}

func (dsl *DSL) SetUseAppCloudId(useAppCloudId bool) {
	dsl.UseAppCloudId = useAppCloudId
}

func (dsl *DSL) SetAppIds(appIds []int) {
	dsl.AppIds = appIds
}

func (dsl *DSL) SetPeriods(periods []map[string]interface{}) {
	dsl.Periods = periods
}

func (dsl *DSL) SetContent(content Content) {
	dsl.Content = &content
}

func (dsl *DSL) SetContents(contents []Content) {
	dsl.Contents = contents
}
func (dsl *DSL) SetOption(option map[string]interface{}) {
	dsl.Option = option
}

func GetDSLBuilder(queryType string) DSLBuilder {
	dsl := NewDSL()
	dslBuilder := DSLBuilder{dsl, queryType}
	dslBuilder.QueryType(queryType)
	return dslBuilder
}

func EventBuilder() DSLBuilder {
	return GetDSLBuilder(consts.EVENT)
}

func FunnelBuilder() DSLBuilder {
	return GetDSLBuilder(consts.FUNNEL)
}

func LifeCycleBuilder() DSLBuilder {
	return GetDSLBuilder(consts.LIFECYCLE)
}

func PathFindBuilder() DSLBuilder {
	return GetDSLBuilder(consts.PATHFIND)
}

func RetentionBuilder() DSLBuilder {
	return GetDSLBuilder(consts.RETENTION)
}

func WebBuilder() DSLBuilder {
	return GetDSLBuilder(consts.WEBSESSION)
}

func ConfidenceBuilder() DSLBuilder {
	return GetDSLBuilder(consts.CONFIDENCE)
}

func TopKBuilder() DSLBuilder {
	return GetDSLBuilder(consts.EVENTOPK)
}

func AdvertiseBuilder() DSLBuilder {
	return GetDSLBuilder(consts.ADVERTISE)
}

func (dsl *DSL) MoveContentToContends() {
	if dsl.Contents == nil && !reflect.DeepEqual(dsl.Content, Content{}) {
		dsl.Contents = []Content{*dsl.Content}
	}
	dsl.Content = nil
}

func MergeDSLs(params map[string]interface{}, dsl []DSL) DSL {
	if len(dsl) == 0 {
		return DSL{}
	}
	d := dsl[0]
	d.MoveContentToContends()
	for i := 1; i < len(dsl); i++ {
		d.AddContents(*dsl[i].Content)
	}
	if len(params) > 0 {
		d.SetOption(params)
	}
	return d
}

func BlendDSLs(base int, dsl []DSL) DSL {
	m := make(map[string]interface{})
	m[consts.BLEND] = map[string]interface{}{consts.STATUS: true, consts.BASE: base}
	return MergeDSLs(m, dsl)
}
