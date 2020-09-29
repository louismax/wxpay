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
	client := NewClient(NewAccount("wx6a7e7323e50a2657", "", "1501889641", "1508876741", "anzhijiaoyu1234567890anzhijiaoyu", true))
	params := make(Params)
	params.SetString("transaction_id", "4200000689202009287876193655")
	//params.SetString("out_trade_no", "202420487769097149658785")
	t.Log(client.OrderQuery(params))
}

func TestClient_DownloadBill(t *testing.T) {
	client := NewClient(NewAccount("wx6a7e7323e50a2657", "", "1501889641", "1508876741", "anzhijiaoyu1234567890anzhijiaoyu", false))
	params := make(Params)
	//params.SetString("transaction_id", "4200000526202005146172457219")
	params.SetString("bill_date", "20200829")
	params.SetString("bill_type", "ALL")
	t.Log(client.DownloadBill(params))
}
