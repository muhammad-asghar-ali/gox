package config

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type (
	LogFormatter struct {
		TimestampFormat string
		LevelDesc       []string
	}
)

func (f *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(f.TimestampFormat)

	return []byte(fmt.Sprintf("%s %s %s\n", timestamp, f.LevelDesc[entry.Level], entry.Message)), nil
}

func NewLogger() {
	log.SetOutput(os.Stdout)
	format := new(LogFormatter)
	format.TimestampFormat = "2006-01-02 15:04:05"
	format.LevelDesc = []string{"PANIC", "FATAL", "ERROR", "WARN", "INFO", "DEBUG", "TRACE"}
	log.SetFormatter(format)
}
