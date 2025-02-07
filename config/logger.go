package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func newLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "->DEBUG: ", logger.Flags()),
		info:    log.New(writer, "->INFO: ", logger.Flags()),
		warning: log.New(writer, "->WARNING: ", logger.Flags()),
		err:     log.New(writer, "->ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create non-foramted logs

func (l *Logger) Debug(v ...interface{}) {

	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {

	l.debug.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {

	l.debug.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {

	l.debug.Println(v...)
}
