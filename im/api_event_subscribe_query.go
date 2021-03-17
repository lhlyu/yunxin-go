package im

const _API_EVENT_SUBSCRIBE_QUERY_URL = "https://api.netease.im/nimserver/event/subscribe/query.action"

type EventSubscribeQueryParam struct {
	// 事件订阅人账号
	// 是否必须: true
	Accid string `json:"accid"`

	// 事件类型，固定设置为1，即 eventType=1
	// 是否必须: true
	EventType int `json:"eventType"`

	// 被订阅人的账号列表，最多100个账号，JSONArray格式。示例：["pub_user1","pub_user2"]
	// 是否必须: true
	PublisherAccids string `json:"publisherAccids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E5%9C%A8%E7%BA%BF%E7%8A%B6%E6%80%81?#%E6%9F%A5%E8%AF%A2%E5%9C%A8%E7%BA%BF%E7%8A%B6%E6%80%81%E4%BA%8B%E4%BB%B6%E8%AE%A2%E9%98%85%E5%85%B3%E7%B3%BB
// 查询在线状态事件订阅关系
func (y *YunxinIM) ApiEventSubscribeQuery(accid string, eventType int, publisherAccids string) *ImResp {
	return y.PostFrom(_API_EVENT_SUBSCRIBE_QUERY_URL, EventSubscribeQueryParam{accid, eventType, publisherAccids})
}
