package im

const _API_HISTORY_QUERY_BROADCAST_MSG_URL = "https://api.netease.im/nimserver/history/queryBroadcastMsg.action"

type HistoryQueryBroadcastMsgParam struct {
	// 查询的起始ID，0表示查询最近的limit条。默认0。
	// 是否必须: false
	BroadcastId int64 `json:"broadcastId,omitempty"`

	// 查询的条数，最大100。默认100。
	// 是否必须: false
	Limit int `json:"limit,omitempty"`

	// 查询的类型，1表示所有，2表示查询存离线的，3表示查询不存离线的。默认1。
	// 是否必须: false
	Type int64 `json:"type,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95?#%E6%89%B9%E9%87%8F%E6%9F%A5%E8%AF%A2%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
// 批量查询广播消息
func (y *YunxinIM) ApiHistoryQueryBroadcastMsg(broadcastId int64, limit int, kind int64) *ImResp {
	return y.PostFrom(_API_HISTORY_QUERY_BROADCAST_MSG_URL, HistoryQueryBroadcastMsgParam{broadcastId, limit, kind})
}
