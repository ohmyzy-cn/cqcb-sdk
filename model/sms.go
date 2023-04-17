package model

type SmsClientConfig struct {
	ServerAddr string // 消息中心服务地址
	Account    string // 消息中心发送短信账号
	Password   string // 密码
	SendType   string // 短信类型，由消息中心提供
	ClassCode  string // 由消息中心提供
	Debug      bool
}

type SmsMessage struct {
	SeqId   string `json:"seqId"`
	Phone   string `json:"phone"`
	Content string `json:"content"`
}

type SmsResponse struct {
	Head struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"head"`
	Body struct {
		List []struct {
			Seqid   string `json:"seqid"`
			Msgid   string `json:"msgid"`
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"list"`
	} `json:"body"`
}
