package api

import (
	"douyin-live-danmusrv/consts"
	"douyin-live-danmusrv/room"
	"github.com/gin-gonic/gin"
)

type EntryRoomParam struct {
	LiveUrl string `json:"live_url" form:"live_url" binding:"required"`
	ReportUrl string `json:"report_url" form:"report_url" binding:"required"`
}

func OnEntryRoom(c *gin.Context) {
	param := EntryRoomParam{}
	if errA := c.ShouldBindJSON(&param); errA != nil {
		FailInvalidParam(c)
		return
	}

	roomMgr := c.MustGet(consts.ROOMMGR).(*room.RoomMgr)
	_, err := roomMgr.Add(param.LiveUrl, param.ReportUrl)
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
