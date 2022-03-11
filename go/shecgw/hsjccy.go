package shecgw

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shanghai-edu/shecdcl-sdk/go/common/request"
)

// ResHscjCy 核酸检测信息最外层
type ResHscjCy struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    []HscjCyData `json:"data"`
	Err
}

type ReqCardNo struct {
	ZJHM string `json:"card_no"` //证件号码
}

// HscjCy 核酸检测采样信息
type HscjCyData struct {
	XM            string `json:"real_name"`           //姓名
	ZJHM          string `json:"card_no"`             //证件号码
	SAMPLEORG     string `json:"sample_org_name"`     //采样机构名称
	SAMPLESTATION string `json:"sample_station_name"` //采样点名称
	PRINTTIME     string `json:"print_time"`          //采样时间
}

// GetSjHsjcjgcxSjgt 获取核酸检测信息
//市大数据中心-核酸采样信息查询接口
func (c *Client) GetSjHscyxxcx(zjhm string) (result ResHscjCy, err error) {
	if err := c.GetAccessToken(); err != nil {
		return result, err
	}
	url := c.ApiGwEndPoint + "/gateway/interface-sj-hscyxxcx/getInfo"

	headers := map[string]string{
		"access_token":    c.Token.AccessToken,
		"authoritytype":   "2",
		"elementsVersion": "1.00",
		"Content-Type":    "application/json",
	}
	req := ReqCardNo{
		ZJHM: zjhm,
	}

	js, _ := json.Marshal(req)
	res, err := request.HTTPPost(url, bytes.NewBuffer(js), headers)
	if err != nil {
		return
	}
	//debug模式下打印原始输出
	if c.Debug {
		fmt.Println(string(res))
	}

	if err := json.Unmarshal(removeEscapes(res), &result); err != nil {
		return result, err
	}

	//如果说令牌过期了，那再来一次
	if result.ErrCode == "GATEWAY0006" {
		return c.GetSjHscyxxcx(zjhm)
	}

	if result.Code != "0" {
		err = errors.New(result.Message)
		return
	}

	if len(result.Data) == 0 {
		err = errors.New("市教委接口没有返回数据")
		return
	}

	return
}
