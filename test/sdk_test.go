package test

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/volcengine/datafinder-sdk-openapi-go/dslcontent"
)

//单元测试版本

type ExampleSDKSuite struct {
	suite.Suite
	ak string
	sk string
	id string
}

func (suite *ExampleSDKSuite) SetupSuite() {
	suite.ak = "xxx"
	suite.sk = "xxx"
	suite.id = "xxx"
}

func (suite *ExampleSDKSuite) TestGetDataExportUrl() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	res, err := client.DataRangers("/openapi/v1/xxx/date/2020-05-03/2020-05-09/downloads", "get")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderAll() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	fmt.Println(suite.ak, suite.sk)
	res, err := client.DataFinder("/openapi/v1/xxx/dashboards/all", "get")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderDashboard() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	res, err := client.DataFinder("/openapi/v1/xxx/dashboards/6760954832762176007", "get")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderDashboardReports() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	res, err := client.DataFinder("/openapi/v1/xxx/dashboards/6760954832762176007/reports", "get")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderDashboardReport() {
	id := suite.id
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	headers := map[string]string{}
	params := map[string]string{"count": "10"}
	res, err := client.DataFinderFull("/openapi/v1/xxx/reports/"+id, "get", headers, params, "")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderCohorts() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	res, err := client.DataFinder("/openapi/v1/xxx/cohorts", "get")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderCohort() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	res, err := client.DataFinder("/openapi/v1/xxx/cohorts/4627", "get")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func (suite *ExampleSDKSuite) TestFinderCohortsSample() {
	client := dslcontent.NewRangersClient(suite.ak, suite.sk)
	headers := map[string]string{}
	params := map[string]string{"count": "10"}
	res, err := client.DataFinderFull("/openapi/v1/xxx/cohorts/4453/sample", "get", headers, params, "")
	suite.Nil(err)
	suite.Equal(200, res.StatusCode)
	data, err := ioutil.ReadAll(res.Body)
	fmt.Println("data----", string(data))
	fmt.Println("error---", err)
}

func TestExampleSDKSuite(t *testing.T) {
	suite.Run(t, new(ExampleSDKSuite))
}
