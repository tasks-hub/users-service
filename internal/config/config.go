package config

type Config struct {
	Environment              string `env:"ENVIRONMENT,required=true"`
	DatabaseType             string `env:"DATABASE_TYPE,required=true"`
	UserServicePort          string `env:"USER_SERVICE_PORT,required=true"`
	UserServiceDatabaseTable string `env:"USER_SERVICE_DATABASE_TABLE,required=true"`
	Database
}

type Database struct {
	PostgresHostFile     string `env:"POSTGRES_HOST_FILE"`
	PostgresDBFile       string `env:"POSTGRES_DB_FILE"`
	PostgresUserFile     string `env:"POSTGRES_USER_FILE"`
	PostgresPasswordFile string `env:"POSTGRES_PASSWORD_FILE"`
}
