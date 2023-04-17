package model

import "time"

type ItsmClientConfig struct {
	CacheTtl   time.Duration // token缓存时间长度
	ServerAddr string        // 服务端地址
	Account    string        // 账号
	Password   string        // 密码
	ModelId    string        // 工单系统的模板ID，由工单系统系统
	VoiceLevel []string      // 是否语言通知的告警级别
	Debug      bool
}

type ItsmLoginResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Obj     struct {
		Token string `json:"token"`
	} `json:"obj"`
}

type ItsmWorkOrderRequest struct {
	Fields       *ItsmWorkOrderFields `json:"fields"`
	ModelId      string               `json:"modelId"`
	TargetUserId string               `json:"targetUserId"`
}

type ItsmWorkOrderFields struct {
	Title  string `json:"title"`
	Memo   string `json:"memo"`
	Xtmc   string `json:"xtmc"`
	BGSJ   string `json:"BGSJ"`
	SFYYTZ string `json:"SFYYTZ"`
}

type ItsmWorkOrderResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Code    int    `json:"code"`
	Obj     string `json:"obj"`
}
