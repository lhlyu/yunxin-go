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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%BB%84%E5%B7%B2%E8%AF%BB%E6%B6%88%E6%81%AF%E7%9A%84%E5%B7%B2%E8%AF%BB%E8%AF%A6%E6%83%85%E4%BF%A1%E6%81%AF
// 获取群组已读消息的已读详情信息
func (y *YunxinIM) ApiTeamGetMarkReadInfo(param *TeamGetMarkReadInfoParam) *ImResp {
	return y.PostFrom(_API_TEAM_GET_MARK_READ_INFO_URL, param)
}
