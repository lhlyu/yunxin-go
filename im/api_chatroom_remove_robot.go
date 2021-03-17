package im

const _API_CHATROOM_REMOVE_ROBOT_URL = "https://api.netease.im/nimserver/chatroom/removeRobot.action"

type ChatroomRemoveRobotParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 机器人账号accid列表，必须是有效账号，账号数量上限100个
	// 是否必须: true
	Accids string `json:"accids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E4%BB%8E%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%86%85%E5%88%A0%E9%99%A4%E6%9C%BA%E5%99%A8%E4%BA%BA
// 从聊天室内删除机器人
func (y *YunxinIM) ApiChatroomRemoveRobot(roomid int64, accids string) *ImResp {
	return y.PostFrom(_API_CHATROOM_REMOVE_ROBOT_URL, ChatroomRemoveRobotParam{roomid, accids})
}
