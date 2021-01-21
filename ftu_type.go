package wxpay

const (
	MERCHANT_ID         = "MERCHANT_ID"         //商户号（mch_id或者sub_mch_id）
	PERSONAL_OPENID     = "PERSONAL_OPENID"     //个人openid（由父商户APPID转换得到）
	PERSONAL_SUB_OPENID = "PERSONAL_SUB_OPENID" //个人sub_openid（由子商户APPID转换得到）

	SERVICE_PROVIDER = "SERVICE_PROVIDER" //服务商
	STORE            = "STORE"            //门店
	STAFF            = "STAFF"            //员工
	STORE_OWNER      = "STORE_OWNER"      //店主
	PARTNER          = "PARTNER"          //合作伙伴
	HEADQUARTER      = "HEADQUARTER"      //总部
	BRAND            = "BRAND"            //品牌方
	DISTRIBUTOR      = "DISTRIBUTOR"      //分销商
	USER             = "USER"             //用户
	SUPPLIER         = "SUPPLIER"         //供应商
	CUSTOM           = "CUSTOM"           //自定义
)

type ReqProfitSharingAddReceiver struct {
	Type           string `json:"type,omitempty"`
	Account        string `json:"account,omitempty"`
	Name           string `json:"name,omitempty"`
	RelationType   string `json:"relation_type,omitempty"`
	CustomRelation string `json:"custom_relation,omitempty"`
}

type ReqProfitSharingRemoveReceiver struct {
	Type    string `json:"type,omitempty"`
	Account string `json:"account,omitempty"`
}

type ReqProfitSharing struct {
	Type        string `json:"type,omitempty"`
	Account     string `json:"account,omitempty"`
	Amount      int    `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}
