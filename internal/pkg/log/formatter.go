package log

import (
	"errors"
	"github.com/sirupsen/logrus"
)

var InvalidFormatter = errors.New("unsupported logrus formatterMap")

var formatterMap = map[string]logrus.Formatter{
	"text": new(logrus.TextFormatter),
	"json": new(logrus.JSONFormatter),
}

func parseLofFormatter(format string) (logrus.Formatter, error) {
	if v, ok := formatterMap[format]; !ok {
		return nil, InvalidFormatter
	} else {
		return v, nil
	}
}
