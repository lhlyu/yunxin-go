package im

const _API_TEAM_UPDATE_TEAM_NICK_URL = "https://api.netease.im/nimserver/team/updateTeamNick.action"

type TeamUpdateTeamNickParam struct {
	// 群唯一标识，创建群时网易云信服务器产生并返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主 accid
	// 是否必须: true
	Owner string `json:"owner"`

	// 要修改群昵称的群成员 accid
	// 是否必须: true
	Accid string `json:"accid"`

	// accid 对应的群昵称，最大长度32字符
	// 是否必须: false
	Nick string `json:"nick,omitempty"`

	// 自定义扩展字段，最大长度1024字节
	// 是否必须: false
	Custom string `json:"custom,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E7%BE%A4%E6%98%B5%E7%A7%B0
// 修改群昵称
func (y *YunxinIM) ApiTeamUpdateTeamNick(param *TeamUpdateTeamNickParam) *ImResp {
	return y.PostFrom(_API_TEAM_UPDATE_TEAM_NICK_URL, param)
}
