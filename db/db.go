package db

import (
    "fmt"
    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
    "log"
    "topcoder.com/skill-builder/golang/models"
)

var db *pg.DB

func Init() {
    db = newDBConn()
    err := createSchema(db)
    if err != nil {
        panic(err)
    } else {
        log.Default().Println("Database initialization completed...")
    }
}

func GetDBObject() *pg.DB {
    return db
}

func newDBConn() (con *pg.DB) {
    //TODO: Easy-250: Read these from environment variables (check config.go)
    address := fmt.Sprintf("%s:%s", "localhost", "5432")
    options := &pg.Options {
        User: "docker",
        Password: "docker",
        Addr: address,
        Database: "postgres",
        PoolSize: 50,
    }

    con = pg.Connect(options)
    if con == nil {
        log.Fatal("Failed to connect to database")
    }
    log.Println("Successfully connected to database.")
    return con
}

func createSchema(db *pg.DB) error {
    dbModels := [] interface{} {
        (*models.Body)(nil),
    }

    for _, model := range dbModels {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{
            IfNotExists: true,
        })
        if err != nil {
            return err
        }
    }

    return nil
}
