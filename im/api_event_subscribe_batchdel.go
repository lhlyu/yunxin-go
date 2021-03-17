package im

const _API_EVENT_SUBSCRIBE_BATCHDEL_URL = "https://api.netease.im/nimserver/event/subscribe/batchdel.action"

type EventSubscribeBatchdelParam struct {
	// 事件订阅人账号
	// 是否必须: true
	Accid string `json:"accid"`

	// 事件类型，固定设置为1，即 eventType=1
	// 是否必须: true
	EventType int `json:"eventType"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%9C%A8%E7%BA%BF%E7%8A%B6%E6%80%81?#%E5%8F%96%E6%B6%88%E5%85%A8%E9%83%A8%E5%9C%A8%E7%BA%BF%E7%8A%B6%E6%80%81%E4%BA%8B%E4%BB%B6%E8%AE%A2%E9%98%85
// 取消全部在线状态事件订阅
func (y *YunxinIM) ApiEventSubscribeBatchdel(accid string, eventType int) *ImResp {
	return y.PostFrom(_API_EVENT_SUBSCRIBE_BATCHDEL_URL, EventSubscribeBatchdelParam{accid, eventType})
}
