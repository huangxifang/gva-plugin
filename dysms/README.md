# 阿里云短信【gin-vue-admin插件】


短信接口： /dysms/register [post] 已配置swagger

    注册验证码： /dysms/register [post] 已配置swagger
    入参：
    type SmsRequest struct {
        SmsType      string //register 注册 sign身份验证
        Mobile       string //手机
        RequestPrams map[string]interface{} //参数
    }