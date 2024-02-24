package env

import (
	"os"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// EnvConfig é a estrutura que armazenará a configuração do ambiente
type Env struct {
	GIN_MODE string `validate:"required,oneof=debug release test" default:"release"`
	PORT     string `validate:"omitempty,number" default:"8000"`
	// Adicione outros campos conforme necessário
}

var config *Env

// InitEnvConfig carrega e inicializa a configuração do ambiente
func InitEnv(logger *logrus.Logger) (*Env, error) {
	// Load environment variables from a .env file
	if err := godotenv.Load("./env/.env"); err != nil {
		logger.Errorf("error loading .env file: %v", err)
		return nil, err
	}

	// Criar uma instância da estrutura de configuração
	config = &Env{}

	// Preencher a configuração dinamicamente a partir das variáveis de ambiente
	if err := fillFromEnv(config); err != nil {
		logger.Errorf("error filling configuration from environment variables: %v", err)
		return nil, err
	}

	// Validate the configuration
	if err := validateConfig(config); err != nil {
		logger.Errorf("validation error: %v", err)
		return nil, err
	}

	return config, nil
}

func fillFromEnv(config interface{}) error {
	configValue := reflect.ValueOf(config).Elem()
	configType := reflect.TypeOf(config).Elem()

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		fieldName := field.Name
		fieldValue := configValue.FieldByName(fieldName)

		// Atribuir valores das variáveis de ambiente à estrutura de configuração
		if envValue := os.Getenv(fieldName); envValue != "" {
			fieldValue.SetString(envValue)
		} else {
			// Definir valor padrão se não houver valor no ambiente
			defaultValue := field.Tag.Get("default")
			if defaultValue != "" {
				fieldValue.SetString(defaultValue)
			}
		}
	}

	return nil
}

func validateConfig(config *Env) error {
	validate := validator.New()
	return validate.Struct(config)
}
