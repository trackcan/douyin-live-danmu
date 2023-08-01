package api

import (
	"douyin-live-danmusrv/config"
	"douyin-live-danmusrv/consts"
	"douyin-live-danmusrv/room"
	"github.com/gin-gonic/gin"
)

type EntryRoomParam struct {
	LiveUrl string `json:"live_url" form:"live_url" binding:"required"`
}

func OnEntryRoom(c *gin.Context) {
	param := EntryRoomParam{}
	if errA := c.ShouldBindJSON(&param); errA != nil {
		FailInvalidParam(c)
		return
	}

	roomMgr := c.MustGet(consts.ROOMMGR).(*room.RoomMgr)
	cfg := c.MustGet(consts.CONFIG).(*config.Config)
	_, err := roomMgr.Add(param.LiveUrl, cfg.ReportApi)
	if err != nil {
		Fail(c, consts.FAIL, err.Error())
		return
	}

	Success(c, nil)
}

func OnExitRoom(c *gin.Context) {
	param := EntryRoomParam{}
	if errA := c.ShouldBindJSON(&param); errA != nil {
		FailInvalidParam(c)
		return
	}

	roomMgr := c.MustGet(consts.ROOMMGR).(*room.RoomMgr)
	roomMgr.Remove(param.LiveUrl)
	Success(c, nil)
}
