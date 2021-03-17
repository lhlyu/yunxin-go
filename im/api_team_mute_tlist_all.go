package im

const _API_TEAM_MUTE_TLIST_ALL_URL = "https://api.netease.im/nimserver/team/muteTlistAll.action"

type TeamMuteTlistAllParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 群主的accid
	// 是否必须: true
	Owner string `json:"owner"`

	// true:禁言，false:解禁(mute和muteType至少提供一个，都提供时按mute处理)
	// 是否必须: false
	Mute string `json:"mute,omitempty"`

	// 禁言类型 0:解除禁言，1:禁言普通成员 3:禁言整个群(包括群主)
	// 是否必须: false
	MuteType int `json:"muteType,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E5%B0%86%E7%BE%A4%E7%BB%84%E6%95%B4%E4%BD%93%E7%A6%81%E8%A8%80
// 将群组整体禁言
func (y *YunxinIM) ApiTeamMuteTlistAll(param *TeamMuteTlistAllParam) *ImResp {
	return y.PostFrom(_API_TEAM_MUTE_TLIST_ALL_URL, param)
}
