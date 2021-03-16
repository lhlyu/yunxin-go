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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（高级群）?#群信息与成员列表查询
// 群信息与成员列表查询
func (y *YunxinIM) ApiTeamQuery(tids string, ope int, ignoreinvalid bool) *ImResp {
	return y.PostFrom(_API_TEAM_QUERY_URL, TeamQueryParam{tids, ope, ignoreinvalid})
}
