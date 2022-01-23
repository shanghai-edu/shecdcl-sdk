package shecgw

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/shanghai-edu/shecdcl-sdk/go/common/request"
)

// ResHscj 核酸检测信息最外层
type ResHscj struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    HscjData `json:"data"`
	Err
}

// HscjData 核酸检测信息第一层
type HscjData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    Hscj   `json:"data"`
	WsjkError
}

// Hscj 核酸检测信息
type Hscj struct {
	XM       string `json:"xm"`       //姓名
	ZJHM     string `json:"zjhm"`     //证件号码
	HSJCJG   string `json:"hsjcjg"`   //核酸检测结果
	HSJCSJ   string `json:"hsjcsj"`   //核酸检测时间
	HSJCJGMC string `json:"hsjcjgmc"` //核酸检测机构名称
}

//卫生监控条线省级降级的错误信息
type WsjkError struct {
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

// GetHscj 获取核酸检测信息
//国家-新冠防疫-核酸检测市级防疫接口
func (c *Client) GetGjXgfyHsjcsjfwjk(xm, zjhm string) (result ResHscj, err error) {
	if err := c.GetAccessToken(); err != nil {
		return result, err
	}
	url := c.ApiGwEndPoint + "/gateway/interface-gj-xgfy-hsjcsjfwjk/getInfo"

	headers := map[string]string{
		"access_token":    c.Token.AccessToken,
		"authoritytype":   "2",
		"elementsVersion": "1.00",
		"Content-Type":    "application/json",
	}
	req := ReqXmZjhm{
		XM:   xm,
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
		return c.GetGjXgfyHsjcsjfwjk(xm, zjhm)
	}

	if result.Code != 200 {
		err = errors.New(result.Message)
		return
	}

	if result.Data.ErrCode != "" {
		err = errors.New(result.Data.ErrMsg)
		return
	}

	if result.Data.Code != "200" {
		err = errors.New(result.Data.Message)
		return
	}
	if result.Data.Data.ZJHM == "" {
		err = errors.New("市教委没有查到数据: " + result.Data.Message)
		return
	}
	if strings.ToUpper(result.Data.Data.ZJHM) != zjhm {
		err = errors.New("市教委返回的数据与查询参数不匹配: " + result.Data.Data.ZJHM)
		return
	}

	return
}
