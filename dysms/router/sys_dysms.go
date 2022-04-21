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
		dysmsRouter.POST("register", DysmsApi.Register) // 注册短信验证码
	}
}
