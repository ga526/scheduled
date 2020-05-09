package thirdparty

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * 容联云
 * 方法  stop  run  start
 */

var (
	baseUrl = "https://app.cloopen.com:8883/"
)

// 创建容联云通讯类
type RLYSMS struct {
	flag        bool
	baseUrl     string
	accountSid  string
	authorToken string
	appId       string
	version     string
}

func NewRLYSMS(accountSid string, authorToken string, appId string, version string) *RLYSMS {
	if accountSid != "" && authorToken != "" && appId != "" {
		sms := new(RLYSMS)
		sms.flag = true
		sms.baseUrl = baseUrl
		sms.accountSid = accountSid
		sms.authorToken = authorToken
		sms.appId = appId
		sms.version = version
		return sms
	} else {
		return nil
	}
}

func (this *RLYSMS) authorizationAndSig() (author string, sig string) {
	// 计算时间戳 格式化yyyyMMddHHmmss
	curTime := time.Now().Format("20060102150405")

	sigStr := this.accountSid + this.authorToken + curTime
	sigMD5 := md5.Sum([]byte(sigStr)) // 获取MD5 签名
	// 获取md5 签名
	sig = strings.ToUpper(fmt.Sprintf("%x", sigMD5)) // 转化成16进制并转大写

	//Authorization 字符构造
	authStr := this.accountSid + ":" + curTime
	// base64编码
	enc := base64.StdEncoding
	author = enc.EncodeToString([]byte(authStr)) // 生成验证信息
	return
}

// 发送的消息与手机号码个数要匹配
func (this *RLYSMS) mSGJSON(data []string, template string,mobiles ...string) string {
	/*
			{"to":"13911281234,15010151234,13811431234","appId":
		 "ff8080813fc70a7b013fc72312324213","templateId":"1","datas":["替换内容","替换内容"]}
	*/
	str := ""
	for _,item := range data{
		str += "'"+item+"'"
	}

	body := map[string]string{
		"to":         strings.Join(mobiles, ","),
		"appId":      this.appId,
		"templateId"  : template,
		"datas" : "["+str+"]",
	}

	// 序列化成json
	res, _ := json.Marshal(body)
	fmt.Println(string(res))
	return string(res)
}

// 给一个手机号码发消息,支持json格式字符串
// message: 要发送的消息内容
//duration: 用于告诉用户验证码有效时间。
//mobiles：要发送给那些手机可以传递多个号码
func (this *RLYSMS) SendMsMs(data []string,temlate string,mobiles ...string) (rsp *http.Response, err error) {
	if this.flag == false {
		return nil, fmt.Errorf("请使用提供的创建方法创建对象")
	}
	//POST /2013-12-26/Accounts/abcdefghijklmnopqrstuvwxyz012345/SMS/TemplateSMS?sig=
	// C1F20E7A9733CE94F680C70A1DBABCDE

	// 创建http客户端
	client := http.Client{}

	// 创建要发送的请求
	// 生成URL
	author, sig := this.authorizationAndSig()
	smsUrl := this.baseUrl + this.version + "/" + "Accounts/" + this.accountSid + "/SMS/TemplateSMS?sig=" + sig

	content := this.mSGJSON(data, temlate, mobiles...)
	// 构造请求体并添加内容
	var body bytes.Buffer
	_, werr := body.Write([]byte(content))
	if err != nil {
		fmt.Println(werr)
	}
	// 构造请求
	req, err := http.NewRequest("POST", smsUrl, &body)
	if err != nil {
		return nil, err
	}
	// 添加请求头
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")
	req.Header.Add("Content-Length", strconv.Itoa(len(content)))
	req.Header.Add("Authorization", author)

	// 同步发送请求，是否异步让用户自己决定！
	rsp, err = client.Do(req)
	defer rsp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}else{
		body, _ := ioutil.ReadAll(rsp.Body)

		fmt.Println(string(body))
	}
	return
}
