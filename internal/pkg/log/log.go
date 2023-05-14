package log

import (
	"github.com/cold-runner/skywing-service-center/config"
	"github.com/sirupsen/logrus"
	"os"
)

var Log logrus.Logger

func InitLogger(cfg *config.ServerConf) {
	Log.SetOutput(os.Stdout)
}
