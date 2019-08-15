package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/nsysu/teacher-education/src/utils/config"
)

var (
	logger       *log.Logger
	file         *os.File
	loggerRegexp = regexp.MustCompile(`src/utils/logger/.*.go`)
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)

	logPath := fmt.Sprintf("%s/../../../%s%s", dir, config.Get("log.path"), config.Get("log.file"))
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	} else {
		logger = log.New(logFile, "", log.LstdFlags)
		logger.SetPrefix("[ ")
	}
}

// Close should be called before main process ends.
func Close() {
	if file != nil {
		file.Close()
	}
	file, logger = nil, nil
}

// Debug outputs the message into log file
func Debug(msg interface{}) {
	output("DEBUG", msg)
}

// Info outputs the message into log file
func Info(msg interface{}) {
	output("INFO", msg)
}

// Warn outputs the message into log file
func Warn(msg interface{}) {
	output("WARN", msg)
}

// Error outputs the message into log file
func Error(msg interface{}) {
	output("ERROR", msg)
}

func output(level string, msg interface{}) {
	logger.Println(fmt.Sprintf("] %s - %s : %v", level, callStack(), msg))
}

func callStack() string {
	preResult, preFile := "", ""
	stack := strings.Split(string(debug.Stack()), "\n")

	for i := 1; i < len(stack)-1; i += 2 {
		file, line := splitStack(stack[i+1])

		if !loggerRegexp.MatchString(file) && strings.Contains(file, config.Get("server.dir").(string)) {
			// trim the stack's absolute path to relative path
			// e.g before /path/to/project/src/utils/logger/logger.go
			//     after  src/utils/logger/logger.go
			temp := strings.Split(file, config.Get("server.dir").(string))
			file := temp[len(temp)-1]
			result := fmt.Sprintf("File \"%v\", line %v", file, line)

			if preFile != "" && file != preFile {
				return preResult
			}
			preFile = file
			preResult = result
		}
	}

	return preResult
}

func splitStack(str string) (file string, line string) {
	temp := strings.Split(strings.TrimSpace(str), ":")
	file = temp[0]

	temp = strings.Split(temp[1], " ")
	line = temp[0]

	return
}
