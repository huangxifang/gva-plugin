package dysms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/router"
	"github.com/gin-gonic/gin"
)

type dysmsPlugin struct{}

func CreateDysmsPlug(AccessKeyId string, AccessKeySecret string, SignName string, RegisterTemplateCode string, SignTemplateCode string) *dysmsPlugin {
	global.GlobalConfig.AccessKeyId = AccessKeyId
	global.GlobalConfig.AccessKeySecret = AccessKeySecret
	global.GlobalConfig.SignName = SignName
	global.GlobalConfig.RegisterTemplateCode = RegisterTemplateCode
	global.GlobalConfig.SignTemplateCode = SignTemplateCode
	return &dysmsPlugin{}
}

func (*dysmsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitDysmsRouter(group)
}

func (*dysmsPlugin) RouterPath() string {
	return "dysms"
}
