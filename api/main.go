package main

import (
	"github.com/nsysu/teacher-education/src/persistence/gorm"
	"github.com/nsysu/teacher-education/src/persistence/redis"
	"github.com/nsysu/teacher-education/src/route"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

func main() {
	defer logger.Close()
	defer gorm.Close()
	defer redis.Close()

	route.Run()
}
