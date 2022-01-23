package shecgw

import (
	"sync"
)

var (
	client *Client
	lock   = new(sync.RWMutex)
)

type Config struct {
	AppId         string
	AppSecret     string
	ApiGwEndpoint string
	Debug         bool
}

// Init 初始化教委账号
func Init(config Config) {
	c := new(Client)
	c.AppID = config.AppId
	c.AppSecret = config.AppSecret
	//默认用公网地址
	if config.ApiGwEndpoint == "" {
		c.ApiGwEndPoint = "https://apigw.shec.edu.cn"
	} else {
		c.ApiGwEndPoint = config.ApiGwEndpoint
	}
	c.Debug = config.Debug
	client = c
}

//GetShecClient 获取教委接口客户端的配置信息
func GetShecClient() *Client {
	lock.RLock()
	defer lock.RUnlock()
	return client
}
