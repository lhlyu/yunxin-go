package im

const _API_FRIEND_ADD_URL = "https://api.netease.im/nimserver/friend/add.action"

type FriendAddParam struct {
	// 加好友发起者accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 加好友接收者accid
	// 是否必须: true
	Faccid string `json:"faccid"`

	// 1直接加好友，2请求加好友，3同意加好友，4拒绝加好友
	// 是否必须: true
	Type int `json:"type"`

	// 加好友对应的请求消息，第三方组装，最长256字符
	// 是否必须: false
	Msg string `json:"msg,omitempty"`

	// 服务器端扩展字段，限制长度256
	// 此字段client端只读，server端读写
	// 是否必须: false
	Serverex string `json:"serverex,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E5%8A%A0%E5%A5%BD%E5%8F%8B
// 加好友
func (y *YunxinIM) ApiFriendAdd(param *FriendAddParam) *ImResp {
	return y.PostFrom(_API_FRIEND_ADD_URL, param)
}
