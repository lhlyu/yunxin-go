package im

const _API_TEAM_QUERY_DETAIL_URL = "https://api.netease.im/nimserver/team/queryDetail.action"

type TeamQueryDetailParam struct {
	// 群id，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid int64 `json:"tid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#获取群组详细信息
// 获取群组详细信息
func (y *YunxinIM) ApiTeamQueryDetail(tid int64) *ImResp {
	return y.PostFrom(_API_TEAM_QUERY_DETAIL_URL, TeamQueryDetailParam{tid})
}
