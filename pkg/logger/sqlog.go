package logger

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type DLog struct {
	logger *logrus.Logger
}

func (l *DLog) LogMode(logLevel string) {
	l.logger.Level, _ = logrus.ParseLevel(logLevel)
}

func (l *DLog) Info(ctx context.Context, args ...interface{}) {
	l.logger.Info(args...)
}

func (l *DLog) Warn(ctx context.Context, args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *DLog) Error(ctx context.Context, args ...interface{}) {
	l.logger.Error(args...)
}

func (l *DLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	l.logger.Trace()
}
