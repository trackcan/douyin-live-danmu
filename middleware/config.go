package middleware

import (
	"douyin-live-danmusrv/config"
	"douyin-live-danmusrv/consts"
	"github.com/gin-gonic/gin"
)

func InitConfig(conf *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(consts.CONFIG, conf)
		c.Next()
	}
}
