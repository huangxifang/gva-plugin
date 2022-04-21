package utils

import (
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/model"
)

//
//  发送短信
//  @Description:
//  @param sendType 类型register注册，sign验证身份
//  @param mobile 手机号
//  @param prams 参数 map类型数据
//
func Send(info model.SmsRequest) (_result *dysmsapi.SendSmsResponse, err error) {
	client, _ := CreateClient()
	request := &dysmsapi.SendSmsRequest{}
	request.SetPhoneNumbers(info.Mobile)
	request.SetSignName(global.GlobalConfig.SignName)
	templateCode := global.GlobalConfig.RegisterTemplateCode
	//模板码
	if info.SmsType == "sign" {
		templateCode = global.GlobalConfig.SignTemplateCode
	}
	request.SetTemplateCode(templateCode)
	// 参数
	data := MapToJson(info.RequestPrams)
	request.SetTemplateParam(data)
	_result, _err := client.SendSms(request)
	return _result, _err
}

//string 转map
func JsonToMap(str string) map[string]interface{} {
	//map 转json
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

// map 转string
func MapToJson(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *dysmsapi.Client, _err error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &global.GlobalConfig.AccessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &global.GlobalConfig.AccessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi.Client{}
	_result, _err = dysmsapi.NewClient(config)
	return _result, _err
}
