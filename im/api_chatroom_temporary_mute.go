package im

const _API_CHATROOM_TEMPORARY_MUTE_URL = "https://api.netease.im/nimserver/chatroom/temporaryMute.action"

type ChatroomTemporaryMuteParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 操作者accid,必须是管理员或创建者
	// 是否必须: true
	Operator string `json:"operator"`

	// 被禁言的目标账号accid
	// 是否必须: true
	Target string `json:"target"`

	// 0:解除禁言;>0设置禁言的秒数，不能超过2592000秒(30天)
	// 是否必须: true
	MuteDuration int64 `json:"muteDuration"`

	// 操作完成后是否需要发广播，true或false，默认true
	// 是否必须: false
	NeedNotify string `json:"needNotify,omitempty"`

	// 通知广播事件中的扩展字段，长度限制2048字符
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E8%AE%BE%E7%BD%AE%E4%B8%B4%E6%97%B6%E7%A6%81%E8%A8%80%E7%8A%B6%E6%80%81
// 设置临时禁言状态
func (y *YunxinIM) ApiChatroomTemporaryMute(param *ChatroomTemporaryMuteParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_TEMPORARY_MUTE_URL, param)
}
