package shecgw

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shanghai-edu/shecdcl-sdk/go/common/request"
)

// ResHscj 核酸检测信息最外层
type ResHscjSh struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    []HscjShData `json:"data"`
	Err
}

type ReqXmSfzh struct {
	XM   string `json:"name"` //姓名
	ZJHM string `json:"sfzh"` //证件号码
}

// Hscj 核酸检测信息
type HscjShData struct {
	XM             string `json:"name"`             //姓名
	ZJHM           string `json:"card_no"`          //证件号码
	CARDTYPE       string `json:"card_type"`        //证件类型
	CARDTYPENAME   string `json:"card_type_name"`   //证件类型名称
	SAMPLEORG      string `json:"sample_orgname"`   //采样机构
	TESTORG        string `json:"test_orgname"`     //检测机构
	SAMPLETYPE     string `json:"sample_type"`      //样本类型
	SAMPLETYPENAME string `json:"sample_type_name"` //样本类型名称
	CHECKPROJECT   string `json:"check_project"`    //核酸检测类型：核酸/核酸抗体
	COLLECTDATE    string `json:"collect_date"`     //收样时间
	SAMPLEDATE     string `json:"sample_date"`      //采样时间
	NATRESULT      string `json:"nat_result"`       //检测结果
	NATRESULTNAME  string `json:"nat_result_name"`  //检测结果名称
	REPORTDATE     string `json:"report_date"`      //检测结果报告时间
}

// GetSjHsjcjgcxSjgt 获取核酸检测信息
//市级-核酸检测结果查询-数据高铁接口
func (c *Client) GetSjHsjcjgcxSjgt(xm, zjhm string) (result ResHscjSh, err error) {
	if err := c.GetAccessToken(); err != nil {
		return result, err
	}
	url := c.ApiGwEndPoint + "/gateway/interface-sj-hsjcjgcx-sjgt/getInfo"

	headers := map[string]string{
		"access_token":    c.Token.AccessToken,
		"authoritytype":   "2",
		"elementsVersion": "1.00",
		"Content-Type":    "application/json",
	}
	req := ReqXmSfzh{
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
		return c.GetSjHsjcjgcxSjgt(xm, zjhm)
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
