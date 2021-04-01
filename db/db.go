package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres password=MinhDuc dbname=RoomBooking port=5432 sslmode=disable", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	return db, err
}