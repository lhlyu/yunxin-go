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

	// 禁言类型
	// 0:解除禁言，1:禁言普通成员
	// 3:禁言整个群(包括群主)
	// 是否必须: false
	MuteType int `json:"muteType,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#将群组整体禁言
// 将群组整体禁言
func (y *YunxinIM) ApiTeamMuteTlistAll(param *TeamMuteTlistAllParam) *ImResp {
	return y.PostFrom(_API_TEAM_MUTE_TLIST_ALL_URL, param)
}
