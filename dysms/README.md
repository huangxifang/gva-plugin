# 阿里云短信【gin-vue-admin插件】


短信接口： /dysms/register [post] 已配置swagger

    注册验证码： /dysms/smsCode [post] 已配置swagger
    入参：
    type SmsRequest struct {
        SmsType      string                 //类型 register注册 sign身份验证
        Mobile       string                 //手机
        RequestPrams map[string]interface{} //短信参数
    }

## 安装阿里云短信SDK
```bash
go get github.com/alibabacloud-go/darabonba-openapi/client
go get github.com/alibabacloud-go/dysmsapi-20170525/v2
```

## 引用仓库
在`server/initialize/plugin.go` 文件内引用

```go
import (
    "github.com/huangxifang/gva-plugin/dysms"
)

```
## 初始化插件

在 `InstallPlugin`方法内新增阿里云短信插件初始化
```go
func InstallPlugin(PublicGroup *gin.RouterGroup, PrivateGroup *gin.RouterGroup) {

    //阿里云短信插件
    PluginInit(PublicGroup, dysms.CreateDysmsPlug(
        global.GVA_CONFIG.Dysms.AccessKeyId,
        global.GVA_CONFIG.Dysms.AccessKeySecret,
        global.GVA_CONFIG.Dysms.SignName,
        global.GVA_CONFIG.Dysms.RegisterTemplateCode,
        global.GVA_CONFIG.Dysms.SignTemplateCode,
    ))
}

```

## 请求举例
请求地址：/dysms/smsCode

Method：POST

Content-Type: application/json

参数：
```json
{
  SmsType: "sign",
  Mobile: "xxxxx",
  RequestPrams:{
    "code":"581111"
  }
}
```


