package config

import (
	v1 "github.com/Gradient-and-Co/sigma-school/internal/adapter/delivery/http/v1"
	"github.com/Gradient-and-Co/sigma-school/pkg/logging"
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	Logging logging.Config
	Web     v1.Config
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath("config")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		if err := bindEnvConfig(); err != nil {
			log.Fatalf("error reading config file: %v", err)
		}

		log.Println("read config file: config/config.yml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("error reading config file: %v", err)
		}

		instance = &Config{}
		if err := viper.Unmarshal(&instance); err != nil {
			log.Fatalf("error unmarshaling config file: %v", err)
		}
	})
	return instance
}

func bindEnvConfig() error {
	bindings := make(map[string]string)
	bindings["web.host"] = "HOST"
	bindings["web.port"] = "PORT"

	for name, binding := range bindings {
		if err := viper.BindEnv(name, binding); err != nil {
			return err
		}
	}

	return nil
}
