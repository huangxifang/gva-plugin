# 阿里云短信【gin-vue-admin插件】


短信接口： /dysms/register [post] 已配置swagger

    注册验证码： /dysms/smsCode [post] 已配置swagger
    入参：
    type SmsRequest struct {
        SmsType      string //register 注册 sign身份验证
        Mobile       string //手机
    }

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
