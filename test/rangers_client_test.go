package test

import (
	"fmt"
	"gosdk/dslcontent"
	"io/ioutil"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/suite"
)

type RangerClientSuite struct {
	suite.Suite
	ak     string
	sk     string
	client dslcontent.RangersClient
}

func (suite *RangerClientSuite) SetupSuite() {
	suite.ak = "***"
	suite.sk = "***"
	suite.client = dslcontent.NewRangersClient(suite.ak, suite.sk)
}

func (suite *RangerClientSuite) AnalysisRequest(dsl dslcontent.DSL) {
	b, err := jsoniter.Marshal(dsl)
	if err != nil {
		fmt.Println(err.Error(), string(b))
	}
	res, err := suite.client.DataFinderFull("/openapi/v1/analysis", "post", nil, nil, string(b))
	if err != nil {
		fmt.Println(err)
	}

	bodyBytes, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println("body", string(bodyBytes))

	//suite.Equal(200, res.StatusCode)
	fmt.Println(res.StatusCode)
}

func (suite *RangerClientSuite) TestRangersOpenAPI() {
	res, err := suite.client.DataRangers("/openapi/v1/xxx/date/2020-02-20/2020-02-23/downloads", "get")
	fmt.Println(res, err)
}

func TestRangerClient(t *testing.T) {
	suite.Run(t, new(RangerClientSuite))
}
