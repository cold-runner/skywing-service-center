package middleware

import "github.com/gin-gonic/gin"

const (
	AUTH     = "auth"
	RECOVERY = "recovery"
	LOGGER   = "logger"
	SECURE   = "secure"
	CORS     = "cors"
)

var Middlewares = map[string]func() gin.HandlerFunc{
	AUTH:     Auth,
	RECOVERY: Recovery,
	LOGGER:   Logger,
	SECURE:   Secure,
	CORS:     Cors,
}
