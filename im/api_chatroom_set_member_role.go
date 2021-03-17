package im

const _API_CHATROOM_SET_MEMBER_ROLE_URL = "https://api.netease.im/nimserver/chatroom/setMemberRole.action"

type ChatroomSetMemberRoleParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 操作者账号accid
	// 是否必须: true
	Operator string `json:"operator"`

	// 被操作者账号accid
	// 是否必须: true
	Target string `json:"target"`

	// 操作：
	// 1: 设置为管理员，operator必须是创建者
	// 2:设置普通等级用户，operator必须是创建者或管理员
	// -1:设为黑名单用户，operator必须是创建者或管理员
	// -2:设为禁言用户，operator必须是创建者或管理员
	// 是否必须: true
	Opt int `json:"opt"`

	// true或false，true:设置；false:取消设置；
	// 执行“取消”设置后，若成员非禁言且非黑名单，则变成游客
	// 是否必须: true
	Optvalue string `json:"optvalue"`

	// 通知扩展字段，长度限制2048，请使用json格式
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E8%AE%BE%E7%BD%AE%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%86%85%E7%94%A8%E6%88%B7%E8%A7%92%E8%89%B2
// 设置聊天室内用户角色
func (y *YunxinIM) ApiChatroomSetMemberRole(param *ChatroomSetMemberRoleParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_SET_MEMBER_ROLE_URL, param)
}
