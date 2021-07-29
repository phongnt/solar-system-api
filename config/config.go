package config

import (
    "flag"
    "os"
)

type DbConfig struct {
    Host        string
    Port        string
    User        string
    Password    string
    Database    string
}

func GetDbConfig() *DbConfig {
    conf := &DbConfig{}

    flag.StringVar(&conf.Host, "hostname", getenv("POSTGRES_HOST", "localhost"), "Database host")
    flag.StringVar(&conf.Port, "port", getenv("POSTGRES_PORT", "5432"), "Database port")
    flag.StringVar(&conf.User, "user", getenv("POSTGRES_USER", "docker"), "Database user")
    flag.StringVar(&conf.Password, "password", getenv("POSTGRES_PASSWORD", "docker"), "Database password")
    flag.StringVar(&conf.Database, "dbname", getenv("POSTGRES_DB", "postgres"), "Database name")

    flag.Parse()

    return conf
}

func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
