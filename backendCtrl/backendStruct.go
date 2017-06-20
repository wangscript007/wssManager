package backendCtrl

import (
	"encoding/json"
)

type ResponseData struct {
	Code int    `json:"code"`
	Data *Data  `json:"data,omitempty"`
	Msg  string `json:"msg"`
}

type Data struct {
	UserData Usr    `json:"usr_data"`
	Action   Action `json:"action"`
}

type Usr struct {
	Usrname string `json:"usrname"`
	Token   string `json:"token"`
}

type Action struct {
	ActionCode  int    `json:"action_code"`
	ActionToken string `json:"action_token"`
}

const (
	WS_SHOW_ALL_STREAM = iota
	WS_GET_LIVE_PLAYER_COUNT
	WS_ENABLE_BLACK_LIST
	WS_SET_BLACK_LIST
	WS_ENABLE_WHITE_LIST
	WS_SET_WHITE_LIST
	WS_SET_UP_STREAM_APP
	WS_PULL_RTMP_STREAM
	WS_ADD_SINK
	WS_DEL_SINK
	WS_ADD_SOURCE
	WS_DEL_SOURCE
	WS_GET_SOURCE
)

func ParseRespData(data []byte) (respData *ResponseData, err error) {
	respData = &ResponseData{}
	err = json.Unmarshal(data, respData)
	return respData, err
}
