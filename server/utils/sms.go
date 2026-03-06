package utils

import (
	"encoding/json"
	"fmt"
	"hz-admin-base/utils/gurl"
	"strconv"
	"time"
)

const Sms_Key = "WcixF2Ae"
const Sms_Secret = "oiYEoPA3"
const Sms_Message = "测试】您的验证码是：%s。请不要把验证码泄露给其他人。"
const Sms_Url = "https://api.laaffic.com/v3/sendSms"
const Sms_App_Id = "3S1XYc01"

type SMS struct {
	Area    string
	Phone   string
	Captcha string
}

type SmsReq struct {
	AppId    string `json:"appId" form:"appId"`
	Numbers  string `json:"numbers" form:"numbers"`
	Content  string `json:"content" form:"content"`
	SenderId string `json:"senderId" form:"senderId"`
}

// {"status":"0","reason":"success","success":"1","fail":"0","array":[{"msgId":"6408061535011609096","number":"6289527907919","orderId":null}]}
type SmsAck struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func NewSMS(area, phone, captcha string) *SMS {
	return &SMS{
		Area:    area,
		Phone:   phone,
		Captcha: captcha,
	}
}

func (s *SMS) SendMess() SmsAck {
	tmp := strconv.FormatInt(time.Now().Unix(), 10)
	var signStr = Sms_Key + Sms_Secret + tmp
	sign := MD5V([]byte(signStr))
	g := gurl.NewGurl("post", Sms_Url)
	g.Header = map[string]string{
		"Content-Type": "application/json;charset=UTF-8",
		"Sign":         sign,
		"Timestamp":    tmp,
		"Api-Key":      Sms_Key,
	}

	g.SetData(SmsReq{
		AppId:    Sms_App_Id,
		Numbers:  s.Area + s.Phone,
		Content:  fmt.Sprintf(Sms_Message, s.Captcha),
		SenderId: Sms_Key,
	})

	res, err := g.Do()
	if err != nil {
		return SmsAck{Status: "-20", Reason: err.Error()}
	}
	ack := SmsAck{}
	err = json.Unmarshal(res, &ack)
	if err != nil {
		return SmsAck{}
	}

	return ack
}
