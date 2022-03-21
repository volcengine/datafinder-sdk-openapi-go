package dslcontent

import (
	"errors"
	"net/http"
	"strings"

	"gosdk/consts"
	"gosdk/util"
)

var (
	org        = consts.ORG
	url        = consts.URL
	expiration = 1800
	services   = map[string]string{
		consts.DATA_FINDER:   consts.DATA_FINDER_URL,
		consts.DATA_TRACER:   consts.DATA_TRACER_URL,
		consts.DATA_TESTER:   consts.DATA_TESTER_URL,
		consts.DATA_ANALYZER: consts.DATA_ANALYZER_URL,
		consts.DATA_RANGGERS: consts.DATA_RANGGERS_URL,
		consts.DATA_PROFILE:  consts.DATA_PROFILE_URL,
	}
)

type RangersClient struct {
	Ak         string
	Sk         string
	Org        string
	Url        string
	Expiration int
	Services   map[string]string
}

func NewRangersClientFull(ak, sk, url string, expiration int, services map[string]string) RangersClient {
	return RangersClient{ak, sk, org, url, expiration, services}
}

func NewRangersClient(ak, sk string) RangersClient {
	return RangersClient{ak, sk, org, url, expiration, services}
}

func NewRangersClientWithUrl(ak, sk, url string) RangersClient {
	return RangersClient{ak, sk, org, url, expiration, services}
}

func NewRangersClientWithUrlAndExp(ak, sk, url string, expiration int) RangersClient {
	return RangersClient{ak, sk, org, url, expiration, services}
}

func (rangersClient *RangersClient) getServicePath(service string) string {
	return rangersClient.Services[service]
}

func (rangersClient *RangersClient) methodValid(method string) bool {
	for _, m := range consts.METHOD_ALLOWED {
		if method == m {
			return true
		}
	}
	return false
}

func handleParams(request *http.Request, headers map[string]string, params map[string]string) {
	req := request.URL.Query()
	if params != nil {
		for p := range params {
			req.Add(p, params[p])
		}
	}
	request.URL.RawQuery = req.Encode()
	for h := range headers {
		request.Header.Set(h, headers[h])
	}
}

func (rangersClient *RangersClient) request(service string, method string, path string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	method = strings.ToUpper(method)
	if !rangersClient.methodValid(method) {
		return http.Response{}, errors.New(consts.METHOD_NOT_SUPPORT)
	}
	if method == consts.POST && body == "" {
		return http.Response{}, errors.New(consts.POST_BODY_NULL)
	}

	servicePath := rangersClient.getServicePath(service)

	if servicePath == "" {
		return http.Response{}, errors.New(consts.SERVICE_NOT_SUPPORT)
	}
	serviceUrl := servicePath + path
	authorization := util.Sign(rangersClient.Ak, rangersClient.Sk, rangersClient.Expiration, method, serviceUrl, params, body)
	if headers == nil {
		headers = map[string]string{}
	}
	headers[consts.AUTHORIZATION] = authorization
	if consts.POST == method {
		headers[consts.CONTENT_TYPE] = consts.APPLICATION_JSON
	}
	url := strings.TrimSpace(rangersClient.Url + serviceUrl)
	client := &http.Client{}
	if method == consts.PUT {
		request, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))
		if err != nil {
			return http.Response{}, err
		}
		handleParams(request, headers, params)
		response, err := client.Do(request)
		if err != nil {
			return http.Response{}, err
		}
		return *response, nil

	} else if method == consts.POST {
		request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
		if err != nil {
			return http.Response{}, err
		}
		handleParams(request, headers, params)
		response, err := client.Do(request)
		if err != nil {
			return http.Response{}, err
		}
		return *response, nil
	} else {
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return http.Response{}, err
		}
		handleParams(request, headers, params)
		response, err := client.Do(request)
		if err != nil {
			return http.Response{}, err
		}
		return *response, nil
	}

}

func formatParams(params map[string]string) string {
	paramStr := "?"
	for k := range params {
		paramStr = paramStr + k + "=" + params[k] + "&"
	}
	return paramStr[0 : len(paramStr)-1]
}

func (rangersClient *RangersClient) SetOrg(org string) {
	rangersClient.Org = org
}

func (rangersClient *RangersClient) SetUrl(url string) {
	rangersClient.Url = url
}

func (rangersClient *RangersClient) SetExpiration(expiration int) {
	rangersClient.Expiration = expiration
}

func (rangersClient *RangersClient) GetAk() string {
	return rangersClient.Ak
}

func (rangersClient *RangersClient) GetSk() string {
	return rangersClient.Sk
}

func (rangersClient *RangersClient) GetOrg() string {
	return rangersClient.Org
}

func (rangersClient *RangersClient) GetUrl() string {
	return rangersClient.Url
}

func (rangersClient *RangersClient) GetExpiration() int {
	return rangersClient.Expiration
}

func (rangersClient *RangersClient) DataRangersFull(path string, method string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	return rangersClient.request(consts.DATA_RANGGERS, method, path, headers, params, body)
}

func (rangersClient *RangersClient) DataRangers(path string, method string) (http.Response, error) {
	return rangersClient.DataRangersFull(path, method, nil, nil, "")
}

func (rangersClient *RangersClient) DataRangersGet(path string) (http.Response, error) {
	return rangersClient.DataRangersFull(path, consts.GET, nil, nil, "")
}

func (rangersClient *RangersClient) DataFinderFull(path string, method string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	return rangersClient.request(consts.DATA_FINDER, method, path, headers, params, body)
}

func (rangersClient *RangersClient) DataFinder(path string, method string) (http.Response, error) {
	return rangersClient.DataFinderFull(path, method, nil, nil, "")
}

func (rangersClient *RangersClient) DataFinderGet(path string) (http.Response, error) {
	return rangersClient.DataFinderFull(path, consts.GET, nil, nil, "")
}

func (rangersClient *RangersClient) DataTracerFull(path string, method string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	return rangersClient.request(consts.DATA_TRACER, method, path, headers, params, body)
}

func (rangersClient *RangersClient) DataTracer(path string, method string) (http.Response, error) {
	return rangersClient.DataTracerFull(path, method, nil, nil, "")
}

func (rangersClient *RangersClient) DataTracerGet(path string) (http.Response, error) {
	return rangersClient.DataTracerFull(path, consts.GET, nil, nil, "")
}

func (rangersClient *RangersClient) DataTesterFull(path string, method string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	return rangersClient.request(consts.DATA_TESTER, method, path, headers, params, body)
}

func (rangersClient *RangersClient) DataTester(path string, method string) (http.Response, error) {
	return rangersClient.DataTesterFull(path, method, nil, nil, "")
}

func (rangersClient *RangersClient) DataTesterGet(path string) (http.Response, error) {
	return rangersClient.DataTesterFull(path, consts.GET, nil, nil, "")
}

func (rangersClient *RangersClient) DataAnalyzerFull(path string, method string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	return rangersClient.request(consts.DATA_ANALYZER, method, path, headers, params, body)
}

func (rangersClient *RangersClient) DataAnalyzer(path string, method string) (http.Response, error) {
	return rangersClient.DataAnalyzerFull(path, method, nil, nil, "")
}

func (rangersClient *RangersClient) DataAnalyzerGet(path string) (http.Response, error) {
	return rangersClient.DataAnalyzerFull(path, consts.GET, nil, nil, "")
}

func (rangersClient *RangersClient) DataProfileFull(path string, method string, headers map[string]string, params map[string]string, body string) (http.Response, error) {
	return rangersClient.request(consts.DATA_PROFILE, method, path, headers, params, body)
}

func (rangersClient *RangersClient) DataProfile(path string, method string) (http.Response, error) {
	return rangersClient.DataProfileFull(path, method, nil, nil, "")
}

func (rangersClient *RangersClient) DataProfileGet(path string) (http.Response, error) {
	return rangersClient.DataProfileFull(path, consts.GET, nil, nil, "")
}
