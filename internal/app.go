package internal

import (
	"github.com/cold-runner/skywing-service-center/config"
)

type App struct {
	config config.ServerConf
	server apiServer
}

func NewApp() *App {
	// 初始化配置文件
	cfg := config.NewConfig()

	return &App{
		config: *cfg,
	}
}

func (a *App) PrepareRun() *App {
	// 初始化服务
	s, err := createServer(&a.config)
	if err != nil {
		panic(err)
	}
	a.server = *s
	return a
}

func (a *App) Run() {
	host := a.config.Server.Host
	port := a.config.Server.Port
	a.server.Run(host + ":" + port)
}
