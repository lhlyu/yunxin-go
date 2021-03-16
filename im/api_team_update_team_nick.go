package im

const _API_TEAM_UPDATE_TEAM_NICK_URL = "https://api.netease.im/nimserver/team/updateTeamNick.action"

type TeamUpdateTeamNickParam struct {
	// 群唯一标识，创建群时网易云信服务器产生并返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主
	// accid
	// 是否必须: true
	Owner string `json:"owner"`

	// 要修改群昵称的群成员
	// accid
	// 是否必须: true
	Accid string `json:"accid"`

	// accid
	// 对应的群昵称，最大长度32字符
	// 是否必须: false
	Nick string `json:"nick,omitempty"`

	// 自定义扩展字段，最大长度1024字节
	// 是否必须: false
	Custom string `json:"custom,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#修改群昵称
// 修改群昵称
func (y *YunxinIM) ApiTeamUpdateTeamNick(param *TeamUpdateTeamNickParam) *ImResp {
	return y.PostFrom(_API_TEAM_UPDATE_TEAM_NICK_URL, param)
}
