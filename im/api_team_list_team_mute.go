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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%BB%84%E7%A6%81%E8%A8%80%E5%88%97%E8%A1%A8
// 获取群组禁言列表
func (y *YunxinIM) ApiTeamListTeamMute(tid string, owner string) *ImResp {
	return y.PostFrom(_API_TEAM_LIST_TEAM_MUTE_URL, TeamListTeamMuteParam{tid, owner})
}
