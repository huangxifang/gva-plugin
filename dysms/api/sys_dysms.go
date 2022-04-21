package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/service"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type DysmsApi struct{}

//注册验证码
func (s *DysmsApi) Register(c *gin.Context) {
	var sendCode model.SmsRequest
	_ = c.ShouldBindJSON(&sendCode)
	if sendCode.SmsType == "" {
		sendCode.SmsType = "register"
	}

	var requestPrams = make(map[string]interface{})
	code := base64Captcha.RandText(6, "0123456789")
	requestPrams["code"] = code
	sendCode.RequestPrams = requestPrams
	_result, err := service.ServiceGroupApp.SendSms(sendCode)
	if err != nil || *_result.Body.Code != "OK" {
		global.GVA_LOG.Error(*_result.Body.Message, zap.Error(err))
		response.FailWithMessage("发送失败", c)
	} else {
		response.OkWithData("发送成功", c)
	}
}
