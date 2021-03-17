package im

const _API_CHATROOM_UPDATE_URL = "https://api.netease.im/nimserver/chatroom/update.action"

type ChatroomUpdateParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 聊天室名称，长度限制128个字符
	// 是否必须: false
	Name string `json:"name,omitempty"`

	// 公告，长度限制4096个字符
	// 是否必须: false
	Announcement string `json:"announcement,omitempty"`

	// 直播地址，长度限制1024个字符
	// 是否必须: false
	Broadcasturl string `json:"broadcasturl,omitempty"`

	// 扩展字段，长度限制4096个字符
	// 是否必须: false
	Ext string `json:"ext,omitempty"`

	// true或false,是否需要发送更新通知事件，默认true
	// 是否必须: false
	NeedNotify string `json:"needNotify,omitempty"`

	// 通知事件扩展字段，长度限制2048
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`

	// 队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更
	// 是否必须: false
	Queuelevel int `json:"queuelevel,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%9B%B4%E6%96%B0%E8%81%8A%E5%A4%A9%E5%AE%A4%E4%BF%A1%E6%81%AF
// 更新聊天室信息
func (y *YunxinIM) ApiChatroomUpdate(param *ChatroomUpdateParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_UPDATE_URL, param)
}
