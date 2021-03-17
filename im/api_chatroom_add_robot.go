package im

const _API_CHATROOM_ADD_ROBOT_URL = "https://api.netease.im/nimserver/chatroom/addRobot.action"

type ChatroomAddRobotParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 机器人账号accid列表，必须是有效账号，账号数量上限100个
	// 是否必须: true
	Accids string `json:"accids"`

	// 机器人信息扩展字段，请使用json格式，长度4096字符
	// 是否必须: false
	RoleExt string `json:"roleExt,omitempty"`

	// 机器人进入聊天室通知的扩展字段，请使用json格式，长度2048字符
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%BE%80%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%86%85%E6%B7%BB%E5%8A%A0%E6%9C%BA%E5%99%A8%E4%BA%BA
// 往聊天室内添加机器人
func (y *YunxinIM) ApiChatroomAddRobot(param *ChatroomAddRobotParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_ADD_ROBOT_URL, param)
}
