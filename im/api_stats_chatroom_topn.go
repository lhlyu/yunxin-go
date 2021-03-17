package im

const _API_STATS_CHATROOM_TOPN_URL = "https://api.netease.im/nimserver/stats/chatroom/topn.action"

type StatsChatroomTopnParam struct {
	// topn值，可选值 1~500，默认值100
	// 是否必须: false
	Topn int `json:"topn,omitempty"`

	// 需要查询的指标所在的时间坐标点，不提供则默认当前时间，单位秒/毫秒皆可
	// 是否必须: false
	Timestamp int64 `json:"timestamp,omitempty"`

	// 统计周期，可选值包括 hour/day, 默认hour
	// 是否必须: false
	Period string `json:"period,omitempty"`

	// 取排序值,可选值 active/enter/message,分别表示按日活排序，进入人次排序和消息数排序， 默认active
	// 是否必须: false
	Orderby string `json:"orderby,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%9F%A5%E8%AF%A2%E8%81%8A%E5%A4%A9%E5%AE%A4%E7%BB%9F%E8%AE%A1%E6%8C%87%E6%A0%87TopN
// 查询聊天室统计指标TopN
func (y *YunxinIM) ApiStatsChatroomTopn(param *StatsChatroomTopnParam) *ImResp {
	return y.PostFrom(_API_STATS_CHATROOM_TOPN_URL, param)
}
