package main

import (
	log "github.com/sirupsen/logrus"
	"go-chatgpt-api/cache"
	"go-chatgpt-api/config"
	"go-chatgpt-api/initialize/banner"
	"go-chatgpt-api/initialize/mysql"
	"go-chatgpt-api/initialize/web"
)

func main() {

	banner.Init()
	err := config.LoadConfig()
	if err != nil {
		log.Errorf("没有找到配置文件，尝试读取环境变量")
		panic("请检查配置文件")
	}
	cache.Init()
	mysql.Init()
	web.Init()
}
