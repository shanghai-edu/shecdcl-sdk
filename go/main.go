package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shanghai-edu/shecdcl-sdk/go/shecgw"
)

//调用示例
func main() {
	/*
		初始化教委网关的连接信息
		政务外网：http://10.101.40.93
		城域网：https://apigw.shec.edu.cn
		申请调用账号和接口文档详见 —— http://dcl.shec.edu.cn/datacatalog/
	*/
	config := shecgw.Config{
		AppId:         "app_id",
		AppSecret:     "app_secert",
		ApiGwEndpoint: "https://apigw.shec.edu.cn", //默认 https://apigw.shec.edu.cn
		Debug:         true,                        //默认 false，true 时将打印接口原始输出
	}
	shecgw.Init(config)
	c := shecgw.GetShecClient()

	/*
		通过姓名+证件号码获取随申码数据
		建议将获取成功的数据缓存24小时
		如高频使用，建议在低峰期预热缓存
	*/

	ssmRes, err := c.GetSjSsmjm("冯小骐", "360732197709013718")
	if err != nil {
		log.Fatalln(err)
	}
	//raw output
	/*
		{"code":"0","message":"success","data":[{"type":"00"}]}
	*/
	bs, _ := json.MarshalIndent(ssmRes.Data, " ", " ")
	fmt.Println(string(bs))
	//sdk data output
	/*
		[
		  {
		   "type": "00"
		  }
		 ]
	*/

	/*
		随申码的扫码接口
		返回的证件号是脱敏的，需要和姓名一起结合匹配
		如果有此需求，注意给姓名所在字段添加索引，否则可能拖累数据库查询效率
	*/
	ssmjlRes, err := c.GetSjJkmjk("https://s.sh.gov.cn/5a055b864fcc853cf2bb2c6dbc55061642941877644")
	if err != nil {
		log.Fatalln(err)
	}
	//raw output
	/*
		{"code":"0","data":"{\"xm\":\"冯小骐\",\"phone\":\"189****9537\",\"type\":\"00\",\"zjhm\":\"360732********3718\",\"dzzz\":\"1\",\"uuid\":\"BBBBBBBBBBdAfcrLDtKovGRjA+l3ZkAAAAAAAAAAAAAA+kPXMsVD/9pkjBQ19geBU9yn5M=\"}","message":""}
	*/
	bs, _ = json.MarshalIndent(ssmjlRes.Data, " ", " ")
	fmt.Println(string(bs))
	//sdk data output
	/*
		{
		  "xm": "冯小骐",
		  "phone": "189****9537",
		  "type": "00",
		  "zjhm": "360732********3718",
		  "dzzz": "1",
		  "uuid": "BBBBBBBBBBdAfcrLDtKovGRjA+l3ZkAAAAAAAAAAAAAA+kPXMsVD/9pkjBQ19geBU9yn5M=\"
		}
	*/
	/*
		通过姓名+证件号码获取疫苗接种信息，核酸检测信息
		数据来源国家卫健委，通常比上海健康云要滞后一天
		如调用过于频繁可能触发国家卫建委的限流降级，建议将获取成功的数据至少缓存24小时
		如高频使用，建议在低峰期预热缓存
		注意返回内容在第二层 Data 内
	*/
	//疫苗接种信息调用示例
	ymRes, err := c.GetGjXgymjzxx("冯小骐", "360732197709013718")
	if err != nil {
		log.Fatalln(err)
		return
	}
	//raw output
	/*
		{"code":200,"message":"","data":"{\"code\":\"200\",\"data\":\"{\\\"gj\\\":\\\"156\\\",\\\"gazt\\\":\\\"01\\\",\\\"xm\\\":\\\"冯小骐\\\",\\\"zjlx\\\":\\\"01\\\",\\\"grdabh\\\":\\\"3607322210326186131\\\",\\\"jzxxlb\\\":\\\"[{\\\\\\\"ymmc\\\\\\\":\\\\\\\"5601\\\\\\\",\\\\\\\"jc\\\\\\\":\\\\\\\"1\\\\\\\",\\\\\\\"jzrq\\\\\\\":\\\\\\\"20210327T095032\\\\\\\",\\\\\\\"scqy\\\\\\\":\\\\\\\"02\\\\\\\",\\\\\\\"jzd\\\\\\\":\\\\\\\"31\\\\\\\"},{\\\\\\\"ymmc\\\\\\\":\\\\\\\"5601\\\\\\\",\\\\\\\"jc\\\\\\\":\\\\\\\"2\\\\\\\",\\\\\\\"jzrq\\\\\\\":\\\\\\\"20210417T094154\\\\\\\",\\\\\\\"scqy\\\\\\\":\\\\\\\"02\\\\\\\",\\\\\\\"jzd\\\\\\\":\\\\\\\"31\\\\\\\"},{\\\\\\\"ymmc\\\\\\\":\\\\\\\"5601\\\\\\\",\\\\\\\"jc\\\\\\\":\\\\\\\"3\\\\\\\",\\\\\\\"jzrq\\\\\\\":\\\\\\\"20211130T153012\\\\\\\",\\\\\\\"scqy\\\\\\\":\\\\\\\"14\\\\\\\",\\\\\\\"jzd\\\\\\\":\\\\\\\"31\\\\\\\"}]\\\",\\\"zjhm\\\":\\\"360732197709013718\\\"}\",\"message\":\"查询成功！\"}"}
	*/
	bs, _ = json.MarshalIndent(ymRes.Data.Data, " ", " ")
	fmt.Println(string(bs))
	// sdk data output
	/*
		{
		  "grdabh": "3607322210326186131",
		  "gazt": "01",
		  "xm": "冯小骐",
		  "gj": "156",
		  "zjlx": "01",
		  "zjhm": "360732197709013718",
		  "jzxxlb": [
		   {
		    "ymmc": "5601",
		    "jc": "1",
		    "jzrq": "20210327T095032",
		    "scqy": "02",
		    "jzd": "31"
		   },
		   {
		    "ymmc": "5601",
		    "jc": "2",
		    "jzrq": "20210417T094154",
		    "scqy": "02",
		    "jzd": "31"
		   },
		   {
		    "ymmc": "5601",
		    "jc": "3",
		    "jzrq": "20211130T153012",
		    "scqy": "14",
		    "jzd": "31"
		   }
		  ]
		 }
	*/

	//最近一次的核酸检测信息调用示例
	hsRes, err := c.GetGjXgfyHsjcsjfwjk("冯小骐", "360732197709013718")
	if err != nil {
		log.Fatalln(err)
		return
	}
	//raw output
	/*
		{"code":200,"message":"","data":"{\"code\":\"200\",\"data\":{\"xm\":\"冯小骐\",\"zjhm\":\"360732197709013718\",\"hsjcjg\":\"阴性\",\"hsjcsj\":\"2022-01-21 13:16:00\",\"hsjcjgmc\":\"苏州第九人民医院\"},\"message\":\"查询成功！\"}"}
	*/
	bs, _ = json.MarshalIndent(hsRes.Data.Data, " ", " ")
	fmt.Println(string(bs))
	//sdk data output
	/*
		{
		  "xm": "冯小骐",
		  "zjhm": "360732197709013718",
		  "hsjcjg": "阴性",
		  "hsjcsj": "2022-01-21 13:16:00",
		  "hsjcjgmc": "苏州第九人民医院"
		}
	*/

	//上海市核酸检测信息调用示例
	hsAllRes, err := c.GetSjHsjcjgcxSjgt("冯小骐", "360732197709013718")
	if err != nil {
		log.Fatalln(err)
		return
	}
	//raw output
	/*
		{"code":"0","message":"success","data":[{"sample_orgname":"上海第九人民医院","test_orgname":"上海第九人民医院","card_type":"1","sample_type":1,"card_type_name":"身份证","check_project":"核酸","sample_type_name":"鼻咽拭子","collect_date":"2022-02-03 17:41:45","card_no":"360732197709013718","name":"冯小骐","sample_date":"2022-02-03 08:46:33","nat_result":1,"nat_result_name":"ORF1a/b阴性，N基因阴性","report_date":"2022-02-03 17:41:45"},{"sample_orgname":"上海第五人民医院","test_orgname":"上海第五人民医院","card_type":"1","sample_type":1,"card_type_name":"身份证","check_project":"核酸","sample_type_name":"鼻咽拭子","collect_date":"2021-12-23 18:48:01","card_no":"360732197709013718","name":"冯小骐","sample_date":"2021-12-23 09:25:06","nat_result":1,"nat_result_name":"ORF1a/b阴性，N基因阴性","report_date":"2021-12-23 18:48:01"}]}
	*/
	bs, _ = json.MarshalIndent(hsAllRes.Data, " ", " ")
	fmt.Println(string(bs))
	//sdk data output
	/*
				[
		  {
		   "name": "冯小骐",
		   "card_no": "360732197709013718",
		   "card_type": "1",
		   "card_type_name": "身份证",
		   "sample_orgname": "上海第九人民医院",
		   "test_orgname": "上海第九人民医院",
		   "sample_type": 1,
		   "sample_type_name": "鼻咽拭子",
		   "check_project": "核酸",
		   "collect_date": "2022-02-03 17:41:45",
		   "sample_date": "2022-02-03 08:46:33",
		   "nat_result": 1,
		   "nat_result_name": "ORF1a/b阴性，N基因阴性",
		   "report_date": "2022-02-03 17:41:45"
		  },
		  {
		   "name": "冯小骐",
		   "card_no": "360732197709013718",
		   "card_type": "1",
		   "card_type_name": "身份证",
		   "sample_orgname": "上海第五人民医院",
		   "test_orgname": "上海第五人民医院",
		   "sample_type": 1,
		   "sample_type_name": "鼻咽拭子",
		   "check_project": "核酸",
		   "collect_date": "2021-12-23 18:48:01",
		   "sample_date": "2021-12-23 09:25:06",
		   "nat_result": 1,
		   "nat_result_name": "ORF1a/b阴性，N基因阴性",
		   "report_date": "2021-12-23 18:48:01"
		  }]
	*/
}
