package autoreply

import (
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"qqBot/bot"
	"qqBot/config"
	"sync"
)

func init() {
	bot.RegisterModule(instance)
}

var instance = &ar{}

type ar struct{}

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
