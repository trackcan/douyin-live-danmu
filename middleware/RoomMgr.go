package middleware

import (
	"douyin-live-danmusrv/consts"
	"douyin-live-danmusrv/room"
	"github.com/gin-gonic/gin"
)

func InitRoomMgr() gin.HandlerFunc {
	mgr := room.InitRoomMgr()
	return func(c *gin.Context) {
		c.Set(consts.ROOMMGR, mgr)
		c.Next()
	}
}
