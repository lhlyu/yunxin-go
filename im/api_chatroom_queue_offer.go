package im

const _API_CHATROOM_QUEUE_OFFER_URL = "https://api.netease.im/nimserver/chatroom/queueOffer.action"

type ChatroomQueueOfferParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// elementKey,新元素的UniqKey,长度限制128字符
	// 是否必须: true
	Key string `json:"key"`

	// elementValue,新元素内容，长度限制4096字符
	// 是否必须: true
	Value string `json:"value"`

	// 提交这个新元素的操作者accid，默认为该聊天室的创建者，若operator对应的帐号不存在，会返回404错误。
	// 若指定的operator不在线，则添加元素成功后的通知事件中的操作者默认为聊天室的创建者；若指定的operator在线，则通知事件的操作者为operator。
	// 是否必须: false
	Operator string `json:"operator,omitempty"`

	// 这个新元素的提交者operator的所有聊天室连接在从该聊天室掉线或者离开该聊天室的时候，提交的元素是否需要删除。
	// true：需要删除；false：不需要删除。默认false。
	// 当指定该参数为true时，若operator当前不在该聊天室内，则会返回403错误。
	// 是否必须: false
	Transient string `json:"transient,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E5%BE%80%E8%81%8A%E5%A4%A9%E5%AE%A4%E6%9C%89%E5%BA%8F%E9%98%9F%E5%88%97%E4%B8%AD%E6%96%B0%E5%8A%A0%E6%88%96%E6%9B%B4%E6%96%B0%E5%85%83%E7%B4%A0
// 往聊天室有序队列中新加或更新元素
func (y *YunxinIM) ApiChatroomQueueOffer(param *ChatroomQueueOfferParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_QUEUE_OFFER_URL, param)
}
