package testHelper

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

var serverCmd *exec.Cmd

func startServer() {
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
		os.Exit(1)
	}
	serverCmd = cmd
	time.Sleep(2 * time.Second)
}

func stopServer() {
	if serverCmd != nil && serverCmd.Process != nil {
		err := serverCmd.Process.Signal(os.Interrupt)
		if err != nil {
			fmt.Println("Erro ao enviar sinal de interrupção ao servidor:", err)
		}

		// Aguarde o término do servidor
		err = serverCmd.Wait()
		if err != nil {
			fmt.Println("Erro ao aguardar o término do servidor:", err)
		}
	}
}

func Init(m *testing.M) {
	startServer()

	exitCode := m.Run()

	stopServer()

	os.Exit(exitCode)
}

func TestRoute(t *testing.T, test string, method string, path string, expectedStatus int) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(
		method,
		"http://localhost:8000"+path,
		nil,
	)

	if err != nil {
		t.Fatal("Erro ao criar a solicitação HTTP:", err)
	}

	res, err := client.Do(req)

	if err != nil {
		t.Fatal("Erro ao fazer a solicitação HTTP:", err)
	}

	defer res.Body.Close()

	assert.NoError(t, err, "Erro ao fazer a solicitação HTTP")

	assert.Equal(t, expectedStatus, res.StatusCode, test)

	if res.StatusCode == expectedStatus {
		t.Logf("SUCCESS: %s", test)
	}
}
