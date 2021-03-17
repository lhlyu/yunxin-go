package im

const _API_HISTORY_QUERY_CHATROOM_MSG_URL = "https://api.netease.im/nimserver/history/queryChatroomMsg.action"

type HistoryQueryChatroomMsgParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 用户账号
	// 是否必须: true
	Accid string `json:"accid"`

	// 查询的时间戳锚点，13位。reverse=1时timetag为起始时间戳，reverse=2时timetag为终止时间戳
	// 是否必须: true
	Timetag int64 `json:"timetag"`

	// 本次查询的消息条数上限(最多200条),小于等于0，或者大于200，会提示参数错误
	// 是否必须: true
	Limit int `json:"limit"`

	// 1按时间正序排列，2按时间降序排列。其它返回参数414错误。默认是2按时间降序排列
	// 是否必须: false
	Reverse int `json:"reverse,omitempty"`

	// 查询指定的多个消息类型，类型之间用","分割，不设置该参数则查询全部类型消息。
	// 格式示例： 0,1,2,3
	// 支持的消息类型：0:文本，1:图片，2:语音，3:视频，4:地理位置，5:通知，6:文件，10:提示，11:智能机器人消息，100:自定义消息。用英文逗号分隔。
	// 是否必须: false
	Type string `json:"type,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95?#%E8%81%8A%E5%A4%A9%E5%AE%A4%E4%BA%91%E7%AB%AF%E5%8E%86%E5%8F%B2%E6%B6%88%E6%81%AF%E6%9F%A5%E8%AF%A2
// 聊天室云端历史消息查询
func (y *YunxinIM) ApiHistoryQueryChatroomMsg(param *HistoryQueryChatroomMsgParam) *ImResp {
	return y.PostFrom(_API_HISTORY_QUERY_CHATROOM_MSG_URL, param)
}
