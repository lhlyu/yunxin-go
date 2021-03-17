package im

const _API_CHATROOM_QUEUE_INIT_URL = "https://api.netease.im/nimserver/chatroom/queueInit.action"

type ChatroomQueueInitParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 队列长度限制，0~1000
	// 是否必须: true
	SizeLimit int64 `json:"sizeLimit"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%88%9D%E5%A7%8B%E5%8C%96%E9%98%9F%E5%88%97
// 初始化队列
func (y *YunxinIM) ApiChatroomQueueInit(roomid int64, sizeLimit int64) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUEUE_INIT_URL, ChatroomQueueInitParam{roomid, sizeLimit})
}
