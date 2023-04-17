package clients

import (
	"cqcb-sdk/model"
	"github.com/imroc/req/v3"
)

type EzsonarClient struct {
	*req.Client
	config   *model.SmsClientConfig
	isLogged bool
}

func (e *EzsonarClient) Login() {
}
