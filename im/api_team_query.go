package im

const _API_TEAM_QUERY_URL = "https://api.netease.im/nimserver/team/query.action"

type TeamQueryParam struct {
	// 群id列表，如["3083","3084"]
	// 是否必须: true
	Tids string `json:"tids"`

	// 1表示带上群成员列表，0表示不带群成员列表，只返回群信息
	// 是否必须: true
	Ope int `json:"ope"`

	// 是否忽略无效的tid，默认为false。设置为true时将忽略无效tid，并在响应结果中返回无效的tid
	// 是否必须: false
	IgnoreInvalid bool `json:"ignoreInvalid,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E9%AB%98%E7%BA%A7%E7%BE%A4%EF%BC%89?#%E7%BE%A4%E4%BF%A1%E6%81%AF%E4%B8%8E%E6%88%90%E5%91%98%E5%88%97%E8%A1%A8%E6%9F%A5%E8%AF%A2
// 群信息与成员列表查询
func (y *YunxinIM) ApiTeamQuery(tids string, ope int, ignoreInvalid bool) *ImResp {
	return y.PostFrom(_API_TEAM_QUERY_URL, TeamQueryParam{tids, ope, ignoreInvalid})
}
