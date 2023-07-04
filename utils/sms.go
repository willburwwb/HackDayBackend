package utils

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendSms(phone string, code string) {
	// 初始化短信客户端
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI5t8xaipr9spnCLTVmEec", "JUkoK0VnM5pZBSIadG6sRFqtLmkQ8e")
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}

	// 构建发送短信请求
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"                                   // 使用 HTTPS 协议
	request.PhoneNumbers = phone                               // 接收短信的手机号码
	request.SignName = "Aidea"                                 // 短信签名
	request.TemplateCode = "SMS_461855427"                     // 短信模板CODE
	request.TemplateParam = fmt.Sprintf(`{"code":"%s"}`, code) // 短信模板参数，JSON 格式

	// 发送短信
	response, err := client.SendSms(request)
	if err != nil {
		fmt.Println("Failed to send SMS:", err)
		return
	}

	DebugF("SMS Sent, RequestId:%s", response.RequestId)
}
