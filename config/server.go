package config

type Server struct {
	Mode           string   `mapstructure:"mode"`
	Health         string   `mapstructure:"health"`
	Middlewares    []string `mapstructure:"middlewares"`
	Host           string   `mapstructure:"host"`
	Port           string   `mapstructure:"port"`
	UseRedis       bool     `mapstructure:"useRedis"`
	TrustedProxies []string `mapstructure:"trustedProxies"`
}
