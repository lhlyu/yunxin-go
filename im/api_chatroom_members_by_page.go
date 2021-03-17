package im

const _API_CHATROOM_MEMBERS_BY_PAGE_URL = "https://api.netease.im/nimserver/chatroom/membersByPage.action"

type ChatroomMembersByPageParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 需要查询的成员类型,0:固定成员;1:非固定成员;2:仅返回在线的固定成员
	// 是否必须: true
	Type int `json:"type"`

	// 单位毫秒，按时间倒序最后一个成员的时间戳,0表示系统当前时间
	// 是否必须: true
	Endtime int64 `json:"endtime"`

	// 返回条数，<=100
	// 是否必须: true
	Limit int64 `json:"limit"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%88%86%E9%A1%B5%E8%8E%B7%E5%8F%96%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8
// 分页获取成员列表
func (y *YunxinIM) ApiChatroomMembersByPage(param *ChatroomMembersByPageParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_MEMBERS_BY_PAGE_URL, param)
}
