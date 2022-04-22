package utils

import (
	"bytes"
	"encoding/json"
	"github.com/Logiase/MiraiGo-Template/utils"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 5 * time.Second}
var logger = utils.GetModuleLogger("logiase.autoreply")

// GetJsonToObject 把消息体中内容赋值给结构体数据
// data 传入的结构体
// url 请求的url
func GetJsonToObject(data interface{}, url string) {
	response, err := client.Get(url)
	if err != nil {
		logger.Warnf("GetJsonToObject client.get(%s) failed", url)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Warn("GetJsonToObject Body.Close() failed")
		}
	}(response.Body)
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := response.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			logger.Warn("GetJsonToObject get data failed")
			return
		}
	}

	err = json.Unmarshal(result.Bytes(), &data)
	if err != nil {
		return
	}
}
