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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E5%88%A0%E9%99%A4%E5%A5%BD%E5%8F%8B
// 删除好友
func (y *YunxinIM) ApiFriendDelete(accid string, faccid string, isDeleteAlias bool) *ImResp {
	return y.PostFrom(_API_FRIEND_DELETE_URL, FriendDeleteParam{accid, faccid, isDeleteAlias})
}
