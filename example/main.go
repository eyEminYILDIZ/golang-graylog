package main

import (
	graylog "github.com/eyEminYILDIZ/golang-graylog"
)

func init() {
	graylog.InitLogger(graylog.Config{
		Address:         "127.0.0.1",
		Port:            12201,
		ApplicationName: "golang-console",
		ShowLogs:        true,
	})
}

func main() {
	graylog.Debug("Debug log...")
	graylog.Information("Information log...")
	graylog.Warning("Warning log...")
	graylog.Error("Error log...")
}
