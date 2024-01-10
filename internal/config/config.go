package config

type Config struct {
	Environment string `env:"ENVIRONMENT,required=true"`
}
