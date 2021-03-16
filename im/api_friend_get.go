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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户关系托管?#获取好友关系
// 获取好友关系
func (y *YunxinIM) ApiFriendGet(accid string, updatetime int64, createtime int64) *ImResp {
	return y.PostFrom(_API_FRIEND_GET_URL, FriendGetParam{accid, updatetime, createtime})
}
