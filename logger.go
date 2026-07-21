package cron

import (
	"io/ioutil"
	"log"
	"os"
)

var DefaultLogger Logger = PrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))

var DiscardLogger Logger = PrintfLogger(log.New(ioutil.Discard, "", 0))

type Logger interface {
	Info(msg string, keysAndValues ...interface{})

	Error(err error, msg string, keysAndValues ...interface{})
}

func PrintfLogger(l interface{ Printf(string, ...interface{}) }) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func VerbosePrintfLogger(l interface{ Printf(string, ...interface{}) }) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

type printfLogger struct {
	logger  interface{ Printf(string, ...interface{}) }
	logInfo bool
}

func (pl printfLogger) Info(msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (pl printfLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func formatString(numKeysAndValues int) string { _ = "STUB: not implemented"; return "" }

func formatTimes(keysAndValues []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }
