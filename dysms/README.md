# GVA 阿里云短信插件

-----------------------------------------------

本插件用于注册验证码、身份验证码，短信通知
1. 使用场景
- 注册会员
- 修改密码，身份验证
- 会员消息通知

2. 配置说明

在gin-vue-admin 主程序的initialize中的plugin的InstallPlugin 函数中写入如下代码
```go
//阿里云短信插件
PluginInit(PublicGroup, dysms.CreateDysmsPlug(
    global.GVA_CONFIG.Dysms.AccessKeyId,
    global.GVA_CONFIG.Dysms.AccessKeySecret,
    global.GVA_CONFIG.Dysms.SignName,
    global.GVA_CONFIG.Dysms.RegisterTemplateCode,
    global.GVA_CONFIG.Dysms.SignTemplateCode,
))
```
在`config.yaml`文件中配置阿里云短信信息
```yaml
dysms:
  access-key-id: xxxxxxx
  access-key-secret: xxxxxxx
  sign-name: 签名
  register-template-code: xxxx
  sign-template-code: xxxxx
```
> 增加配置到后台配置

在`web/src/view/systemTools/system/system.vue` 配置底部添加以下代码

```js
<el-collapse-item title="阿里云短信验证码" name="14">
  <el-form-item label="AccessKeyId">
    <el-input v-model="config.dysms.accessKeyId" />
  </el-form-item>
  <el-form-item label="AccessKeySecret">
    <el-input v-model="config.dysms.accessKeySecret" />
  </el-form-item>
  <el-form-item label="签名">
    <el-input v-model="config.dysms.signName" />
  </el-form-item>
  <el-form-item label="注册会员验证码模板编码">
    <el-input v-model="config.dysms.registerTemplateCode" />
  </el-form-item>
  <el-form-item label="验证身份模板编码">
    <el-input v-model="config.dysms.signTemplateCode" />
  </el-form-item>
</el-collapse-item>
```
在`config`变量内容新增
```js
dysms: {}
```


3. 参数说明
```go
type SmsRequest struct {
    SmsType      string                 //类型 register注册 sign身份验证
    Mobile       string                 //手机
    RequestPrams map[string]interface{} //短信参数
}
```

4. 方法API（可调用方法）
```go
var info model.SmsRequest
utils.Send(info)
```

5. 可调用接口


    接口： /dysms/smsCode [post] 
    参数：
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
