package lib

import (
	"fmt"
	_ "golang.org/x/mobile/bind"
	"qqBot/bot"
	"qqBot/config"
)

func init() {
	fmt.Println("bot as lib")
}

func InitBot(configJSONContent string, deviceJSONContent string) {
	config.InitWithContent([]byte(configJSONContent))
	bot.InitWithDeviceJSONContent([]byte(deviceJSONContent))
	bot.StartService()
	bot.UseProtocol(bot.IPad)
	bot.Login()
	bot.RefreshList()
}
