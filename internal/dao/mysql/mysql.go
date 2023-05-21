package mysql

import (
	"fmt"
	"github.com/cold-runner/skywing-service-center/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// New create a new gorm db instance with the given options.
func New(cfg *config.Mysql) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(cfg.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnections)

	return db, nil
}
