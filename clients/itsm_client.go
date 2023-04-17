package clients

import (
	"cqcb-sdk/model"
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
)

type ItsmClient struct {
	*req.Client
	config   *model.ItsmClientConfig
	isLogged bool
}

func NewItsmClient(config *model.ItsmClientConfig, logger req.Logger) *ItsmClient {
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

	return &ItsmClient{
		Client: c,
		config: config,
	}
}

func (i *ItsmClient) Login() (*model.ItsmLoginResponse, error) {
	var res model.ItsmLoginResponse
	_, err := i.R().
		SetBody(map[string]string{
			"loginId":  i.config.Account,
			"loginPwd": i.config.Password,
		}).
		SetSuccessResult(&res).
		Post("o/simo_itsm/apis/http/v1/auth")
	if err != nil {
		return nil, fmt.Errorf("request sms api failed, %s", err)
	}
	return &res, nil
}

func (i *ItsmClient) SendWorkOrder(token string, userIds string, fields *model.ItsmWorkOrderFields) (*model.ItsmWorkOrderResponse, error) {
	var resp model.ItsmWorkOrderResponse
	form, err := json.Marshal(&model.ItsmWorkOrderRequest{
		Fields:       fields,
		ModelId:      i.config.ModelId,
		TargetUserId: userIds,
	})
	if err != nil {
		return nil, fmt.Errorf("serialization failed, %s", err)
	}
	_, err = i.R().SetHeaders(map[string]string{
		"Authorization": token,
		"Content-Type":  "multipart/form-data",
	}).SetFiles(map[string]string{
		"workOrderVo": string(form),
	}).SetSuccessResult(&resp).
		Post("o/simo_itsm/apis/http/v1/workorder/create")
	return &resp, err
}
