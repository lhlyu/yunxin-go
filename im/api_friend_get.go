package im

const _API_FRIEND_GET_URL = "https://api.netease.im/nimserver/friend/get.action"

type FriendGetParam struct {
	// 发起者accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 更新时间戳，接口返回该时间戳之后有更新的好友列表
	// 是否必须: true
	Updatetime int64 `json:"updatetime"`

	// 【Deprecated】定义同updatetime
	// 是否必须: false
	Createtime int64 `json:"createtime,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E8%8E%B7%E5%8F%96%E5%A5%BD%E5%8F%8B%E5%85%B3%E7%B3%BB
// 获取好友关系
func (y *YunxinIM) ApiFriendGet(accid string, updatetime int64, createtime int64) *ImResp {
	return y.PostFrom(_API_FRIEND_GET_URL, FriendGetParam{accid, updatetime, createtime})
}
