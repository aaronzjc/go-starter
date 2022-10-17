package logger

import (
	"go-starter/pkg/helper"
	"net/http"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

const (
	TYPE_COMMON = "common"
	TYPE_REQ    = "request"
)

type AppLogger struct {
	logger *logrus.Logger
}

var appLogger *AppLogger

func SetLevel(l string) {
	level, _ := logrus.ParseLevel(l)
	appLogger.logger.SetLevel(level)
}

func Fatal(args ...interface{}) {
	appLogger.logger.Fatal(args)
}

func Info(args ...interface{}) {
	appLogger.logger.WithFields(logrus.Fields{
		"log_type": TYPE_COMMON,
	}).Info(args)
}

func Error(args ...interface{}) {
	appLogger.logger.WithFields(logrus.Fields{
		"log_type": TYPE_COMMON,
	}).Error(args)
}

func ErrorWithStack(args ...interface{}) {
	stack := make([]byte, 2048)
	stack = stack[:runtime.Stack(stack, true)]
	appLogger.logger.WithFields(logrus.Fields{
		"log_type": TYPE_COMMON,
		"stack":    string(stack),
	}).Error(args)
}

func Request(req *http.Request, resp *http.Response, ts float64, err error) {
	fields := logrus.Fields{
		"log_type": "request",
		"consume":  ts,
	}
	if req != nil {
		fields["req_host"] = req.URL.Host
		fields["req_path"] = req.URL.Path
		fields["req_params"] = req.URL.RawQuery
	}
	if resp != nil {
		fields["resp_code"] = resp.StatusCode
	}
	if err != nil {
		fields["err"] = err.Error()
		appLogger.logger.WithFields(fields).Error()
	} else {
		appLogger.logger.WithFields(fields).Info()
	}
}

func init() {
	appLogger = &AppLogger{
		logger: logrus.New(),
	}
}

func Setup(appName string, path string) error {
	appLogger.logger.SetFormatter(&logrus.JSONFormatter{})
	appLogger.logger.SetOutput(os.Stdout)
	if path == "" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	appLogger.logger.SetOutput(f)
	appLogger.logger.WithFields(logrus.Fields{
		"app_name": appName,
		"hostname": helper.LocalHostname(),
		"ip":       helper.LocalAddr(),
	})
	return nil
}
