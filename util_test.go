package main

import (
	"testing"
	. "gopkg.in/check.v1"
)

func TestUtilMethods(t *testing.T) {
	TestingT(t)
}

type UtilTestSuite struct{}

var _ = Suite(&UtilTestSuite{})

func (suite *UtilTestSuite) TestGetSettings(c *C) {
	testSettings := GetSettings("testSettings.json")

	c.Assert(len(testSettings.Services), Equals, 2)

	c.Assert(testSettings.Port, Equals, 1008)
	c.Assert(testSettings.Hostname, Equals, "testhost")

	testService := testSettings.Services[0]
	c.Assert(testService.Name, Equals, "couchdb")
	c.Assert(testService.Port, Equals, 5984)

	testService = testSettings.Services[1]
	c.Assert(testService.Name, Equals, "Google")
	c.Assert(testService.Port, Equals, 443)
}

func (suite *UtilTestSuite) TestStatusMatrixGenerator(c *C) {
	testSettings := GetSettings("testSettings.json")

	testStatusMatrix := StatusMatrixGenerator(testSettings.Services)

	c.Check(testStatusMatrix[0].Name, Equals, testSettings.Services[0].Label)
	c.Assert(testStatusMatrix[0].Status, Equals, "DOWN")

	c.Check(testStatusMatrix[1].Name, Equals, testSettings.Services[1].Label)
	c.Assert(testStatusMatrix[1].Status, Equals, "UP")

}
