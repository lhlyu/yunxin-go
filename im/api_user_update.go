package im

const _API_USER_UPDATE_URL = "https://api.netease.im/nimserver/user/update.action"

type UserUpdateParam struct {
	// 网易云信IM账号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`

	// 该参数已不建议使用。
	// 是否必须: false
	Props string `json:"props,omitempty"`

	// 网易云信IM账号可以指定登录token值，最大长度128字符
	// 是否必须: false
	Token string `json:"token,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/网易云通信ID?#更新网易云信IM token
// 更新网易云信IM token
func (y *YunxinIM) ApiUserUpdate(accid string, props string, token string) *ImResp {
	return y.PostFrom(_API_USER_UPDATE_URL, UserUpdateParam{accid, props, token})
}
