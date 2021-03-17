package im

const _API_CHATROOM_GET_URL = "https://api.netease.im/nimserver/chatroom/get.action"

type ChatroomGetParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 是否需要返回在线人数，true或false，默认false
	// 是否必须: false
	NeedOnlineUserCount string `json:"needOnlineUserCount,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%9F%A5%E8%AF%A2%E8%81%8A%E5%A4%A9%E5%AE%A4%E4%BF%A1%E6%81%AF
// 查询聊天室信息
func (y *YunxinIM) ApiChatroomGet(roomid int64, needOnlineUserCount string) *ImResp {
	return y.PostFrom(_API_CHATROOM_GET_URL, ChatroomGetParam{roomid, needOnlineUserCount})
}
