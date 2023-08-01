package main

import (
	"douyin-live-danmusrv/api"
	"douyin-live-danmusrv/config"
	"douyin-live-danmusrv/middleware"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

var tomlFile string

func init() {
	flag.StringVar(&tomlFile, "conf", "docs/conf.toml", "toml config file")
}

func main() {
	flag.Parse()

	cfg, err := config.UnmarshalConfig(tomlFile)
	if err != nil {
		fmt.Println("UnmarshalConfig: err:", err)
		return
	}

	r := gin.Default()
	r.Use(gin.Recovery())

	v1 := r.Group("/api")
	v1.Use(middleware.InitRoomMgr())
	v1.Use(middleware.InitConfig(cfg))
	{
		v1.POST("/room/entry", api.OnEntryRoom) // 进入直播间
		v1.POST("/room/exit", api.OnExitRoom)   // 退出直播间
	}

	r.Run(fmt.Sprintf(":%v", cfg.Port))
}
