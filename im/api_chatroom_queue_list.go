package im

const _API_CHATROOM_QUEUE_LIST_URL = "https://api.netease.im/nimserver/chatroom/queueList.action"

type ChatroomQueueListParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E6%8E%92%E5%BA%8F%E5%88%97%E5%87%BA%E9%98%9F%E5%88%97%E4%B8%AD%E6%89%80%E6%9C%89%E5%85%83%E7%B4%A0
// 排序列出队列中所有元素
func (y *YunxinIM) ApiChatroomQueueList(roomid int64) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUEUE_LIST_URL, ChatroomQueueListParam{roomid})
}
