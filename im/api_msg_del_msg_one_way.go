package im

const _API_MSG_DEL_MSG_ONE_WAY_URL = "https://api.netease.im/nimserver/msg/delMsgOneWay.action"

type MsgDelMsgOneWayParam struct {
	// 要撤回消息的msgid
	// 是否必须: true
	DeleteMsgid string `json:"deleteMsgid"`

	// 要撤回消息的创建时间
	// 是否必须: true
	Timetag int64 `json:"timetag"`

	// 13:表示点对点消息撤回，14:表示群消息撤回，其它为参数错误
	// 是否必须: true
	Type int `json:"type"`

	// 发消息的accid
	// 是否必须: true
	From string `json:"from"`

	// 如果点对点消息，为接收消息的accid,如果群消息，为对应群的tid
	// 是否必须: true
	To string `json:"to"`

	// 可以带上对应的描述
	// 是否必须: false
	Msg string `json:"msg,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/消息功能?#批量发送点对点普通消息
// 单向撤回消息
func (y *YunxinIM) ApiMsgDelMsgOneWay(param *MsgDelMsgOneWayParam) *ImResp {
	return y.PostFrom(_API_MSG_DEL_MSG_ONE_WAY_URL, param)
}
