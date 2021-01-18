package wxpay

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//SubmitMicroMch 小微商户申请
func (c *Client)SubmitMicroMch(params Params) (Params, error)  {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = SubmitMicroMchUrl
	}
	params["nonce_str"] = GetGUID()
	params["sign_type"] = HMACSHA256
	params["sign"] = c.SignV2(params,HMACSHA256)

	//请求参数组装
	fmt.Println("请求参数组装：", MapToXml(params))

	// 将pkcs12证书转成pem
	cert := pkcs12ToPem(c.account.certData, c.account.mchID)

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	transport := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	h := &http.Client{Transport: transport}

	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(params)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return c.processResponseXmlNoSign(string(res))
}