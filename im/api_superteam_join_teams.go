package im

const _API_SUPERTEAM_JOIN_TEAMS_URL = "https://api.netease.im/nimserver/superteam/joinTeams.action"

type SuperteamJoinTeamsParam struct {
	// 用户帐号，最大长度32字符
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#获取某用户所加入的群信息
// 获取某用户所加入的群信息
func (y *YunxinIM) ApiSuperteamJoinTeams(accid string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_JOIN_TEAMS_URL, SuperteamJoinTeamsParam{accid})
}
