package im

const _API_CHATROOM_TOGGLE_CLOSE_STAT_URL = "https://api.netease.im/nimserver/chatroom/toggleCloseStat.action"

type ChatroomToggleCloseStatParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 操作者账号，必须是创建者才可以操作
	// 是否必须: true
	Operator string `json:"operator"`

	// true或false，false:关闭聊天室；true:打开聊天室
	// 是否必须: true
	Valid string `json:"valid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E4%BF%AE%E6%94%B9%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%BC%80/%E5%85%B3%E9%97%AD%E7%8A%B6%E6%80%81
// 修改聊天室开/关闭状态
func (y *YunxinIM) ApiChatroomToggleCloseStat(roomid int64, operator string, valid string) *ImResp {
	return y.PostFrom(_API_CHATROOM_TOGGLE_CLOSE_STAT_URL, ChatroomToggleCloseStatParam{roomid, operator, valid})
}
