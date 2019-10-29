package gorm

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql" // init mysql driver
	"github.com/jinzhu/gorm"
	"github.com/nsysu/teacher-education/src/utils/config"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

var (
	db       *gorm.DB
	interval = config.Get("db.interval").(int)
	dialect  = config.Get("db.dialect").(string)
	source   = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.Get("db.user"),
		config.Get("db.password"),
		config.Get("db.host"),
		config.Get("db.port"),
		config.Get("db.database"),
		config.Get("db.flag"),
	)
)

func init() {
	db = connect(dialect, source)
}

// DB return database instance
func DB() *gorm.DB {
	if db == nil {
		connect(dialect, source)
	}
	return db
}

// Close close database connection
func Close() {
	if db != nil {
		if err := db.Close(); err != nil {
			logger.Error(err.Error())
		}
	}
}

func connect(dialect string, source string) *gorm.DB {
	conn, err := gorm.Open(dialect, source)

	if err != nil {
		panic(err)
	}

	// use singular table by default
	conn.SingularTable(true)

	// generates an error on update/delete without where clause.
	// prevent eventual error with empty objects updates/deletions
	conn.BlockGlobalUpdate(true)

	// sets the maximum number of connections in the idle connection pool.
	conn.DB().SetMaxIdleConns(10)

	// sets the maximum number of open connections to the database.
	conn.DB().SetMaxOpenConns(100)

	// sets the maximum amount of time a connection may be reused.
	conn.DB().SetConnMaxLifetime(time.Hour)

	// the number of seconds the server waits for activity on a noninteractive connection before closing it.
	conn.Exec("SET wait_timeout=300")

	return conn
}
