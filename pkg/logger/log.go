package logger

import (
	"go-starter/pkg/helper"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type AppLogger struct {
	logger *logrus.Logger
}

func (l *AppLogger) Info(msg string) {
	l.logger.WithField("log_type", "common").Info(msg)
}

func (l *AppLogger) Error(msg string, stack []byte) {
	l.logger.WithFields(logrus.Fields{
		"log_type": "common",
		"stack":    string(stack),
	}).Error(msg)
}

func (l *AppLogger) LogReq(req *http.Request, resp *http.Response, ts float64, err error) {
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
		l.logger.WithFields(fields).Error()
	} else {
		l.logger.WithFields(fields).Info()
	}
}

var appLogger *AppLogger

func init() {
	appLogger = &AppLogger{
		logger: logrus.New(),
	}
}

func Setup(appId string, path string) error {
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
		"app_id":   appId,
		"hostname": helper.LocalHostname(),
		"ip":       helper.LocalAddr(),
	})
	return nil
}

func GetLog() *AppLogger {
	return appLogger
}
