package im

const _API_CHATROOM_QUEUE_DROP_URL = "https://api.netease.im/nimserver/chatroom/queueDrop.action"

type ChatroomQueueDropParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%88%A0%E9%99%A4%E6%B8%85%E7%90%86%E6%95%B4%E4%B8%AA%E9%98%9F%E5%88%97
// 删除清理整个队列
func (y *YunxinIM) ApiChatroomQueueDrop(roomid int64) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUEUE_DROP_URL, ChatroomQueueDropParam{roomid})
}
