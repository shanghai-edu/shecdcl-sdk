package shecgw

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shanghai-edu/shecdcl-sdk/go/common/request"
)

type SsmsmRawResult struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
	Err
}

type SsmsmAPIResult struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    SsmsmData `json:"data"`
	Err
}

type SsmsmData struct {
	XM    string `json:"xm"`
	Phone string `json:"phone"`
	Type  string `json:"type"`
	ZJHM  string `json:"zjhm"`
	DZZZ  string `json:"dzzz"`
	UUID  string `json:"uuid"`
}

type SsmsmReq struct {
	Data string `json:"data"`
}

//市级-健康码接口（随申码被扫）
func (c *Client) GetSjJkmjk(urldata string) (result SsmsmAPIResult, err error) {
	if err := c.GetAccessToken(); err != nil {
		return result, err
	}
	url := c.ApiGwEndPoint + "/gateway/interface-sj-jkmjk/getInfo"

	headers := map[string]string{
		"access_token":    c.Token.AccessToken,
		"authoritytype":   "2",
		"elementsVersion": "1.00",
		"Content-Type":    "application/json",
	}
	req := SsmsmReq{
		Data: urldata,
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
	//似乎在传入非url参数时，会返回 null
	if string(res) == "null" {
		err = errors.New("市教委接口没有返回数据")
		return
	}

	//先不解析data层的内容，避免直接抛错
	var rawResult SsmsmRawResult
	if err := json.Unmarshal(res, &rawResult); err != nil {
		return result, err
	}

	//如果说令牌过期了，那再来一次
	if rawResult.ErrCode == "GATEWAY0006" {
		return c.GetSjJkmjk(urldata)
	}

	if rawResult.Code != "0" {
		err = errors.New(rawResult.Message)
		return
	}
	//说明正常返回了，那么拆解 data 里的转义符，重新解码到标准结构
	if err := json.Unmarshal(removeEscapes(res), &result); err != nil {
		return result, err
	}

	return
}
