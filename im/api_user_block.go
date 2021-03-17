package im

const _API_USER_BLOCK_URL = "https://api.netease.im/nimserver/user/block.action"

type UserBlockParam struct {
	// 网易云信IM账号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`

	// 是否踢掉被禁用户，true或false，默认false
	// 是否必须: false
	Needkick string `json:"needkick,omitempty"`

	// 踢人时的扩展字段，SDK版本需要大于等于v7.7.0
	// 是否必须: false
	KickNotifyExt string `json:"kickNotifyExt,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BD%91%E6%98%93%E4%BA%91%E9%80%9A%E4%BF%A1ID?#%E5%B0%81%E7%A6%81%E7%BD%91%E6%98%93%E4%BA%91%E4%BF%A1IM%E8%B4%A6%E5%8F%B7
// 封禁网易云信IM账号
func (y *YunxinIM) ApiUserBlock(accid string, needkick string, kickNotifyExt string) *ImResp {
	return y.PostFrom(_API_USER_BLOCK_URL, UserBlockParam{accid, needkick, kickNotifyExt})
}
