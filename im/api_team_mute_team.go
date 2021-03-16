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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#修改消息提醒开关
// 修改消息提醒开关
func (y *YunxinIM) ApiTeamMuteTeam(tid string, accid string, ope int) *ImResp {
	return y.PostFrom(_API_TEAM_MUTE_TEAM_URL, TeamMuteTeamParam{tid, accid, ope})
}
