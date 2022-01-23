# 上海教育数据资源目录-SDK

数据资源目录详见 —— http://dcl.shec.edu.cn/datacatalog/

资源目录需要上海教育城域网内的教育用户才可访问

## 约定

`SDK` 不限定开发语言，数据结构和调用方式，但不同语言所开发的 `SDK` 应遵循以下约定
### 函数名

函数命名为对应接口路径中，有特殊含义部分的驼峰改写。

例如随申码的接口路径为 `/gateway/interface-sj-ssmjm/getInfo`，,其中代表接口含义的部分是 `sj-ssmjm`, 则 `SDK` 中的函数为 `GetSjSsmjm()`

以此类推， 疫苗接口的路径为 `/gateway/interface-gj-xgymjzxx/getInfo`，则 `SDK` 中的函数为 `GetGjXgymjzxx()`

### 数据结构

`SDK` 将接口返回的数据映射为对应的结构体或字典类型，但不对接口数据做额外处理，例如疫苗等涉及字典代码部分，需要应用根据代码表再自行进行映射。以尽可能的减少耦合，为应用提供尽可能多的灵活性。

但对于原接口以转义符方式返回的字符串类型的`json`数据结构，则将其解码为标准结构，以简化调用。

例如随申码的被扫接口，其原始接口返回内容为：
```
{"code":"0","data":"{\"xm\":\"冯小骐\",\"phone\":\"189****9537\",\"type\":\"00\",\"zjhm\":\"360732********3718\",\"dzzz\":\"1\",\"uuid\":\"BBBBBBBBBBdAfcrLDtKovGRjA+l3ZkAAAAAAAAAAAAAA+kPXMsVD/9pkjBQ19geBU9yn5M=\"}","message":""}
```

则由 `SDK` 返回的数据结构所生成的 `json` 应该是如下结构（或者是兼容这个结构的扩展）：
```
{
	"code": "0",
	"data": {
		"xm": "冯小骐",
		"phone": "189****9537",
		"type": "00",
		"zjhm": "360732********3718",
		"dzzz": "1",
		"uuid": "BBBBBBBBBBdAfcrLDtKovGRjA+l3ZkAAAAAAAAAAAAAA+kPXMsVD/9pkjBQ19geBU9yn5M="
	},
	"message": ""
}
```

## 目前支持的语言

- [Go](https://github.com/shanghai-edu/shecdcl-sdk/blob/main/go/README.MD)

## 参与贡献

- [freedomkk-qfeng](https://github.com/freedomkk-qfeng)

## 说明

SDK 由上海教育社区自发开源，不代表官方立场。

欢迎大家贡献 PR，共建开放生态。
