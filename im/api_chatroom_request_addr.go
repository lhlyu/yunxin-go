package im

const _API_CHATROOM_REQUEST_ADDR_URL = "https://api.netease.im/nimserver/chatroom/requestAddr.action"

type ChatroomRequestAddrParam struct {
	// 聊天室id
	// 是否必须: true
	Roomid int64 `json:"roomid"`

	// 进入聊天室的账号
	// 是否必须: true
	Accid string `json:"accid"`

	// 1:weblink（客户端为web端时使用）; 2:commonlink（客户端为非web端时使用）;3:wechatlink(微信小程序使用), 默认1
	// 是否必须: false
	Clienttype int `json:"clienttype,omitempty"`

	// 客户端ip，传此参数时，会根据用户ip所在地区，返回合适的地址
	// 是否必须: false
	Clientip string `json:"clientip,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E8%81%8A%E5%A4%A9%E5%AE%A4?#%E8%AF%B7%E6%B1%82%E8%81%8A%E5%A4%A9%E5%AE%A4%E5%9C%B0%E5%9D%80
// 请求聊天室地址
func (y *YunxinIM) ApiChatroomRequestAddr(param *ChatroomRequestAddrParam) *ImResp {
	return y.PostFrom(_API_CHATROOM_REQUEST_ADDR_URL, param)
}
