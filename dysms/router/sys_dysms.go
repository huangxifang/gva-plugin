package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/api"
	"github.com/gin-gonic/gin"
)

type DysmsRouter struct{}

func (s *DysmsRouter) InitDysmsRouter(Router *gin.RouterGroup) {
	dysmsRouter := Router
	DysmsApi := api.ApiGroupApp.DysmsApi
	{
		dysmsRouter.POST("smsCode", DysmsApi.SmsCode) // 注册-身份验证短信验证码
	}
}
