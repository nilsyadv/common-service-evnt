package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nilsyadv/common-service-evnt/pkg/config"
)

// Database Struct
type DB struct {
	dbs *gorm.DB
}

// initiating db instance
func Initdb(config config.IConfig) (*DB, error) {
	dbs, err := createdb(config)
	if err != nil {
		log.Println("Failed to Initiat DB Instance: ", err)
		return nil, err
	}
	log.Println("database insatnce created successfully.....")
	return &DB{
		dbs: dbs,
	}, nil
}

func (db *DB) Getdb() *gorm.DB {
	return db.dbs
}

// set database debug
func (db *DB) SetDebug() {
	db.dbs = db.dbs.Debug()
}

// Will Return gorm Instance
func createdb(config config.IConfig) (*gorm.DB, error) {
	dbconn := createDBString(config)
	fmt.Println("db Connection String:", dbconn)
	db, err := gorm.Open(postgres.Open(dbconn), &gorm.Config{})
	if err != nil {
		log.Println("Error in db Connection start:", err.Error())
		log.Fatal("Failed to connect db....")
		return nil, err
	}
	return db, nil
}

// Creating db connection string
func createDBString(config config.IConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.GetString("db.db_host"), config.GetString("db.db_user"),
		config.GetString("db.db_pass"), config.GetString("db.db_name"),
		config.GetString("db.db_port"), config.GetString("db.db_ssl"),
		config.GetString("db.db_timezone"))
}

// migrate db table.
func (db *DB) AutoMigrate(in interface{}) error {
	return db.dbs.AutoMigrate(in)
}

// migrate db multiple table.
func (db *DB) AutoMigrates(ins []interface{}) []error {
	var errs []error
	for _, in := range ins {
		er := db.dbs.AutoMigrate(in)
		if er != nil {
			errs = append(errs, er)
		}
	}
	return errs
}
