package logger

import (
	"log"
)

// returning defailt logger
func NewLogger() *log.Logger {
	return log.Default()
}
