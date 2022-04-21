package config

type Dysms struct {
	AccessKeyId          string `mapstructure:"access-key-id" json:"accessKeyId" yaml:"access-key-id"`
	AccessKeySecret      string `mapstructure:"access-key-secret" json:"accessKeySecret" yaml:"access-key-secret"`
	SignName             string `mapstructure:"sign-name" json:"signName" yaml:"sign-name"`
	RegisterTemplateCode string `mapstructure:"register-template-code" json:"registerTemplateCode" yaml:"register-template-code"` //注册会员验证码模板编码
	SignTemplateCode     string `mapstructure:"sign-template-code" json:"signTemplateCode" yaml:"sign-template-code"`             //验证身份模板编码
}
