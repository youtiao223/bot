package autoreply

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"

	"github.com/youtiao223/bot/bot"
	"github.com/youtiao223/bot/config"
	"github.com/youtiao223/bot/utils"
	"os"

	"sync"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &ar{}
var logger = utils.GetModuleLogger("logiase.autoreply")
var tem map[string]string

type ar struct {
}

var ArConfig *config.Config

func (a *ar) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "youtiao223.autoreply",
		Instance: instance,
	}
}

func (a *ar) Init() {
	// 从配置文件里读取自动回复消息内容
	path := config.GlobalConfig.GetString("logiase.autoreply.path")

	if path == "" {
		path = "./autoreply.yaml"
	}
	tem = make(map[string]string)
	initArConfig(path)
}

func (a *ar) PostInit() {
}

func (a *ar) Serve(b *bot.Bot) {
	b.OnGroupMessage(func(c *client.QQClient, msg *message.GroupMessage) {
		hasPrefix, data := checkPrefixAndGetData(msg.ToString())
		if !hasPrefix {
			return
		}
		out := autoreply(data)
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendGroupMessage(msg.GroupCode, m)
	})

	b.OnPrivateMessage(func(c *client.QQClient, msg *message.PrivateMessage) {
		out := autoreply(msg.ToString())
		if out == "" {
			return
		}
		m := message.NewSendingMessage().Append(message.NewText(out))
		c.SendPrivateMessage(msg.Sender.Uin, m)
	})
}

func (a *ar) Start(bot *bot.Bot) {
}

func (a *ar) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
}

func autoreply(in string) string {
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
