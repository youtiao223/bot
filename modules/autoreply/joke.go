package autoreply

import (
	"qqBot/config"
	"qqBot/utils"
)

type Response struct {
	StatusCode string
	Desc       string
	Result     []Joke
}
type Joke struct {
	Id         int
	Content    string
	UpdateTime string
}

// 从公共API获取笑话
//test := []byte(`{
//    "statusCode":"000000",
//    "desc":"请求成功",
//    "result":[
//        {
//            "id":1,
//            "content":"姐夫外地出差，大姐带着小外甥回家，老妈整天喊大姐小棉袄。 住了几天，姐夫出差回来接这娘俩，小外甥去开门，老妈问谁呀，小外甥说：来取棉袄的！！",
//            "updateTime":"2017-04-07 20:53:50"
//        },
//        {
//            "id":2,
//            "content":"妈妈：“别人给你东西吃，你该怎么说？” 儿子：“还有吗？”",
//            "updateTime":"2017-04-07 20:53:50"
//        }
//    ]
//}`)

func getJoke() string {
	var res Response
	url := config.GlobalConfig.GetString("autoreply.joke.url")
	utils.GetJsonToObject(&res, url)
	jokeContent := res.Result[0].Content
	//jokeContent := "这是一个笑话"
	return jokeContent
}
