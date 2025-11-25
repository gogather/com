package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger struct {
	log.Logger
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return New(out, prefix, flag)
}

func (logger *Logger) Fatal(values ...interface{}) {
	logger.Logger.Output(2, fmt.Sprint(values...))
	os.Exit(1)
}

func (logger *Logger) Fatalf(format string, values ...interface{}) {
	logger.Logger.Output(2, fmt.Sprintf(format, values...))
	os.Exit(1)
}

func (logger *Logger) Fatalln(values ...interface{}) {
	logger.Logger.Output(2, fmt.Sprintln(values...))
	os.Exit(1)
}

func (logger *Logger) Panic(values ...interface{}) {
	message := fmt.Sprint(values...)
	logger.Logger.Output(2, message)
	panic(message)
}

func (logger *Logger) Panicf(format string, values ...interface{}) {
	message := fmt.Sprintf(format, values...)
	logger.Logger.Output(2, message)
	panic(message)
}

func (logger *Logger) Panicln(values ...interface{}) {
	message := fmt.Sprintln(values...)
	logger.Logger.Output(2, message)
	panic(message)
}

func (logger *Logger) Flags() int {
	return logger.Logger.Flags()
}

func (logger *Logger) SetFlags(flag int) {
	logger.Logger.SetFlags(flag)
}

func (logger *Logger) Prefix() string {
	return logger.Logger.Prefix()
}

func (logger *Logger) SetPrefix(prefix string) {
	logger.Logger.SetPrefix(prefix)
}
