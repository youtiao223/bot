package main

import (
	"os"
	"os/signal"

	"github.com/youtiao223/bot/bot"
	"github.com/youtiao223/bot/config"
	"github.com/youtiao223/bot/utils"

	_ "github.com/youtiao223/bot/modules/logging"
)

func init() {
	utils.WriteLogToFS(utils.LogInfoLevel, utils.LogWithStack)
	config.Init()
}

func main() {
	// 快速初始化
	bot.Init()

	// 初始化 Modules
	bot.StartService()

	// 使用协议
	// 不同协议可能会有部分功能无法使用
	// 在登陆前切换协议
	bot.UseProtocol(bot.AndroidPhone)

	// 登录
	bot.Login()

	// 刷新好友列表，群列表
	bot.RefreshList()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	bot.Stop()
}
