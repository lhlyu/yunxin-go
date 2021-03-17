package im

const _API_TEAM_ADD_MANAGER_URL = "https://api.netease.im/nimserver/team/addManager.action"

type TeamAddManagerParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`

	// ["aaa","bbb"](JSONArray对应的accid，如果解析出错会报414)，长度最大1024字符（一次添加最多10个管理员）
	// 是否必须: true
	Members string `json:"members"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%BB%BB%E5%91%BD%E7%AE%A1%E7%90%86%E5%91%98
// 任命管理员
func (y *YunxinIM) ApiTeamAddManager(tid string, owner string, members string) *ImResp {
	return y.PostFrom(_API_TEAM_ADD_MANAGER_URL, TeamAddManagerParam{tid, owner, members})
}
