package database

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresDBDSNBuilder(username string, password string, host string, port uint16, dbname string, searchPath string, sslMode string, timeZone string) (dsn string) {
	if len(dbname) == 0 {
		dbname = username
	}
	if len(searchPath) == 0 {
		searchPath = "public"
	}

	if len(timeZone) == 0 {
		timeZone = "UTC"
	} else {
		timeZone = url.QueryEscape(timeZone)
	}
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s search_path=%s sslmode=%s TimeZone=%s",
		username, password, host, port, dbname, searchPath, sslMode, timeZone)
}

func NewPostgresDB(username string, password string, host string, port uint16, dbname string, searchPath string, sslMode string, timeZone string) (db *gorm.DB, err error) {
	dsn := postgresDBDSNBuilder(username, password, host, port, dbname, searchPath, sslMode, timeZone)
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return
	}

	err = applyConnectionPool(db, 2, 10, time.Hour)

	return
}
