package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ContextTimeout      int    `mapstructure:"CONTEXT_TIMEOUT"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUser              string `mapstructure:"DB_USER"`
	DBPass              string `mapstructure:"DB_PASSWORD"`
	DBName              string `mapstructure:"DB_NAME"`
	DatabaseURL         string `mapstructure:"DB_URL"`
	EmbeddingServiceURL string `mapstructure:"EMBEDDING_SERVICE_URL"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile("../../.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return &env

}
