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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%A7%BB%E4%BA%A4%E7%BE%A4%E4%B8%BB
// 移交群主
func (y *YunxinIM) ApiTeamChangeOwner(param *TeamChangeOwnerParam) *ImResp {
	return y.PostFrom(_API_TEAM_CHANGE_OWNER_URL, param)
}
