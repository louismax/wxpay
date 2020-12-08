package wxpay

import "testing"

func TestClient_UnifiedOrder(t *testing.T) {
	client := NewClient(NewAccount("", "", "", "", "", false))
	params := make(Params)
	params.SetString("body", "test").
		SetString("out_trade_no", "58867657575757").
		SetInt64("total_fee", 1).
		SetString("spbill_create_ip", "127.0.0.1").
		SetString("notify_url", "http://louismax.com/notify").
		SetString("trade_type", "JSAPI")
	t.Log(client.UnifiedOrder(params))
}

func TestClient_OrderQuery(t *testing.T) {
	client := NewClient(NewAccount("wx6a7exxxxxxxxxxxx", "", "1500000000", "15122222222", "xxxxxxxxxxxxxxxxxxxxxxxxx", false))
	params := make(Params)
	//params.SetString("transaction_id", "4200000728202011181565873981")
	params.SetString("out_trade_no", "203231414102946822671332")
	t.Log(client.OrderQuery(params))
}

func TestClient_DownloadBill(t *testing.T) {
	client := NewClient(NewAccount("wx6a7exxxxxxxxxxxx", "", "1500000000", "15122222222", "xxxxxxxxxxxxxxxxxxxxxxxxx", false))
	params := make(Params)
	//params.SetString("transaction_id", "4200000526202005146172457219")
	params.SetString("bill_date", "20200829")
	params.SetString("bill_type", "ALL")
	t.Log(client.DownloadBill(params))
}

func TestClient_Sendminiprogramhb(t *testing.T) {
	apiaccount := NewAccount("wx6a7exxxxxxxxxxxx", "", "1500000000", "15122222222", "xxxxxxxxxxxxxxxxxxxxxxxxx", false)
	apiaccount.SetCertData("apiclient_cert.p12")
	client := NewClient(apiaccount)
	params := make(Params)
	params.SetString("mch_billno", GenValidateCode(28)).
		SetString("mch_id","1500000000").
		SetString("sub_mch_id","15122222222").
		SetString("wxappid","wx9cxxxxxxxxxxxxxxxxx").
		SetString("msgappid","wx9cxxxxxxxxxxxxxxxxx").
		SetString("send_name","测试百货").
		SetString("re_openid","owEr25blwvxEn3Xk5yuhieleMy9I").
		SetInt("total_amount",1).
		SetInt("total_num",1).
		SetString("wishing","测试111").
		SetString("act_name","测试222").
		SetString("remark","备注备注").
		SetString("notify_way","MINI_PROGRAM_JSAPI").
		SetString("nonce_str",GetGUID())

	params.SetString("sign",client.Sign(params))



	t.Log(client.Sendminiprogramhb(params))
}

func TestClient_Sendredpack(t *testing.T) {
	apiaccount := NewAccount("wx6a7exxxxxxxxxxxx", "", "1500000000", "15122222222", "xxxxxxxxxxxxxxxxxxxxxxxxx", false)
	apiaccount.SetCertData("apiclient_cert.p12")
	client := NewClient(apiaccount)
	params := make(Params)
	params.SetString("mch_billno", GenValidateCode(28)).
		SetString("mch_id","1500000000").
		SetString("sub_mch_id","15122222222").
		SetString("wxappid","wx9cxxxxxxxxxxxxxxxxx").
		SetString("msgappid","wx9cxxxxxxxxxxxxxxxxx").
		SetString("send_name","test1").
		SetString("re_openid","ovLnMvs1AaeG3YOBTjSQx7l_frcA").
		SetInt("total_amount",100).
		SetInt("total_num",1).
		SetString("wishing","测试111").
		SetString("client_ip","127.0.0.1").
		SetString("act_name","测试222").
		SetString("remark","备注备注").
		SetString("nonce_str",GetGUID())

	params.SetString("sign",client.Sign(params))
	t.Log(client.Sendredpack(params))
}

func TestClient_Gethbinfo(t *testing.T) {
	apiaccount := NewAccount("wx6a7exxxxxxxxxxxx", "", "1500000000", "15122222222", "xxxxxxxxxxxxxxxxxxxxxxxxx", false)
	apiaccount.SetCertData("apiclient_cert.p12")
	client := NewClient(apiaccount)
	params := make(Params)
	params.SetString("mch_billno", "0239670536479219974099409931").
		SetString("mch_id","1500000000").
		SetString("appid","wx9cxxxxxxxxxxxxxxxxx").
		SetString("bill_type","MCHT").
		SetString("nonce_str",GetGUID())

	params.SetString("sign",client.Sign(params))
	t.Log(client.Gethbinfo(params))
}