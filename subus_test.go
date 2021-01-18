package wxpay

import (
	"encoding/json"
	"testing"
)

func TestClient_SubmitMicroMch(t *testing.T) {
	apiaccount := NewAccount("wxddba5d27a2c4b9eb", "", "1501889641", "1509940251", "anzhijiaoyu1234567890anzhijiaoyu", false)
	apiaccount.SetCertData("apiclient_cert.p12")
	client := NewClient(apiaccount)

	//获取平台证书
	params := make(Params)
	params.SetString("mch_id", "1501889641")
	res,_ :=client.GetCertficates(params)
	result := SystemOauthTokenRsp{}
	err := json.Unmarshal([]byte(res.GetString("certificates")), &result)
	if err != nil {
		t.Log(err)
		return
	}
	//fmt.Println(res.GetString("certificates"))
	if len(result.Data)<1{
		t.Log("获取证书失败")
		return
	}

	apiv3key := "anzhixiaopay2019anzhixiaopay9876"
	//apiv3key := result.Data[0].Serial_no
	//证书解密
	cretinfoint, err := result.CertificateDecryption(apiv3key)
	if err != nil {
		t.Log(err)
		return
	}
	cretinfo := cretinfoint.(CertificateInfo)

	//申请内容
	params2 := make(Params)
	params2.SetString("version", "3.0").
		SetString("cert_sn",result.Data[0].Serial_no).
		SetString("mch_id","1501889641").
		SetString("business_code",GetGUID()).
		SetString("id_card_copy","tnRKR6S1W9YZeCwNUp7wYVWjdBtmRHgoc2-CeF-QgpnEhlpNP7GwGUtqKPINlmLRnyfuDtqs5dhCdCvFT7MtLTan3WnQ9TM7KYfjRmFiD6o").
		SetString("id_card_national","tnRKR6S1W9YZeCwNUp7wYVWjdBtmRHgoc2-CeF-QgpnEhlpNP7GwGUtqKPINlmLRnyfuDtqs5dhCdCvFT7MtLTan3WnQ9TM7KYfjRmFiD6o")

	jmname,err := SensitiveDataEncryption("张路平",cretinfo.Publickey)
	if err != nil{
		t.Log(err)
		return
	}
	jmcm,err:= SensitiveDataEncryption("53233119950401261X",cretinfo.Publickey)
	if err != nil{
		t.Log(err)
		return
	}
	jmphone,err:= SensitiveDataEncryption("13135116121",cretinfo.Publickey)
	if err != nil{
		t.Log(err)
		return
	}

	jmanm,err:= SensitiveDataEncryption("6214837317836322",cretinfo.Publickey)
	if err != nil{
		t.Log(err)
		return
	}

	jmemail,err:= SensitiveDataEncryption("louis8@163.com",cretinfo.Publickey)
	if err != nil{
		t.Log(err)
		return
	}

	params2.SetString("id_card_name",jmname).
		SetString("id_card_number",jmcm).
		SetString("id_card_valid_time","[\"1970-01-01\",\"长期\"]").
		SetString("account_name",jmname).
		SetString("account_bank","工商银行").
		SetString("bank_address_code","110000").
		SetString("account_number",jmanm).
		SetString("store_name","安智零食服务测试").
		SetString("store_address_code","110000").
		SetString("store_street","无").
		SetString("store_entrance_pic","tnRKR6S1W9YZeCwNUp7wYVWjdBtmRHgoc2-CeF-QgpnEhlpNP7GwGUtqKPINlmLRnyfuDtqs5dhCdCvFT7MtLTan3WnQ9TM7KYfjRmFiD6o").
		SetString("indoor_pic","tnRKR6S1W9YZeCwNUp7wYVWjdBtmRHgoc2-CeF-QgpnEhlpNP7GwGUtqKPINlmLRnyfuDtqs5dhCdCvFT7MtLTan3WnQ9TM7KYfjRmFiD6o").
		SetString("merchant_shortname","安智零食").
		SetString("service_phone","13135116121").
		SetString("product_desc","其他").
		SetString("rate","0.6%").
		SetString("contact",jmname).
		SetString("contact_phone",jmphone).
		SetString("contact_email",jmemail)

	res2,err :=client.SubmitMicroMch(params2)
if err != nil{
	t.Log(err)
	return
}

	//params.SetString("sign",client.Sign(params))
	t.Log(res2)
}
