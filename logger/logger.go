package logger

import (
	log "github.com/sirupsen/logrus"
)

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
}
