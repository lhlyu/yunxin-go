package im

const _API_TEAM_QUERY_DETAIL_URL = "https://api.netease.im/nimserver/team/queryDetail.action"

type TeamQueryDetailParam struct {
	// 群id，群唯一标识，创建群时会返回
	// 是否必须: true
	Tid int64 `json:"tid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E7%BB%84%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF
// 获取群组详细信息
func (y *YunxinIM) ApiTeamQueryDetail(tid int64) *ImResp {
	return y.PostFrom(_API_TEAM_QUERY_DETAIL_URL, TeamQueryDetailParam{tid})
}
