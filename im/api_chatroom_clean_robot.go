package im

const _API_CHATROOM_CLEAN_ROBOT_URL = "https://api.netease.im/nimserver/chatroom/cleanRobot.action"

type ChatroomCleanRobotParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 是否发送退出聊天室通知消息，默认为false
	// 是否必须: false
	Notify bool `json:"notify,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%B8%85%E7%A9%BA%E8%81%8A%E5%A4%A9%E5%AE%A4%E6%9C%BA%E5%99%A8%E4%BA%BA
// 清空聊天室机器人
func (y *YunxinIM) ApiChatroomCleanRobot(roomid int64, notify bool) *ImResp {
	return y.PostFrom(_API_CHATROOM_CLEAN_ROBOT_URL, ChatroomCleanRobotParam{roomid, notify})
}
