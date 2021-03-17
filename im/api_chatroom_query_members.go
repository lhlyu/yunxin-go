package im

const _API_CHATROOM_QUERY_MEMBERS_URL = "https://api.netease.im/nimserver/chatroom/queryMembers.action"

type ChatroomQueryMembersParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// ["abc","def"], 账号列表，最多200条
	// 是否必须: true
	Accids string `json:"accids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%89%B9%E9%87%8F%E8%8E%B7%E5%8F%96%E5%9C%A8%E7%BA%BF%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF
// 批量获取在线成员信息
func (y *YunxinIM) ApiChatroomQueryMembers(roomid int64, accids string) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUERY_MEMBERS_URL, ChatroomQueryMembersParam{roomid, accids})
}
