package im

const _API_HISTORY_QUERY_USER_EVENTS_URL = "https://api.netease.im/nimserver/history/queryUserEvents.action"

type HistoryQueryUserEventsParam struct {
	// 要查询用户的accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 开始时间，毫秒级
	// 是否必须: true
	Begintime string `json:"begintime"`

	// 截止时间，毫秒级
	// 是否必须: true
	Endtime string `json:"endtime"`

	// 本次查询的消息条数上限(最多100条),小于等于0，或者大于100，会提示参数错误
	// 是否必须: true
	Limit int `json:"limit"`

	// 1按时间正序排列，2按时间降序排列。其它返回参数414错误。默认是按降序排列
	// 是否必须: false
	Reverse int `json:"reverse,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95?#%E7%94%A8%E6%88%B7%E7%99%BB%E5%BD%95%E7%99%BB%E5%87%BA%E4%BA%8B%E4%BB%B6%E8%AE%B0%E5%BD%95%E6%9F%A5%E8%AF%A2
// 用户登录登出事件记录查询
func (y *YunxinIM) ApiHistoryQueryUserEvents(param *HistoryQueryUserEventsParam) *ImResp {
	return y.PostFrom(_API_HISTORY_QUERY_USER_EVENTS_URL, param)
}
