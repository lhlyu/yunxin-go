package im

const _API_CHATROOM_QUEUE_POLL_URL = "https://api.netease.im/nimserver/chatroom/queuePoll.action"

type ChatroomQueuePollParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 目前元素的elementKey,长度限制128字符，不填表示取出头上的第一个
	// 是否必须: false
	Key string `json:"key,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E4%BB%8E%E9%98%9F%E5%88%97%E4%B8%AD%E5%8F%96%E5%87%BA%E5%85%83%E7%B4%A0
// 从队列中取出元素
func (y *YunxinIM) ApiChatroomQueuePoll(roomid int64, key string) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUEUE_POLL_URL, ChatroomQueuePollParam{roomid, key})
}
