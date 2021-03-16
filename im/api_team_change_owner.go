package im

const _API_TEAM_CHANGE_OWNER_URL = "https://api.netease.im/nimserver/team/changeOwner.action"

type TeamChangeOwnerParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// 新群主帐号，最大长度32字符
	// 是否必须: true
	Newowner string `json:"newowner"`

	// 1:群主解除群主后离开群，2：群主解除群主后成为普通成员。其它414
	// 是否必须: true
	Leave int `json:"leave"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#移交群主
// 移交群主
func (y *YunxinIM) ApiTeamChangeOwner(param *TeamChangeOwnerParam) *ImResp {
	return y.PostFrom(_API_TEAM_CHANGE_OWNER_URL, param)
}
