package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Setup(path string) error {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	if path == "" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	logrus.SetOutput(f)
	return nil
}
