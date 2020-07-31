# wxpay
微信支付普通商户版、服务商版(WeChat Pay) SDK for Golang 

 ![Language](https://img.shields.io/badge/language-Go-orange.svg) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE.md)


wxpay兼容微信普通商户版、服务商版接口，目前提供以下方法：

| 方法名              | 说明          |
| ---------------- | ----------- |
| UnifiedOrder     | 统一下单        |
| OrderQuery       | 查询订单        |
| CloseOrder       | 关闭订单        |
| Refund           | 申请退款        |
| RefundQuery      | 查询退款        |
| DownloadBill     | 下载对账单      |
| Report           | 交易保障        |
| MicroPay         | 付款码支付      |
| Reverse          | 付款码撤销订单   |
| AuthCodeToOpenid | 付款码查询openid |
| ShortUrl         | Native支付转换短链接 |

* 支持微信支付沙箱调试；
* 默认使用MD5进行签名，支持配置HMAC-SHA256；
* 通过HTTPS请求得到返回数据后会对其做必要的处理（例如验证签名，签名错误则抛出异常）。
* 对于DownloadBill，无论是否成功都返回Map，且都含有`return_code`和`return_msg`。若成功，其中`return_code`为`SUCCESS`，另外`data`对应对账单数据。
* 普通商户版与服务商版区别只在于Sub_Appid、Sub_Machid参数是否参与签名

## 安装

```bash
$ go get github.com/louismax/wxpay

```

## 示例

```cgo
// 创建支付账户
account := wxpay.NewAccount("appid", "subappid", "mchid", "submchid", "apikey", false)

// 设置证书
account.SetCertData("Certification path")

// 新建微信支付客户端
client := wxpay.NewClient(account)

// 设置支付账户
client.SetAccount(account)

// 设置http请求超时时间
client.SetHttpConnectTimeoutMs(2000)

// 设置http读取信息流超时时间
client.SetHttpReadTimeoutMs(1000)

// 更改签名类型
client.SetSignType(wxpay.HMACSHA256)

```

```cgo
// 统一下单
params := make(wxpay.Params)
params.SetString("body", "test").
		SetString("out_trade_no", "123456789").
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://louismax.com/notify").
		SetString("trade_type", "JSAPI")
p, _ := client.UnifiedOrder(params)

// 订单查询
params := make(wxpay.Params)
params.SetString("out_trade_no", "123456789")
p, _ := client.OrderQuery(params)

// 退款
params := make(wxpay.Params)
params.SetString("out_trade_no", "123456789").
		SetString("out_refund_no", "1234567890").
		SetInt64("total_fee", 1).
		SetInt64("refund_fee", 1)
p, _ := client.Refund(params)

// 退款查询
params := make(wxpay.Params)
params.SetString("out_refund_no", "123456789")
p, _ := client.RefundQuery(params)

```


```cgo
// 签名
signStr := client.Sign(params)

// 校验签名
b := client.ValidSign(params)

```

```cgo
// xml解析
params := wxpay.XmlToMap(xmlStr)

// map封装xml请求参数
b := wxpay.MapToXml(params)

```



## License
MIT license

## 联系我
louis8@163.com

## 感谢您的支持与厚爱
![wechatpay](https://louisweb.oss-cn-shenzhen.aliyuncs.com/louispay.jpg) 

