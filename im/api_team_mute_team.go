package im

const _API_TEAM_MUTE_TEAM_URL = "https://api.netease.im/nimserver/team/muteTeam.action"

type TeamMuteTeamParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 要操作的群成员accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 1：关闭消息提醒，2：打开消息提醒，其他值无效
	// 是否必须: true
	Ope int `json:"ope"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E4%BF%AE%E6%94%B9%E6%B6%88%E6%81%AF%E6%8F%90%E9%86%92%E5%BC%80%E5%85%B3
// 修改消息提醒开关
func (y *YunxinIM) ApiTeamMuteTeam(tid string, accid string, ope int) *ImResp {
	return y.PostFrom(_API_TEAM_MUTE_TEAM_URL, TeamMuteTeamParam{tid, accid, ope})
}
