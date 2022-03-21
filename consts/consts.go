package consts

const (
	ORG = "dataRangers"
	URL = "https://analytics.volcengineapi.com"

	METHOD_NOT_SUPPORT  = "method not support"
	POST                = "POST"
	GET                 = "GET"
	PUT                 = "PUT"
	POST_BODY_NULL      = "post method mush contains body"
	SERVICE_NOT_SUPPORT = "service not support"
	AUTHORIZATION       = "Authorization"
	CONTENT_TYPE        = "Content-Type"
	APPLICATION_JSON    = "application/json;charset=utf-8"

	ANALYSIS_BASE = "analysis_base"
	DATA_FINDER   = "data_finder"
	DATA_TRACER   = "data_tracer"
	DATA_TESTER   = "data_tester"
	DATA_ANALYZER = "data_analyzer"
	DATA_RANGGERS = "data_rangers"

	ANALYSIS_BASE_URL = "/analysisbase"
	DATA_FINDER_URL   = "/datafinder"
	DATA_TRACER_URL   = "/datatracer"
	DATA_TESTER_URL   = "/datatester"
	DATA_ANALYZER_URL = "/dataanalyzer"
	DATA_RANGGERS_URL = "/datarangers"

	DATA_PROFILE     = "dataprofile"
	DATA_PROFILE_URL = "/dataprofile"

	MEASURE_TYPE  = "measure_type"
	MEASURE_VALUE = "measure_value"
	PROPERTY_NAME = "property_name"

	ASC       = "asc"
	FIELD     = "field"
	DIRECTION = "direction"

	LOGIC     = "logic"
	LOGIC_AND = "and"
	LOGIC_OR  = "or"

	LIMIt  = "limit"
	OFFSET = "offset"

	EVENT      = "event"
	FUNNEL     = "funnel"
	LIFECYCLE  = "life_cycle"
	PATHFIND   = "path_find"
	RETENTION  = "retention"
	WEBSESSION = "web_session"
	CONFIDENCE = "confidence"
	EVENTOPK   = "event_topk"
	ADVERTISE  = "advertise"

	BLEND  = "blend"
	STATUS = "status"
	BASE   = "base"

	OPTIMIZED   = "optimized"
	TYPE        = "type"
	RANGE       = "range"
	GRANULARITY = "granularity"
	REALTIME    = "real_time"
	LAST        = "last"
	TODAY       = "today"
	AMOUNT      = "amount"
	UNIT        = "unit"

	SKIP_CACHE           = "skip_cache"
	IS_STACK             = "is_stack"
	LIFECYCLE_QUERY_TYPE = "lifecycle_query_type"
	LIFECYCLE_PERIOD     = "lifecycle_period"
	PERIOD               = "period"
	STICKINESS           = "stickiness"
	RETENTION_TYPE       = "retention_type"
	RETENTION_N_DAYS     = "retention_n_days"
	WeB_SESSION_PARAMS   = "web_session_params"
	SESSION_PARAMS_TYPE  = "session_params_type"
	SESSION_TIMEOUT      = "session_timeout"
	PRODUCT              = "product"
	WINDOW_PERIOD_TYPE   = "window_period_type"
	WINDOW_PERIOD        = "window_period"

	STRING  = "string"
	INt     = "int"
	PROFILE = "profile"

	CONDITIONS = "conditions"
)

var METHOD_ALLOWED = [5]string{"POST", "GET", "DELETE", "PUT", "PATCH"}
