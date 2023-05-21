package internal

import (
	"github.com/cold-runner/skywing-service-center/config"
	"github.com/cold-runner/skywing-service-center/internal/dao/mysql"
	"github.com/cold-runner/skywing-service-center/internal/pkg/log"
	"github.com/gin-gonic/gin"
)

type apiServer struct {
	*gin.Engine
}

func createServer(cfg *config.ServerConf) (server *apiServer, err error) {
	server = &apiServer{}
	// 初始化日志
	log.InitLogger(cfg)

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
	// 设置信任代理
	if cfg.TrustedProxies != nil {
		err := server.Engine.SetTrustedProxies(cfg.TrustedProxies)
		if err != nil {
			panic(err)
		}
	}
	// 初始化数据库
	mysql.Db, err = mysql.New(&cfg.Mysql)
	if err != nil {
		panic(err)
	}
	installMiddleware(server.Engine, cfg)
	installController(server.Engine)
	return server, nil
}
