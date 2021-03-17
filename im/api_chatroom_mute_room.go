package im

const _API_CHATROOM_MUTE_ROOM_URL = "https://api.netease.im/nimserver/chatroom/muteRoom.action"

type ChatroomMuteRoomParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 操作者accid，必须是管理员或创建者
	// 是否必须: true
	Operator string `json:"operator"`

	// true或false
	// 是否必须: true
	Mute string `json:"mute"`

	// true或false，默认true
	// 是否必须: false
	NeedNotify string `json:"needNotify,omitempty"`

	// 通知扩展字段
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%B0%86%E8%81%8A%E5%A4%A9%E5%AE%A4%E6%95%B4%E4%BD%93%E7%A6%81%E8%A8%80
// 将聊天室整体禁言
func (y *YunxinIM) ApiChatroomMuteRoom(param *ChatroomMuteRoomParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_MUTE_ROOM_URL, param)
}
