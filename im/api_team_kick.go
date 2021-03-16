package im

const _API_TEAM_KICK_URL = "https://api.netease.im/nimserver/team/kick.action"

type TeamKickParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 管理员的accid，用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 被移除人的accid，用户账号，最大长度32字符;注：member或members任意提供一个，优先使用member参数
	// 是否必须: false
	Member string `json:"member,omitempty"`

	// ["aaa","bbb"]（JSONArray对应的accid，如果解析出错，会报414）一次最多操作200个accid;
	// 注：member或members任意提供一个，优先使用member参数
	// 是否必须: false
	Members string `json:"members,omitempty"`

	// 自定义扩展字段，最大长度512
	// 是否必须: false
	Attach string `json:"attach,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#踢人出群
// 踢人出群
func (y *YunxinIM) ApiTeamKick(param *TeamKickParam) *ImResp {
	return y.PostFrom(_API_TEAM_KICK_URL, param)
}
