package modules

import (
	"qqBot/bot"
	"sync"
)

// init.go 样例写法

// 注册 module 实例
func init() {
	bot.RegisterModule(instance)
}

var instance = &sample{}

type sample struct{}

func (s sample) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "",
		Instance: instance,
	}
}

func (s sample) Init() {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
}

func (s sample) PostInit() {
	// 第二次初始化
	// 再次过程中可以进行跨Module的动作
	// 如通用数据库等等
}

func (s sample) Serve(bot *bot.Bot) {
	// 注册服务函数部分
}

func (s sample) Start(bot *bot.Bot) {
	// 此函数会新开携程进行调用
	// ```go
	// 		go exampleModule.Start()
	// ```

	// 可以利用此部分进行后台操作
	// 如http服务器等等
}

func (s sample) Stop(bot *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
}
