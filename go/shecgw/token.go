package shecgw

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/shanghai-edu/shecdcl-sdk/go/common/request"
)

// 目前教委反馈 token 24 小时过期
const EXPIRE = 86400

// Err 教委数据中心返回的相关状态信息
type Err struct {
	ErrCode   string `json:"code"`
	ErrStatus int    `json:"status"`
	ErrMsg    string `json:"msg"`
}

// AccessToken 教委数据中心请求Token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	Err
	ExpiresInTime time.Time
}

// Client 微信企业号应用配置信息
type Client struct {
	AppID         string
	AppSecret     string
	ApiGwEndPoint string
	Debug         bool
	Token         AccessToken
}

// GetAccessToken 获取会话token
func (c *Client) GetAccessToken() error {
	var err error

	if c.Token.AccessToken == "" || c.Token.ExpiresInTime.Before(time.Now()) {
		c.Token, err = getAccessToken(c.AppID, c.AppSecret, c.ApiGwEndPoint)
		if err != nil {
			return fmt.Errorf("invoke getAccessToken fail: %v", err)
		}
		//实际给的有效期为官方有效期的一半，也就是有效期不足一半时续新的token
		c.Token.ExpiresInTime = time.Now().Add(time.Duration(EXPIRE/2) * time.Second)
	}

	return err
}

// getAccessToken 获取token
func getAccessToken(appID, appSecret, apiGwEndpoint string) (accessToken AccessToken, err error) {
	url := apiGwEndpoint + "/gateway/auth/accesstoken/create?appId=" + appID + "&appSecret=" + appSecret

	var res []byte
	res, err = request.HTTPGet(url, nil, nil)
	if err != nil {
		return accessToken, err
	}

	err = json.Unmarshal(res, &accessToken)
	if err != nil {
		return accessToken, fmt.Errorf("parse gettoken response body fail: %v", err)
	}

	if accessToken.AccessToken == "" {
		err = fmt.Errorf("invoke api gettoken fail, ErrCode: %v, ErrMsg: %v", accessToken.ErrCode, accessToken.ErrMsg)
		return accessToken, err
	}

	return accessToken, err
}
