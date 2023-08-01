package api

import (
	"douyin-live-danmusrv/consts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

func FailInvalidParam(c *gin.Context) {
	Fail(c, consts.ERR_INVALID_PARAM, "invalid param")
}
