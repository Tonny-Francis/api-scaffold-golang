package main

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	apiURL string
}

func (suite *APITestSuite) SetupTest() {
	suite.apiURL = "http://localhost:8000"
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

func (suite *APITestSuite) TestHealthz() {
	req, err := http.NewRequest("GET", suite.apiURL+"/healthz", nil)

	assert.NoError(suite.T(), err, "No error expected")

	resp, err := http.DefaultClient.Do(req)

	assert.NoError(suite.T(), err, "No error expected")

	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode, "Status code should be 200")
}

func (suite *APITestSuite) TestExample() {
	reqquestBody := []byte(`{"name":"Tonny","phone":"1234567890"}`)

	req, err := http.NewRequest("POST", suite.apiURL+"/v1/example/0000?page=12", bytes.NewBuffer(reqquestBody))

	assert.NoError(suite.T(), err, "No error expected")

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("tenant", "test")

	resp, err := http.DefaultClient.Do(req)

	assert.NoError(suite.T(), err, "No error expected")

	defer resp.Body.Close()

	assert.Equal(suite.T(), http.StatusOK, resp.StatusCode, "Status code should be 200")
}
