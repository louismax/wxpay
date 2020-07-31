package wxpay

import (
	"fmt"
	"io/ioutil"
)

type Account struct {
	appID     string
	subappID 	string
	mchID     string
	submchID  string
	apiKey    string
	certData  []byte
	isSandbox bool //沙箱环境
}

// 创建微信支付账号
func NewAccount(appID,subappID , mchID , submchid , apiKey string, isSandbox bool) *Account {
	return &Account{
		appID:     appID,
		subappID: subappID,
		mchID:     mchID,
		submchID:  submchid,
		apiKey:    apiKey,
		isSandbox: isSandbox,
	}
}

// 设置证书
func (a *Account) SetCertData(certPath string) {
	certData, err := ioutil.ReadFile(certPath)
	if err != nil {
		fmt.Println("读取证书失败,ERR:",err)
		return
	}
	a.certData = certData
}