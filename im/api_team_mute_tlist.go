package im

const _API_TEAM_MUTE_TLIST_URL = "https://api.netease.im/nimserver/team/muteTlist.action"

type TeamMuteTlistParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主accid
	// 是否必须: true
	Owner string `json:"owner"`

	// 禁言对象的accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 1-禁言，0-解禁
	// 是否必须: true
	Mute int `json:"mute"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#禁言群成员
// 禁言群成员
func (y *YunxinIM) ApiTeamMuteTlist(param *TeamMuteTlistParam) *ImResp {
	return y.PostFrom(_API_TEAM_MUTE_TLIST_URL, param)
}
