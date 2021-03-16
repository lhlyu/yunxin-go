package im

const _API_USER_UNBLOCK_URL = "https://api.netease.im/nimserver/user/unblock.action"

type UserUnblockParam struct {
	// 网易云信IM账号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/网易云通信ID?#解禁网易云信IM账号
// 解禁网易云信IM账号
func (y *YunxinIM) ApiUserUnblock(accid string) *ImResp {
	return y.PostFrom(_API_USER_UNBLOCK_URL, UserUnblockParam{accid})
}
