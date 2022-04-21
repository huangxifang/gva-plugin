package model

type SmsRequest struct {
	SmsType      string
	Mobile       string
	RequestPrams map[string]interface{}
}
