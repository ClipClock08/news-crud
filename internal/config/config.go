package config

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"
)

const (
	defaultServerPort         = "8080"
	defaultServerReadTimeout  = 5 * time.Second
	defaultServerWriteTimeout = 10 * time.Second
	defaultServerIdleTimeout  = 1 * time.Second
)

type Environment string

const (
	Development Environment = "development"
	Stage       Environment = "stage"
	Production  Environment = "production"
)

type Config struct {
	Server Server
	DB     Database
}

type Server struct {
	Environment  Environment
	LogLevel     slog.Leveler
	Address      string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Database struct {
	ConnectionString string
	Name             string
}

func LoadConfigFromEnv() (Config, error) {
	server, err := loadServerConfig()
	if err != nil {
		return Config{}, err
	}

	db, err := loadDatabaseConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		Server: server,
		DB:     db,
	}, nil
}

func loadServerConfig() (Server, error) {
	var env Environment
	var logLevel slog.Leveler

	port, err := mustParseEnvVar("SERVER_PORT")
	if err != nil {
		port = defaultServerPort
	}

	e := os.Getenv("ENVIRONMENT")

	switch e {
	case "production":
		env = Production
		logLevel = slog.LevelError
	case "stage":
		env = Stage
		logLevel = slog.LevelInfo
	default:
		env = Development
		logLevel = slog.LevelDebug
	}

	return Server{
		Environment:  env,
		LogLevel:     logLevel,
		Address:      net.JoinHostPort("", port),
		Port:         port,
		ReadTimeout:  defaultServerReadTimeout,
		WriteTimeout: defaultServerWriteTimeout,
		IdleTimeout:  defaultServerIdleTimeout,
	}, nil
}

func loadDatabaseConfig() (Database, error) {
	connStr, err := mustParseEnvVar("CONNECTION_STRING")
	if err != nil {
		return Database{}, err
	}

	dbName, err := mustParseEnvVar("DB_NAME")
	if err != nil {
		return Database{}, err
	}

	return Database{
		ConnectionString: connStr,
		Name:             dbName,
	}, nil
}

func mustParseEnvVar(key string) (string, error) {
	if value, exist := os.LookupEnv(key); exist {
		return value, nil
	}

	return "", fmt.Errorf("environment variable %s not set", key)
}
