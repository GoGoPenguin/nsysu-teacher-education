package main

import (
	"github.com/nsysu/teacher-education/src/route"
	"github.com/nsysu/teacher-education/src/utils/logger"
)

func main() {
	defer logger.Close()

	route.Run()
}
