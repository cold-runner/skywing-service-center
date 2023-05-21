package internal

import (
	"github.com/cold-runner/skywing-service-center/config"
	"github.com/cold-runner/skywing-service-center/internal/dao/mysql"
	"github.com/cold-runner/skywing-service-center/internal/middleware"
	"github.com/cold-runner/skywing-service-center/internal/model"
	"github.com/cold-runner/skywing-service-center/internal/pkg/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(r *gin.Engine, cfg *config.ServerConf) {

}

func installMiddleware(r *gin.Engine, cfg *config.ServerConf) {
	m := cfg.Server.Middlewares
	for _, v := range m {
		if install, ok := middleware.Middlewares[v]; ok {
			r.Use(install())
		}
	}
}

func installController(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		log.Log.Info("access health api!")
		c.String(http.StatusOK, "I'm alive!")
	})

	r.GET("/getSkyAppliers", func(c *gin.Context) {
		var appliers []model.Applier
		mysql.Db.Find(&appliers)
		c.JSON(http.StatusOK, appliers)
	})
}
