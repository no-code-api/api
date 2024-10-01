package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	fatal   *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, formatPrefix(prefix, "DEBUG"), logger.Flags()),
		info:    log.New(writer, formatPrefix(prefix, "INFO"), logger.Flags()),
		warning: log.New(writer, formatPrefix(prefix, "WARNING"), logger.Flags()),
		err:     log.New(writer, formatPrefix(prefix, "ERROR"), logger.Flags()),
		fatal:   log.New(writer, formatPrefix(prefix, "FATAL"), logger.Flags()),
		writer:  writer,
	}
}

func formatPrefix(prefix string, logType string) string {
	return fmt.Sprintf("[%s] %s: ", prefix, logType)
}

// Create Non-formatted Logs
func (logger *Logger) Debug(values ...interface{}) {
	logger.debug.Println(values...)
}

func (logger *Logger) Info(values ...interface{}) {
	logger.info.Println(values...)
}

func (logger *Logger) Warning(values ...interface{}) {
	logger.warning.Println(values...)
}

func (logger *Logger) Error(values ...interface{}) {
	logger.err.Println(values...)
}

func (logger *Logger) Fatal(values ...interface{}) {
	logger.fatal.Fatal(values...)
}

// Create Formated Logs
func (logger *Logger) DebugF(format string, values ...interface{}) {
	logger.debug.Printf(format, values...)
}

func (logger *Logger) InfoF(format string, values ...interface{}) {
	logger.info.Printf(format, values...)
}

func (logger *Logger) WarningF(format string, values ...interface{}) {
	logger.warning.Printf(format, values...)
}

func (logger *Logger) ErrorF(format string, values ...interface{}) {
	logger.err.Printf(format, values...)
}

func (logger *Logger) FatalF(format string, values ...interface{}) {
	logger.fatal.Fatalf(format, values...)
}
