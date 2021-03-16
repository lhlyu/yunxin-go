package im

const _API_SUPERTEAM_RECALL_MSG_URL = "https://api.netease.im/nimserver/superteam/recallMsg.action"

type SuperteamRecallMsgParam struct {
	// 要撤回消息的msgid
	// 是否必须: true
	DeleteMsgid string `json:"deleteMsgid"`

	// 要撤回消息的创建时间(创建时间为云信服务器生成的消息发送时间戳)
	// 是否必须: true
	Timetag string `json:"timetag"`

	// 发送者accid,
	// 用户帐号，最大长度32字符
	// 是否必须: true
	From string `json:"from"`

	// 群tid
	// 是否必须: true
	To string `json:"to"`

	// 可以带上对应的描述
	// 是否必须: false
	Msg string `json:"msg,omitempty"`

	// 1表示忽略撤回时间检测，0表示不忽略，其它为非法参数，默认0，如果需要撤回时间检测，不填即可
	// 是否必须: false
	IgnoreTime string `json:"ignoreTime,omitempty"`

	// 推送内容，不超过500字符
	// 是否必须: false
	PushContent string `json:"pushContent,omitempty"`

	// 推送对应的payload，必须是JSON，不能超过2k字符
	// 是否必须: false
	PushPayload string `json:"pushPayload,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#撤回消息
// 撤回消息
func (y *YunxinIM) ApiSuperteamRecallMsg(param *SuperteamRecallMsgParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_RECALL_MSG_URL, param)
}
