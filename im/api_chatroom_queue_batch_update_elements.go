package im

const _API_CHATROOM_QUEUE_BATCH_UPDATE_ELEMENTS_URL = "https://api.netease.im/nimserver/chatroom/queueBatchUpdateElements.action"

type ChatroomQueueBatchUpdateElementsParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 操作者accid,必须是管理员或创建者
	// 是否必须: true
	Operator string `json:"operator"`

	// 更新的key-value对，最大200个，示例：{"k1":"v1","k2":"v2"}
	// 是否必须: true
	Elements string `json:"elements"`

	// true或false,是否需要发送更新通知事件，默认true
	// 是否必须: false
	NeedNotify bool `json:"needNotify,omitempty"`

	// 通知事件扩展字段，长度限制2048
	// 是否必须: false
	NotifyExt string `json:"notifyExt,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%89%B9%E9%87%8F%E6%9B%B4%E6%96%B0%E8%81%8A%E5%A4%A9%E5%AE%A4%E9%98%9F%E5%88%97%E5%85%83%E7%B4%A0
// 批量更新聊天室队列元素
func (y *YunxinIM) ApiChatroomQueueBatchUpdateElements(param *ChatroomQueueBatchUpdateElementsParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUEUE_BATCH_UPDATE_ELEMENTS_URL, param)
}
