package im

const _API_MSG_BROADCAST_MSG_URL = "https://api.netease.im/nimserver/msg/broadcastMsg.action"

type MsgBroadcastMsgParam struct {
	// 广播消息内容，最大4096字符
	// 是否必须: true
	Body string `json:"body"`

	// 发送者accid, 用户帐号，最大长度32字符，必须保证一个APP内唯一
	// 是否必须: false
	From string `json:"from,omitempty"`

	// 是否存离线，true或false，默认false
	// 是否必须: false
	IsOffline string `json:"isOffline,omitempty"`

	// 存离线状态下的有效期，单位小时，默认7天
	// 是否必须: false
	Ttl int `json:"ttl,omitempty"`

	// 目标客户端，默认所有客户端，jsonArray，格式：["ios","aos","pc","web","mac"]
	// 是否必须: false
	TargetOs string `json:"targetOs,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E6%B6%88%E6%81%AF%E5%8A%9F%E8%83%BD?#%E5%8F%91%E9%80%81%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
// 发送广播消息
func (y *YunxinIM) ApiMsgBroadcastMsg(param *MsgBroadcastMsgParam) *ImResp {
	return y.PostFrom(_API_MSG_BROADCAST_MSG_URL, param)
}
