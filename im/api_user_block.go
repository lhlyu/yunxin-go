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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/网易云通信ID?#封禁网易云信IM账号
// 封禁网易云信IM账号
func (y *YunxinIM) ApiUserBlock(accid string, needkick string, kicknotifyext string) *ImResp {
	return y.PostFrom(_API_USER_BLOCK_URL, UserBlockParam{accid, needkick, kicknotifyext})
}
