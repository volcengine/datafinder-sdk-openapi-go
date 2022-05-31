package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"gosdk/dslcontent"
)

var (
	ak     = os.Getenv("ak")
	sk     = os.Getenv("sk")
	url    = os.Getenv("url")
	client = dslcontent.NewRangersClientWithUrl(ak, sk, url)
)

func TestDataTag(t *testing.T) {
	//testUpload()
	//testCreateTag()
	//testQueryResult()
	testQueryHistory()
	//testExportTag()
	//testTagInfo()
	//testQueryTags()
	//testCalTag()
}

func testUpload() {
	method := "POST"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag/file/upload"

	fileDir, _ := os.Getwd()
	fileName := "user_tag.csv"
	filePath := path.Join(fileDir, fileName)

	res, err := client.UploadFile(method, serviceUrl, nil, nil, filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testCreateTag() {
	method := "POST"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag"
	body := "{\n" +
		"    \"name\": \"tag_test_tag_go\",\n" +
		"    \"show_name\": \"测试标签_go\",\n" +
		"    \"value_type\": \"string\",\n" +
		"    \"description\": \"\",\n" +
		"    \"create_type\": \"upload\",\n" +
		"    \"refresh_rule\": \"manual\",\n" +
		"    \"tag_rule\": {\n" +
		"        \"file\": {\n" +
		"            \"file_key\": \"tag_upload_uuid/164314/20220531/90984d7203954cb9b91be7f6a3369059.json\",\n" +
		"            \"detail\": {\n" +
		"                \"name\": \"user_tag.csv\"\n" +
		"            }\n" +
		"        }\n" +
		"    }\n" +
		"}"
	res, err := client.Request(method, serviceUrl, nil, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testQueryResult() {
	method := "GET"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag/tag_test_tag_go/result"
	res, err := client.Request(method, serviceUrl, nil, nil, "")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testQueryHistory() {
	method := "POST"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag/tag_test_tag_go/result/history"
	body := "{\n" +
		"    \"granularity\":\"day\",\n" +
		"    \"type\":\"past_range\",\n" +
		"    \"spans\":[\n" +
		"        {\n" +
		"            \"type\":\"past\",\n" +
		"            \"past\":{\n" +
		"                \"amount\":7,\n" +
		"                \"unit\":\"day\"\n" +
		"            }\n" +
		"        },\n" +
		"        {\n" +
		"            \"type\":\"past\",\n" +
		"            \"past\":{\n" +
		"                \"amount\":1,\n" +
		"                \"unit\":\"day\"\n" +
		"            }\n" +
		"        }\n" +
		"    ],\n" +
		"    \"timezone\":\"Asia/Shanghai\",\n" +
		"    \"week_start\":1\n" +
		"}"
	res, err := client.Request(method, serviceUrl, nil, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testExportTag() {
	method := "POST"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag/tag_test_tag_go/download"
	body := "{\n" +
		"    \"type\": \"user\",\n" +
		"    \"condition\": {\n" +
		"        \"property_operation\": \"is_not_null\",\n" +
		"        \"snapshot\": {\n" +
		"            \"type\": \"day\",\n" +
		"            \"day\": \"2022-05-31\"\n" +
		"        }\n" +
		"    },\n" +
		"    \"period\": {\n" +
		"        \"timezone\": \"Asia/Shanghai\"\n" +
		"    }\n" +
		"}"
	res, err := client.Request(method, serviceUrl, nil, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testTagInfo() {
	method := "GET"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag/tag_test_tag_go"
	body := ""
	res, err := client.Request(method, serviceUrl, nil, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testQueryTags() {
	method := "GET"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag"
	body := ""
	res, err := client.Request(method, serviceUrl, nil, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}

func testCalTag() {
	method := "POST"
	serviceUrl := "/datatag/openapi/v1/app/164314/tag/tag_test_tag_go/calculation"
	body := ""
	res, err := client.Request(method, serviceUrl, nil, nil, body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println(err, string(data))
}
