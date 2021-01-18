package wxpay

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io"
	"log"
	"os"
	"strings"
	"time"
	mathrand "math/rand"
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
	Sendminiprogramhb = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendminiprogramhb"
	SendredpackUrl = "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack"
	GethbinfoUrl = "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo"

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

	GetsignkeyUrl ="https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey"

	//冷门接口
	GetCertficatesUrl= "https://api.mch.weixin.qq.com/risk/getcertficates"
	UploadMediaUrl ="https://api.mch.weixin.qq.com/secapi/mch/uploadmedia"

	//小微商户
	SubmitMicroMchUrl = "https://api.mch.weixin.qq.com/applyment/micro/submit"
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

//GenValidateCode 生成指定长度数字
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	mathrand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[mathrand.Intn(r)])
	}
	return sb.String()
}


type SystemOauthTokenRsp struct {
	Data []struct {
		Serial_no           string `json:"serial_no"`
		Effective_time      string `json:"effective_time"`
		Expire_time         string `json:"expire_time"`
		Encrypt_certificate struct {
			Algorithm       string `json:"algorithm"`
			Nonce           string `json:"nonce"`
			Associated_data string `json:"associated_data"`
			Ciphertext      string `json:"ciphertext"`
		} `json:"encrypt_certificate"`
	} `json:"data"`
}

func (this SystemOauthTokenRsp) CertificateDecryption(apiv3key string) (interface{}, error) {
	if len(this.Data) < 1 {
		return nil, errors.New("没有找到平台密文证书")
	} else {
		cpinfo := this.Data[0]
		// 对编码密文进行base64解码
		decodeBytes, err := base64.StdEncoding.DecodeString(cpinfo.Encrypt_certificate.Ciphertext)
		if err != nil {
			return nil, err
		}

		c, err := aes.NewCipher([]byte(apiv3key))
		if err != nil {
			return nil, err
		}

		gcm, err := cipher.NewGCM(c)
		if err != nil {
			return nil, err
		}

		nonceSize := gcm.NonceSize()
		if len(decodeBytes) < nonceSize {
			return nil, errors.New("密文证书长度不够")
		}
		res := CertificateInfo{}
		res.Serial_no = cpinfo.Serial_no
		if cpinfo.Encrypt_certificate.Associated_data != "" {
			plaintext, err := gcm.Open(nil, []byte(cpinfo.Encrypt_certificate.Nonce), decodeBytes, []byte(cpinfo.Encrypt_certificate.Associated_data))
			if err != nil {
				return nil, err
			}
			res.Publickey = string(plaintext)
		} else {
			plaintext, err := gcm.Open(nil, []byte(cpinfo.Encrypt_certificate.Nonce), decodeBytes, nil)
			if err != nil {
				return nil, err
			}
			res.Publickey = string(plaintext)
		}
		return res, nil
	}
}
type CertificateInfo struct {
	Serial_no string `json:"serial_no"`
	Publickey string `json:"publickey"`
}


func CalcFileMD5(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	md5Handle := md5.New()      //创建 md5 句柄
	_, err = io.Copy(md5Handle, f)  //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	md := md5Handle.Sum(nil)    //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}

//敏感数据加密
func SensitiveDataEncryption(plaintext,Publickey string) (string ,error) {
	block_pub, _ := pem.Decode([]byte(Publickey))

	if block_pub == nil || block_pub.Type != "CERTIFICATE" {
		log.Println("解码包含平台公钥的PEM块失败！")
		return "", errors.New("解码包含平台公钥的PEM块失败")
	}
	var err error
	PFPubc, err := x509.ParseCertificate(block_pub.Bytes)
	if err != nil {
		return "", err
	}

	secretMessage := []byte(plaintext)
	rng := rand.Reader

	cipherdata, err := rsa.EncryptOAEP(sha1.New(), rng, PFPubc.PublicKey.(*rsa.PublicKey), secretMessage, nil)
	if err != nil {
		return "", err
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherdata)
	return ciphertext, nil
}
