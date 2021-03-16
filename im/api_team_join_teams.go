package im

const _API_TEAM_JOIN_TEAMS_URL = "https://api.netease.im/nimserver/team/joinTeams.action"

type TeamJoinTeamsParam struct {
	// 要查询用户的accid
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#获取某用户所加入的群信息
// 获取某用户所加入的群信息
func (y *YunxinIM) ApiTeamJoinTeams(accid string) *ImResp {
	return y.PostFrom(_API_TEAM_JOIN_TEAMS_URL, TeamJoinTeamsParam{accid})
}
