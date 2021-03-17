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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E6%9B%B4%E6%96%B0%E5%A5%BD%E5%8F%8B%E7%9B%B8%E5%85%B3%E4%BF%A1%E6%81%AF
// 更新好友相关信息
func (y *YunxinIM) ApiFriendUpdate(param *FriendUpdateParam) *ImResp {
	return y.PostFrom(_API_FRIEND_UPDATE_URL, param)
}
