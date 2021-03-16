package im

const _API_FRIEND_UPDATE_URL = "https://api.netease.im/nimserver/friend/update.action"

type FriendUpdateParam struct {
	// 发起者accid
	// 是否必须: true
	Accid string `json:"accid"`

	// 要修改朋友的accid
	// 是否必须: true
	Faccid string `json:"faccid"`

	// 给好友增加备注名，限制长度128，可设置为空字符串
	// 是否必须: false
	Alias string `json:"alias,omitempty"`

	// 修改ex字段，限制长度256，可设置为空字符串
	// 是否必须: false
	Ex string `json:"ex,omitempty"`

	// 修改serverex字段，限制长度256，可设置为空字符串
	// 此字段client端只读，server端读写
	// 是否必须: false
	Serverex string `json:"serverex,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户关系托管?#更新好友相关信息
// 更新好友相关信息
func (y *YunxinIM) ApiFriendUpdate(param *FriendUpdateParam) *ImResp {
	return y.PostFrom(_API_FRIEND_UPDATE_URL, param)
}
