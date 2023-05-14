package internal

import (
	"github.com/cold-runner/skywing-service-center/config"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	*gin.Engine
}

func createServer(cfg *config.ServerConf) (*apiServer, error) {
	server := &apiServer{}
	// 设置模式
	switch cfg.Server.Mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
		server.Engine = gin.New()
	case "debug":
		server.Engine = gin.Default()
	default:
		panic("unsupported apiServer mode")
	}
	if cfg.TrustedProxies != nil {
		err := server.Engine.SetTrustedProxies(cfg.TrustedProxies)
		if err != nil {
			panic(err)
		}
	}

	installMiddleware(server.Engine, cfg)
	installController(server.Engine)
	return server, nil
}
