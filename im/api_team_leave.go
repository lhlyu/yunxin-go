package im

const _API_TEAM_LEAVE_URL = "https://api.netease.im/nimserver/team/leave.action"

type TeamLeaveParam struct {
	// 网易云信服务器产生，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid string `json:"tid"`

	// 退群的accid
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#主动退群
// 主动退群
func (y *YunxinIM) ApiTeamLeave(tid string, accid string) *ImResp {
	return y.PostFrom(_API_TEAM_LEAVE_URL, TeamLeaveParam{tid, accid})
}
