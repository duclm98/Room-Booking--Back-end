package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"echo-demo/config"
)

var env config.Config

func init() {
	var err error
	env, err = config.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load database config:", err)
		return
	}
}

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + env.DbHost + " user=" + env.DbUser + " password=" + env.DbPassword + " dbname=" + env.DbDatabaseName + " port=" + env.DbPort + " sslmode=disable", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	return db, err
}