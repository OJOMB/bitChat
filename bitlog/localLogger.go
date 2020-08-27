package bitlog

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/OJOMB/bitChat/bitutils"
)

var once sync.Once

// LocalLogger is a logger implementation that generates local log files to bitChat/bitlog/bitLogFiles
type LocalLogger struct {
	filename string
	warning  *log.Logger
	info     *log.Logger
	err      *log.Logger
	fatal    *log.Logger
}

// Info level log message
func (l *LocalLogger) Info(msg ...interface{}) {
	passToPrint(l.info, msg...)
}

// Warn level log message
func (l *LocalLogger) Warn(msg ...interface{}) {
	passToPrint(l.warning, msg...)
}

// Error level log message
func (l *LocalLogger) Error(msg ...interface{}) {
	passToPrint(l.err, msg...)
}

// Fatal level log message, logs then exits
func (l *LocalLogger) Fatal(msg ...interface{}) {
	passToPrint(l.fatal, msg...)
	os.Exit(1)
}

func passToPrint(logger *log.Logger, args ...interface{}) {
	if len(args) > 1 {
		logger.Printf(args[0].(string)+"\n", args[1:]...)
		return
	}
	logger.Println(args[0])
}

// GetLocalLogger returns a local logger
func GetLocalLogger() *LocalLogger {
	var globalLocalLogger *LocalLogger
	once.Do(func() { globalLocalLogger = initLocalLogger() })
	return globalLocalLogger
}

func initLocalLogger() *LocalLogger {
	filename := "log-" + strings.ReplaceAll(time.Now().UTC().Format(time.UnixDate), " ", "_")
	filepath := bitutils.ConstructPath(
		[]string{"wikiLog", "wikiLogFiles", filename},
		".log",
	)
	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var WarningLogger *log.Logger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	var InfoLogger *log.Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	var ErrorLogger *log.Logger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	var FatalLogger *log.Logger = log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	return &LocalLogger{
		filename: filename,
		warning:  WarningLogger,
		info:     InfoLogger,
		err:      ErrorLogger,
		fatal:    FatalLogger,
	}
}
