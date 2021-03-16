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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户关系托管?#加好友
// 加好友
func (y *YunxinIM) ApiFriendAdd(param *FriendAddParam) *ImResp {
	return y.PostFrom(_API_FRIEND_ADD_URL, param)
}
