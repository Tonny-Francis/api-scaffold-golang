package main

import (
	"net/http"
	"testing"

	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/testHelper"
)

func TestMain(m *testing.M) {
	loggerHelper.InitLogger()
	testHelper.Init(m)
}

func TestRouteExample(t *testing.T) {
	testHelper.TestRoute(t, "Teste de Funcionamento", "GET", "/v1/example", http.StatusFailedDependency)
	testHelper.TestRoute(t, "Teste", "GET", "/v1/example", http.StatusBadRequest)
}
