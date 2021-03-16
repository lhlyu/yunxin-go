package im

const _API_TEAM_REMOVE_URL = "https://api.netease.im/nimserver/team/remove.action"

type TeamRemoveParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主用户帐号，最大长度32字符
	// 是否必须: true
	Owner string `json:"owner"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#解散群
// 解散群
func (y *YunxinIM) ApiTeamRemove(tid string, owner string) *ImResp {
	return y.PostFrom(_API_TEAM_REMOVE_URL, TeamRemoveParam{tid, owner})
}
