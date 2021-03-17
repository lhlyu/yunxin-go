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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E5%8D%95%E5%90%91%E6%92%A4%E5%9B%9E%E6%B6%88%E6%81%AF
// 单向撤回消息
func (y *YunxinIM) ApiMsgDelMsgOneWay(param *MsgDelMsgOneWayParam) *ImResp {
	return y.PostFrom(_API_MSG_DEL_MSG_ONE_WAY_URL, param)
}
