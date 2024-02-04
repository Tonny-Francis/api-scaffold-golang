package envConfig

import (
	"os"
	"reflect"
	"github.com/joho/godotenv"
	"github.com/go-playground/validator/v10"
	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
)

// EnvConfig é a estrutura que armazenará a configuração do ambiente
type EnvConfig struct {
	GIN_MODE string `validate:"required,oneof=debug release test" default:"release"`
	PORT     string `validate:"omitempty,number" default:"8000"`
	// Adicione outros campos conforme necessário
}

var config *EnvConfig

// InitEnvConfig carrega e inicializa a configuração do ambiente
func InitEnvConfig() error {
	// Load environment variables from a .env file
	if err := godotenv.Load(); err != nil {
		loggerHelper.Logger.Errorf("error loading .env file: %v", err)
		return err
	}

	// Criar uma instância da estrutura de configuração
	config = &EnvConfig{}

	// Preencher a configuração dinamicamente a partir das variáveis de ambiente
	if err := fillFromEnv(config); err != nil {
		loggerHelper.Logger.Errorf("error filling configuration from environment variables: %v", err)
		return err
	}

	// Validate the configuration
	if err := validateConfig(config); err != nil {
		loggerHelper.Logger.Errorf("validation error: %v", err)
		return err
	}

	return nil
}

// GetEnvConfig retorna a instância da configuração do ambiente
func GetEnvConfig() *EnvConfig {
	return config
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

func validateConfig(config *EnvConfig) error {
	validate := validator.New()
	return validate.Struct(config)
}