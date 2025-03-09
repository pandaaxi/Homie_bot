package logger

import (
	"log"
	"os"
)

// Init configures the standard logger.
func Init(logLevel string) {
	// For a production system you might configure different loggers and levels.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.Printf("Logger initialized with level: %s", logLevel)
}
