package test

import (
	"gosdk/util"
	"strconv"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/suite"
)

type DslSignTestSuite struct {
	suite.Suite
	accessKey         string
	secrteKey         string
	expirationSeconds int
}

func (suite *DslSignTestSuite) SetupSuite() {
	suite.accessKey = "****"
	suite.secrteKey = "****"
	suite.expirationSeconds = 300
}

func (suite *DslSignTestSuite) TestSign() {
	method := "POST"
	uri := "/dataprofile/openapi/v1/xx/users/xxx"
	queryParams := map[string]string{"set_once": "true"}
	QueryBodyJson := "{\"name\":\"name\",\"value\":\"zhangsan\"}"
	patches := gomonkey.ApplyFunc(strconv.FormatInt, func(_ int64, _ int) string {
		return "1596784160"
	})
	suite.expirationSeconds = 1596784160
	defer patches.Reset()
	authorization := util.Sign(suite.accessKey, suite.secrteKey, suite.expirationSeconds, method, uri, queryParams, QueryBodyJson)
	suite.Equal("ak-v1/****/1596784160/1596784160/xxx", authorization)
}

func TestDslSignTestSuite(t *testing.T) {
	suite.Run(t, new(DslSignTestSuite))
}
