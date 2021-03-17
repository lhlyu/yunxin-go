package im

const _API_CHATROOM_UPDATE_IN_OUT_NOTIFICATION_URL = "https://api.netease.im/nimserver/chatroom/updateInOutNotification.action"

type ChatroomUpdateInOutNotificationParam struct {
	// 聊天室ID
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// true/false, 是否关闭进出通知
	// 是否必须: true
	Close bool `json:"close"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%85%B3%E9%97%AD%E6%8C%87%E5%AE%9A%E8%81%8A%E5%A4%A9%E5%AE%A4%E8%BF%9B%E5%87%BA%E9%80%9A%E7%9F%A5
// 关闭指定聊天室进出通知
func (y *YunxinIM) ApiChatroomUpdateInOutNotification(roomid int64, close bool) *ImResp {
	return y.PostFrom(_API_CHATROOM_UPDATE_IN_OUT_NOTIFICATION_URL, ChatroomUpdateInOutNotificationParam{roomid, close})
}
