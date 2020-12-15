package gorm

import (
	"fmt"
	"time"

	"github.com/nsysu/teacher-education/src/utils/config"
	"github.com/nsysu/teacher-education/src/utils/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Get("db.user"),
		config.Get("db.password"),
		config.Get("db.host"),
		config.Get("db.port"),
		config.Get("db.database"),
		config.Get("db.flag"),
	)
)

func init() {
	db = connect(dsn)
}

// DB return database instance
func DB() *gorm.DB {
	if db == nil {
		connect(dsn)
	}
	return db
}

// Close close database connection
func Close() {
	if db != nil {
		sqldb, _ := db.DB()
		if err := sqldb.Close(); err != nil {
			logger.Error(err.Error())
		}
	}
}

func connect(dsn string) *gorm.DB {
	conn, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			AllowGlobalUpdate: false,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)

	if err != nil {
		panic(err)
	}

	sqldb, err := conn.DB()

	if err != nil {
		panic(err)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqldb.SetConnMaxLifetime(time.Hour)

	return conn
}
