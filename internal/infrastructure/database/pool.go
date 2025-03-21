package database

import (
	"time"

	"gorm.io/gorm"
)

func applyConnectionPool(db *gorm.DB, maxIdleCons int, maxOpenCons int, connMaxLifetime time.Duration) (err error) {
	sqlDB, err := db.DB()
	if err != nil {
		return
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleCons)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenCons)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	return
}
