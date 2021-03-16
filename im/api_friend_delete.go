package im

const _API_FRIEND_DELETE_URL = "https://api.netease.im/nimserver/friend/delete.action"

type FriendDeleteParam struct {
	// 发起者accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 要删除朋友的accid
	// 是否必须: true
	Faccid string `json:"faccid"`

	// 是否需要删除备注信息
	// 默认false:不需要，true:需要
	// 是否必须: false
	IsDeleteAlias bool `json:"isDeleteAlias,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户关系托管?#删除好友
// 删除好友
func (y *YunxinIM) ApiFriendDelete(accid string, faccid string, isdeletealias bool) *ImResp {
	return y.PostFrom(_API_FRIEND_DELETE_URL, FriendDeleteParam{accid, faccid, isdeletealias})
}
