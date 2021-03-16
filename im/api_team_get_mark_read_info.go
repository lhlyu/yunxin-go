package im

const _API_TEAM_GET_MARK_READ_INFO_URL = "https://api.netease.im/nimserver/team/getMarkReadInfo.action"

type TeamGetMarkReadInfoParam struct {
	// 群id，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid int64 `json:"tid"`

	// 发送群已读业务消息时服务器返回的消息ID
	// 是否必须: true
	Msgid int64 `json:"msgid"`

	// 消息发送者账号
	// 是否必须: true
	FromAccid string `json:"fromAccid"`

	// 是否返回已读、未读成员的accid列表，默认为false
	// 是否必须: false
	Snapshot bool `json:"snapshot,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#获取群组已读消息的已读详情信息
// 获取群组已读消息的已读详情信息
func (y *YunxinIM) ApiTeamGetMarkReadInfo(param *TeamGetMarkReadInfoParam) *ImResp {
	return y.PostFrom(_API_TEAM_GET_MARK_READ_INFO_URL, param)
}
