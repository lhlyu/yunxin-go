package im

const _API_HISTORY_QUERY_BROADCAST_MSG_BY_ID_URL = "https://api.netease.im/nimserver/history/queryBroadcastMsgById.action"

type HistoryQueryBroadcastMsgByIdParam struct {
	// 广播消息ID，应用内唯一。
	// 是否必须: true
	BroadcastId int64 `json:"broadcastId"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%8E%86%E5%8F%B2%E8%AE%B0%E5%BD%95?#%E6%9F%A5%E8%AF%A2%E5%8D%95%E6%9D%A1%E5%B9%BF%E6%92%AD%E6%B6%88%E6%81%AF
// 查询单条广播消息
func (y *YunxinIM) ApiHistoryQueryBroadcastMsgById(broadcastId int64) *ImResp {
	return y.PostFrom(_API_HISTORY_QUERY_BROADCAST_MSG_BY_ID_URL, HistoryQueryBroadcastMsgByIdParam{broadcastId})
}
