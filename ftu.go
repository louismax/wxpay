package wxpay

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ProfitSharingMerchantRatioQuery 查询分账比例
func (c *Client) ProfitSharingMerchantRatioQuery() (Params, error) {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingMerchantRatioQueryUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}
	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	fmt.Println("请求参数组装：", params)
	h := &http.Client{}
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(params)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(string(res))
}

// ProfitSharingAddReceiver 添加分账接收方
func (c *Client) ProfitSharingAddReceiver(receiver *ReqProfitSharingAddReceiver) (Params, error) {
	if receiver.Type != MERCHANT_ID && receiver.Type != PERSONAL_OPENID && receiver.Type != PERSONAL_SUB_OPENID {
		return nil, errors.New("分账接收方类型无效")
	}
	if receiver.Account == "" {
		return nil, errors.New("分账接收方账号无效")
	}
	if receiver.Type == MERCHANT_ID {
		if receiver.Name == "" {
			return nil, errors.New("分账接收方全称不能为空")
		}
	}
	if receiver.RelationType != SERVICE_PROVIDER && receiver.RelationType != STORE && receiver.RelationType != STAFF && receiver.RelationType != STORE_OWNER && receiver.RelationType != PARTNER && receiver.RelationType != HEADQUARTER && receiver.RelationType != BRAND && receiver.RelationType != DISTRIBUTOR && receiver.RelationType != USER && receiver.RelationType != SUPPLIER && receiver.RelationType != CUSTOM {
		return nil, errors.New("关系类型无效")
	}
	if receiver.RelationType == CUSTOM {
		if receiver.CustomRelation == "" {
			return nil, errors.New("自定义的分账关系不能为空")
		}
	}

	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingAddReceiverUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["appid"] = c.account.appID
	if c.account.subappID != "" {
		params["sub_appid"] = c.account.subappID
	}
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}

	jsons, errs := json.Marshal(receiver)
	if errs != nil {
		return nil, errs
	}
	params["receiver"] = string(jsons)

	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	fmt.Println("请求参数组装：", params)
	h := &http.Client{}
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(params)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(string(res))
}

// ProfitSharingRemoveReceiver 删除分账接收方
func (c *Client) ProfitSharingRemoveReceiver(receiver *ReqProfitSharingRemoveReceiver) (Params, error) {
	if receiver.Type != MERCHANT_ID && receiver.Type != PERSONAL_OPENID && receiver.Type != PERSONAL_SUB_OPENID {
		return nil, errors.New("分账接收方类型无效")
	}
	if receiver.Account == "" {
		return nil, errors.New("分账接收方账号无效")
	}

	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingRemoveReceiverUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["appid"] = c.account.appID
	if c.account.subappID != "" {
		params["sub_appid"] = c.account.subappID
	}
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}

	jsons, errs := json.Marshal(receiver)
	if errs != nil {
		return nil, errs
	}
	params["receiver"] = string(jsons)

	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	fmt.Println("请求参数组装：", params)
	h := &http.Client{}
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(params)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(string(res))
}

// ProfitSharingOrderAmountQuery 查询订单待分账金额
func (c *Client) ProfitSharingOrderAmountQuery(transaction_id string) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingOrderAmountQueryUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["mch_id"] = c.account.mchID
	params["transaction_id"] = transaction_id
	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	fmt.Println("请求参数组装：", params)
	h := &http.Client{}
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(params)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(string(res))
}

// ProfitSharing 请求单次分账
func (c *Client) ProfitSharing(transaction_id, out_order_no string, receivers *[]ReqProfitSharing) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["appid"] = c.account.appID
	if c.account.subappID != "" {
		params["sub_appid"] = c.account.subappID
	}
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}
	params["out_order_no"] = out_order_no
	params["transaction_id"] = transaction_id

	jsons, errs := json.Marshal(receivers)
	if errs != nil {
		return nil, errs
	}
	params["receivers"] = string(jsons)

	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	xmlStr, err := c.postWithCert(url, params, true)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(xmlStr)
}

// MultiProfitSharing 请求多次分账
func (c *Client) MultiProfitSharing(transaction_id, out_order_no string, receivers *[]ReqProfitSharing) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = MultiProfitSharingUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["appid"] = c.account.appID
	if c.account.subappID != "" {
		params["sub_appid"] = c.account.subappID
	}
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}
	params["out_order_no"] = out_order_no
	params["transaction_id"] = transaction_id

	jsons, errs := json.Marshal(receivers)
	if errs != nil {
		return nil, errs
	}
	params["receivers"] = string(jsons)

	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	xmlStr, err := c.postWithCert(url, params, true)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(xmlStr)
}

// ProfitSharingQuery 查询分账结果
func (c *Client) ProfitSharingQuery(transaction_id, out_order_no string) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingQueryUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}
	params["transaction_id"] = transaction_id
	params["out_order_no"] = out_order_no
	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	fmt.Println("请求参数组装：", params)
	h := &http.Client{}
	response, err := h.Post(url, bodyType, strings.NewReader(MapToXml(params)))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(string(res))
}

// ProfitSharingFinish 完结分账
func (c *Client) ProfitSharingFinish(transaction_id, out_order_no, description string) (Params, error) {
	var url string
	if c.account.isSandbox {
		url = ""
	} else {
		url = ProfitSharingFinishUrl
	}
	c.signType = HMACSHA256
	params := Params{}
	params["appid"] = c.account.appID
	params["mch_id"] = c.account.mchID
	if c.account.submchID != "" {
		params["sub_mch_id"] = c.account.submchID
	}
	params["out_order_no"] = out_order_no
	params["transaction_id"] = transaction_id
	params["description"] = description

	params["nonce_str"] = GetGUID()
	params["sign_type"] = c.signType
	params["sign"] = c.Sign(params)

	xmlStr, err := c.postWithCert(url, params, true)
	if err != nil {
		return nil, err
	}
	return c.processResponseXml(xmlStr)
}
