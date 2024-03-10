package env

import (
	"os"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env struct {
	GIN_MODE string `validate:"required,oneof=debug release test" default:"release"`
	PORT     string `validate:"omitempty,number" default:"8000"`
}

var config *Env

func InitEnv(logger *logrus.Logger, mode string) (*Env, error) {
	var path string

	if mode == "test" {
		path = "../../env/.env"
	} else {
		path = "./env/.env"
	}

	if err := godotenv.Load(path); err != nil {
		logger.Errorf("error loading .env file: %v", err)
		return nil, err
	}

	config = &Env{}

	if err := fillFromEnv(config); err != nil {
		logger.Errorf("error filling configuration from environment variables: %v", err)
		return nil, err
	}

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

		if envValue := os.Getenv(fieldName); envValue != "" {
			fieldValue.SetString(envValue)
		} else {
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
