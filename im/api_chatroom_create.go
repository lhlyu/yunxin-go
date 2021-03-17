package im

const _API_CHATROOM_CREATE_URL = "https://api.netease.im/nimserver/chatroom/create.action"

type ChatroomCreateParam struct {
	// 聊天室属主的账号accid
	// 是否必须: true
	Creator string `json:"creator"`

	// 聊天室名称，长度限制128个字符
	// 是否必须: true
	Name string `json:"name"`

	// 公告，长度限制4096个字符
	// 是否必须: false
	Announcement string `json:"announcement,omitempty"`

	// 直播地址，长度限制1024个字符
	// 是否必须: false
	Broadcasturl string `json:"broadcasturl,omitempty"`

	// 扩展字段，最长4096字符
	// 是否必须: false
	Ext string `json:"ext,omitempty"`

	// 队列管理权限：0:所有人都有权限变更队列，1:只有主播管理员才能操作变更。默认0
	// 是否必须: false
	Queuelevel int `json:"queuelevel,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%88%9B%E5%BB%BA%E8%81%8A%E5%A4%A9%E5%AE%A4
// 创建聊天室
func (y *YunxinIM) ApiChatroomCreate(param *ChatroomCreateParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_CREATE_URL, param)
}
