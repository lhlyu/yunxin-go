package im

const _API_CHATROOM_DELETE_HISTORY_MESSAGE_URL = "https://api.netease.im/nimserver/chatroom/deleteHistoryMessage.action"

type ChatroomDeleteHistoryMessageParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 消息发送者的accid
	// 是否必须: true
	FromAcc string `json:"fromAcc"`

	// 消息的时间戳，单位毫秒，应该拿到原始消息中的时间戳为参数
	// 是否必须: true
	MsgTimetag int64 `json:"msgTimetag"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95?#%E5%88%A0%E9%99%A4%E8%81%8A%E5%A4%A9%E5%AE%A4%E4%BA%91%E7%AB%AF%E5%8E%86%E5%8F%B2%E6%B6%88%E6%81%AF
// 删除聊天室云端历史消息
func (y *YunxinIM) ApiChatroomDeleteHistoryMessage(roomid int64, fromAcc string, msgTimetag int64) *ImResp {
	return y.PostFrom(_API_CHATROOM_DELETE_HISTORY_MESSAGE_URL, ChatroomDeleteHistoryMessageParam{roomid, fromAcc, msgTimetag})
}
