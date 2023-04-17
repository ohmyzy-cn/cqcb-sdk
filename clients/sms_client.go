package clients

import (
	"cqcb-sdk/model"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
)

type SmsClient struct {
	*req.Client
	config *model.SmsClientConfig
}

func NewSmsClient(config *model.SmsClientConfig, logger req.Logger) *SmsClient {
	c := req.C().
		SetCommonHeaders(map[string]string{
			"Accept":      "application/json",
			"ContentType": "application/json",
		}).
		SetBaseURL(config.ServerAddr).
		SetLogger(logger)
	if config.Debug {
		c.EnableDumpAll()
		c.EnableDebugLog()
	}

	return &SmsClient{
		Client: c,
		config: config,
	}
}

func (s *SmsClient) SendMessage(messages ...*model.SmsMessage) (*model.SmsResponse, error) {
	resp, err := s.R().SetBody(map[string]interface{}{
		"head": map[string]string{
			"account":   s.config.Account,
			"password":  s.config.Password,
			"sendtype":  s.config.SendType,
			"classcode": s.config.ClassCode,
		},
		"body": map[string][]*model.SmsMessage{
			"list": messages,
		},
	}).Post("SERVICE_HTTP/submitJsonMessage")
	if err != nil {
		return nil, fmt.Errorf("request sms api failed, %s", err)
	}
	bytes, err := resp.ToBytes()
	if err != nil {
		return nil, fmt.Errorf("get response result failed, %s", err)
	}
	// 由于消息中心响应头为xml格式，数据又是json格式，需要手动对结果转换成结构体
	var res model.SmsResponse
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		return nil, fmt.Errorf("decode response failed, %s", err)
	}
	return &res, nil
}
