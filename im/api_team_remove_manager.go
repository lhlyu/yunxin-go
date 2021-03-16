package im

const _API_TEAM_REMOVE_MANAGER_URL = "https://api.netease.im/nimserver/team/removeManager.action"

type TeamRemoveManagerParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// ["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，长度最大1024字符（一次解除最多10个管理员）
	// 是否必须: true
	Members string `json:"members"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#移除管理员
// 移除管理员
func (y *YunxinIM) ApiTeamRemoveManager(tid string, owner string, members string) *ImResp {
	return y.PostFrom(_API_TEAM_REMOVE_MANAGER_URL, TeamRemoveManagerParam{tid, owner, members})
}