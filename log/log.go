package log

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {

	fmt.Println("thwerthwrth")

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
}

// Info ...
func Info(args ...interface{}) {
	log.Info(args)
}

// Error ...
func Error(args ...interface{}) {
	log.Error(args)
}

// Fatal ...
func Fatal(args ...interface{}) {
	log.Fatal(args)
}
