package config

import (
    "flag"
    "os"
)

type DbConfig struct {
    host   string
    port   string
    user   string
    pass   string
    dbName string
}

func GetDbConfig() *DbConfig {
    conf := &DbConfig{}

    flag.StringVar(&conf.host, "hostname", os.Getenv("POSTGRES_HOST"), "Database Host")
    flag.StringVar(&conf.port, "port", os.Getenv("POSTGRES_PORT"), "Database port")
    flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "Database name")

    flag.Parse()

    return conf
}
