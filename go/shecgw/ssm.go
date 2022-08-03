package shecgw

import (
	"bytes"
	"fmt"

	//	"strings"

	"encoding/json"
	"errors"

	"github.com/shanghai-edu/shecdcl-sdk/go/common/request"
)

type SsmAPIResult struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    []SSM  `json:"data"`
	Err
}

type SSM struct {
	Type string `json:"type"`
}

type SsmReq struct {
	//	XM   string `json:"XM"`
	ZJHM string `json:"ZJHM"`
}

//市级-随申码解码
func (c *Client) GetSjSsmjm(zjhm string) (result SsmAPIResult, err error) {
	if err := c.GetAccessToken(); err != nil {
		return result, err
	}
	url := c.ApiGwEndPoint + "/gateway/interface-sj-ssmjm/getInfo"

	headers := map[string]string{
		"access_token":    c.Token.AccessToken,
		"authoritytype":   "2",
		"elementsVersion": "1.00",
		"Content-Type":    "application/json",
	}
	req := SsmReq{
		//	XM:   xm,
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

	if err := json.Unmarshal(res, &result); err != nil {
		return result, err
	}
	//如果说令牌过期了，那再来一次
	if result.Code == "GATEWAY0006" {
		return c.GetSjSsmjm(zjhm)
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
