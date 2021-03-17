package im

const _API_TEAM_JOIN_TEAMS_URL = "https://api.netease.im/nimserver/team/joinTeams.action"

type TeamJoinTeamsParam struct {
	// 要查询用户的accid
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E6%9F%90%E7%94%A8%E6%88%B7%E6%89%80%E5%8A%A0%E5%85%A5%E7%9A%84%E7%BE%A4%E4%BF%A1%E6%81%AF
// 获取某用户所加入的群信息
func (y *YunxinIM) ApiTeamJoinTeams(accid string) *ImResp {
	return y.PostFrom(_API_TEAM_JOIN_TEAMS_URL, TeamJoinTeamsParam{accid})
}
