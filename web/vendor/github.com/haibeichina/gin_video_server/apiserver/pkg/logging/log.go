package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/haibeichina/gin_video_server/0commonpkg/file"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func SetUp() {
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err := file.MustOpen(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	setPrefix(DEBUG)
	logger.Printf(format, v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

func Infof(format string, v ...interface{}) {
	setPrefix(INFO)
	logger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	setPrefix(WARNING)
	logger.Printf(format, v...)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	setPrefix(ERROR)
	logger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

func Fatalf(format string, v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalf(format, v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		//logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], file, line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
