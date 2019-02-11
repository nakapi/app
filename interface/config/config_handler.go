package config

type ConfigHandler interface {
	Load() error
	DatabaseConfig
	LoggerConfig
}

type DatabaseConfig interface {
	GetDatabaseUser() string
	GetDatabasePassword() string
	GetDatabaseHost() string
	GetDatabasePort() string
	GetDatabase() string
}

type LoggerConfig interface {
	GetLogger() []byte
}
