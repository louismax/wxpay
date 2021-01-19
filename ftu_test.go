package wxpay

import "testing"

func TestClient_ProfitSharingMerchantRatioQuery(t *testing.T) {
	client := NewClient(NewAccount("", "", "", "", "", false))

	t.Log(client.ProfitSharingMerchantRatioQuery())
}

func TestClient_ProfitSharingAddReceiver(t *testing.T) {
	receiver := ReqProfitSharingAddReceiver{
		Type:         MERCHANT_ID,
		Account:      "",
		Name:         "",
		RelationType: PARTNER,
	}
	client := NewClient(NewAccount("", "", "", "", "", false))
	t.Log(client.ProfitSharingAddReceiver(&receiver))
}

func TestClient_ProfitSharingRemoveReceiver(t *testing.T) {
	receiver := ReqProfitSharingRemoveReceiver{
		Type:    MERCHANT_ID,
		Account: "",
	}
	client := NewClient(NewAccount("", "", "", "", "", false))
	t.Log(client.ProfitSharingRemoveReceiver(&receiver))
}

func TestClient_ProfitSharingOrderAmountQuery(t *testing.T) {
	client := NewClient(NewAccount("", "", "", "", "", false))

	t.Log(client.ProfitSharingOrderAmountQuery(""))
}

func TestClient_ProfitSharing(t *testing.T) {
	var receivers []ReqProfitSharing
	receiver := ReqProfitSharing{
		Type:        "",
		Account:     "",
		Amount:      3,
		Description: "",
	}
	receivers = append(receivers, receiver)

	apiaccount := NewAccount("", "", "", "", "", false)

	apiaccount.SetCertData("")
	client := NewClient(apiaccount)

	t.Log(client.ProfitSharing("", GetGUID(), &receivers))
}

func TestClient_MultiProfitSharing(t *testing.T) {
	var receivers []ReqProfitSharing
	receiver := ReqProfitSharing{
		Type:        "",
		Account:     "",
		Amount:      2,
		Description: "",
	}
	receivers = append(receivers, receiver)

	apiaccount := NewAccount("", "", "", "", "", false)

	apiaccount.SetCertData("")
	client := NewClient(apiaccount)

	t.Log(client.MultiProfitSharing("", GetGUID(), &receivers))
}

func TestClient_ProfitSharingQuery(t *testing.T) {
	client := NewClient(NewAccount("", "", "", "", "", false))
	t.Log(client.ProfitSharingQuery("", ""))
}
