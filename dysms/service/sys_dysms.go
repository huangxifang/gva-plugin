package service

import (
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/dysms/utils"
)

type DysmsService struct{}

func (s *DysmsService) SendSms(info model.SmsRequest) (_result *dysmsapi.SendSmsResponse, err error) {
	return utils.Send(info)
}
