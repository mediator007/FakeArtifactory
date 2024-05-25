package log

import (
	"log"
)

// ANSI color codes for colored logging
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

// Wrapper functions to include log levels with colors
func Info(v ...interface{}) {
	log.SetPrefix(Blue + "INFO: " + Reset)
	log.Println(v...)
}

func Success(v ...interface{}) {
	log.SetPrefix(Green + "SUCCESS: " + Reset)
	log.Println(v...)
}

func Warn(v ...interface{}) {
	log.SetPrefix(Yellow + "WARN: " + Reset)
	log.Println(v...)
}

func Error(v ...interface{}) {
	log.SetPrefix(Red + "ERROR: " + Reset)
	log.Println(v...)
}
