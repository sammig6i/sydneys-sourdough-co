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
	BackendPort         string `mapstructure:"BACKEND_PORT"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetDefault("CONTEXT_TIMEOUT", 30)
	viper.SetDefault("BACKEND_PORT", ":8080")
	viper.SetDefault("DB_NAME", "postgres")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Printf("Warning: Error unmarshaling config: %v", err)
	}

	if env.BackendPort != "" {
		env.BackendPort = ":" + env.BackendPort
	}

	return &env
}
