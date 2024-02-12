package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/testHelper"
)

func TestMain(m *testing.M) {
	loggerHelper.InitLogger()

	exit := testHelper.Init(m)

	os.Exit(exit)
}

func TestRouteExample(t *testing.T) {
	testHelper.TestRoute(t, "Teste", "GET", "/v1/example", http.StatusBadRequest)
}
