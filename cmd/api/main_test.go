package main

import (
	"bytes"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite
	apiURL      string
	serverReady chan bool
	serverCmd   *exec.Cmd
	serverOut   bytes.Buffer
	serverErr   bytes.Buffer
}

var serverCmd *exec.Cmd // Declarar a variável fora da função SetupSuite

func (suite *APITestSuite) SetupSuite() {
	suite.apiURL = "http://localhost:8000"
	suite.serverReady = make(chan bool)

	// Iniciar o servidor em uma goroutine
	go func() {
		suite.serverCmd = exec.Command("go", "run", "main.go")
		suite.serverCmd.Stdout = &suite.serverOut
		suite.serverCmd.Stderr = &suite.serverErr

		// Esperar até que o servidor esteja pronto antes de sinalizar
		go func() {
			time.Sleep(2 * time.Second) // Ajuste conforme necessário para garantir a inicialização
			close(suite.serverReady)
		}()

		suite.serverCmd.Start()
	}()

	// Aguardar até que o servidor esteja pronto
	<-suite.serverReady
}

func (suite *APITestSuite) TearDownSuite() {
	if serverCmd != nil && serverCmd.Process != nil {
		serverCmd.Process.Kill()
	}
}

func TestSuite(t *testing.T) {
	// Chamar suite.Run somente após o servidor estar pronto
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
