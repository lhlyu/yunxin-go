package im

const _API_CHATROOM_UPDATE_MY_ROOM_ROLE_URL = "https://api.netease.im/nimserver/chatroom/updateMyRoomRole.action"

type ChatroomUpdateMyRoomRoleParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 需要变更角色信息的accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 变更的信息是否需要持久化，默认false，仅对聊天室固定成员生效
	// 是否必须: false
	Save bool `json:"save,omitempty"`

	// 是否需要做通知
	// 是否必须: false
	NeedNotify bool `json:"needNotify,omitempty"`

	// 通知的内容，长度限制2048
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`

	// 聊天室室内的角色信息：昵称，不超过64个字符
	// 是否必须: false
	Nick string `json:"nick,omitempty"`

	// 聊天室室内的角色信息：头像
	// 是否必须: false
	Avator string `json:"avator,omitempty"`

	// 聊天室室内的角色信息：开发者扩展字段
	// 是否必须: false
	Ext string `json:"ext,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%8F%98%E6%9B%B4%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%86%85%E7%9A%84%E8%A7%92%E8%89%B2%E4%BF%A1%E6%81%AF
// 变更聊天室内的角色信息
func (y *YunxinIM) ApiChatroomUpdateMyRoomRole(param *ChatroomUpdateMyRoomRoleParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_UPDATE_MY_ROOM_ROLE_URL, param)
}
