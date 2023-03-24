package graylog

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

var gelfWriter *gelf.TCPWriter
var extra map[string]interface{}
var config Config

func InitLogger(_config Config) {
	config = _config
	graylogAddr := fmt.Sprintf("%s:%d", config.Address, config.Port)
	var err error
	gelfWriter, err = gelf.NewTCPWriter(graylogAddr)
	if err != nil {
		log.Fatalf("gelf.NewTCPWriter: %s", err)
	}

	extra = make(map[string]interface{})
	extra["Application"] = "golang-console"
}

func send(message string, level int32) {

	gelfWriter.WriteMessage(&gelf.Message{
		Facility: "golang-console",
		Extra:    extra,
		Version:  "2.0",
		Host:     "golang",
		Short:    message,
		Full:     "Messages...",
		TimeUnix: float64(time.Now().Unix()),
		Level:    level,
	})

	if config.ShowLogs {
		log.Println(message)
	}
}

func Debug(message string) {
	send(message, gelf.LOG_DEBUG)
}

func Information(message string) {
	send(message, gelf.LOG_ERR)
}

func Warning(message string) {
	send(message, gelf.LOG_WARNING)
}

func Error(message string) {
	send(message, gelf.LOG_ERR)
}
