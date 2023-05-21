package log

import (
	"github.com/cold-runner/skywing-service-center/config"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func InitLogger(cfg *config.ServerConf) {
	level, err := logrus.ParseLevel(cfg.Log.Level)
	if err != nil {
		panic(err)
	}
	formatter, err := parseLofFormatter(cfg.Format)
	if err != nil {
		panic(err)
	}

	Log = &logrus.Logger{
		Out:       os.Stderr,
		Formatter: formatter,
		Level:     level,
	}
}
