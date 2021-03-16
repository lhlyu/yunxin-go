package im

const _API_TEAM_LIST_TEAM_MUTE_URL = "https://api.netease.im/nimserver/team/listTeamMute.action"

type TeamListTeamMuteParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主的accid
	// 是否必须: true
	Owner string `json:"owner"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#获取群组禁言列表
// 获取群组禁言列表
func (y *YunxinIM) ApiTeamListTeamMute(tid string, owner string) *ImResp {
	return y.PostFrom(_API_TEAM_LIST_TEAM_MUTE_URL, TeamListTeamMuteParam{tid, owner})
}
