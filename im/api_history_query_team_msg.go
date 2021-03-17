package im

const _API_HISTORY_QUERY_TEAM_MSG_URL = "https://api.netease.im/nimserver/history/queryTeamMsg.action"

type HistoryQueryTeamMsgParam struct {
	// 群id
	// 是否必须: true
	Tid string `json:"tid"`

	// 查询用户对应的accid.
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

	// 1按时间正序排列，2按时间降序排列。其它返回参数414错误。默认是按降序排列，即时间戳最晚的消息排在最前面。
	// 是否必须: false
	Reverse int `json:"reverse,omitempty"`

	// 查询指定的多个消息类型，类型之间用","分割，不设置该参数则查询全部类型消息格式示例： 0,1,2,3
	// 类型支持： 1:图片，2:语音，3:视频，4:地理位置，5:通知，6:文件，10:提示，11:Robot，100:自定义
	// 是否必须: false
	Type string `json:"type,omitempty"`

	// true(默认值)：表示需要检查群是否有效,accid是否为有效的群成员；设置为false则仅检测群是否存在，accid是否曾经为群成员。
	// 是否必须: false
	CheckTeamValid bool `json:"checkTeamValid,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95?#%E7%BE%A4%E8%81%8A%E4%BA%91%E7%AB%AF%E5%8E%86%E5%8F%B2%E6%B6%88%E6%81%AF%E6%9F%A5%E8%AF%A2
// 群聊云端历史消息查询
func (y *YunxinIM) ApiHistoryQueryTeamMsg(param *HistoryQueryTeamMsgParam) *ImResp {
	return y.PostFrom(_API_HISTORY_QUERY_TEAM_MSG_URL, param)
}
