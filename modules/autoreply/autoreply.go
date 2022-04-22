package autoreply

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"qqBot/config"
	"qqBot/utils"
	"strings"
)

var logger = utils.GetModuleLogger("logiase.autoreply")
var tem map[string]string
var ArConfig *config.Config

// 查询自动回复语句
func autoreply(in string) string {
	// 先转发模块
	if in == "来个笑话" {
		return getJoke()
	}
	// 模块中没有再从配置文件中返回
	out, ok := tem[in]
	if !ok {
		return "不好意思,我听不懂"
	}
	return out
}

// 读取 ar 的配置文件
func initArConfig(path string) {
	ArConfig = &config.Config{
		Viper: viper.New(),
	}

	configFile, _ := os.Open(path)

	ArConfig.SetConfigFile(path)
	err := ArConfig.ReadConfig(configFile)

	if err != nil {
		logrus.WithField("config", "ArConfig").WithError(err).Panicf("unable to read autoreply config")
	}

	for _, s := range ArConfig.AllKeys() {
		tem[s] = ArConfig.GetString(s)
	}

	// 热更新配置文件
	ArConfig.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
		for _, s := range ArConfig.AllKeys() {
			tem[s] = ArConfig.GetString(s)
		}
	})
	ArConfig.WatchConfig()
}

// 检查前缀若合法就从中提取消息
func checkPrefixAndGetData(msg string) (bool, string) {
	name := config.GlobalConfig.GetString("bot.name")
	var prefix = "@" + name

	hasPrefix := strings.HasPrefix(msg, prefix)

	var data = ""
	if hasPrefix {
		data = msg[len(prefix):]
	}
	return hasPrefix, strings.TrimSpace(data)
}
