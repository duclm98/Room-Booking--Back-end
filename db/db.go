package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"echo-demo/config"
	dto "echo-demo/DTOs"
)

var DB *gorm.DB

func init() {
	var err error
	var env config.Config
	
	env, err = config.LoadConfig(".")
	if err != nil {
		fmt.Println("Cannot load database config:", err)
		os.Exit(1)
	}

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=" + env.DbHost + " user=" + env.DbUser + " password=" + env.DbPassword + " dbname=" + env.DbDatabaseName + " port=" + env.DbPort + " sslmode=disable", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Error in connect DB: ", err)
		os.Exit(1)
	}

	DB.AutoMigrate(
		dto.User{},
		dto.Room{},
		dto.Building{},
		dto.Booking{},
	)

	fmt.Println("Connected to DB...")
}