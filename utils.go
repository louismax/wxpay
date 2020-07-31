package wxpay

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io"
)

const (
	Fail       = "FAIL"
	Success    = "SUCCESS"
	HMACSHA256 = "HMAC-SHA256"
	MD5        = "MD5"
	Sign       = "sign"

	UnifiedOrderUrl     = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryUrl       = "https://api.mch.weixin.qq.com/pay/orderquery"
	CloseOrderUrl       = "https://api.mch.weixin.qq.com/pay/closeorder"
	RefundUrl           = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	RefundQueryUrl      = "https://api.mch.weixin.qq.com/pay/refundquery"
	DownloadBillUrl     = "https://api.mch.weixin.qq.com/pay/downloadbill"
	ReportUrl           = "https://api.mch.weixin.qq.com/payitil/report"
	MicroPayUrl         = "https://api.mch.weixin.qq.com/pay/micropay"
	ReverseUrl          = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	AuthCodeToOpenidUrl = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
	ShortUrl            = "https://api.mch.weixin.qq.com/tools/shorturl"

	//微信支付沙箱测试接口
	SandboxUnifiedOrderUrl     = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	SandboxOrderQueryUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
	SandboxCloseOrderUrl       = "https://api.mch.weixin.qq.com/sandboxnew/pay/closeorder"
	SandboxRefundUrl           = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
	SandboxRefundQueryUrl      = "https://api.mch.weixin.qq.com/sandboxnew/pay/refundquery"
	SandboxDownloadBillUrl     = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
	SandboxReportUrl           = "https://api.mch.weixin.qq.com/sandboxnew/payitil/report"
	SandboxMicroPayUrl         = "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay"
	SandboxReverseUrl          = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/reverse"
	SandboxAuthCodeToOpenidUrl = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"
	SandboxShortUrl            = "https://api.mch.weixin.qq.com/sandboxnew/tools/shorturl"
)

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//GetGUID 产生GUID
func GetGUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

// 将Pkcs12转成Pem
func pkcs12ToPem(p12 []byte, password string) tls.Certificate {

	blocks, err := pkcs12.ToPEM(p12, password)

	defer func() {
		if x := recover(); x != nil {
			fmt.Println("ERR:", x)
		}
	}()

	if err != nil {
		panic(err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}
