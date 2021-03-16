package im

const _API_USER_LIST_BLACK_AND_MUTE_LIST_URL = "https://api.netease.im/nimserver/user/listBlackAndMuteList.action"

type UserListBlackAndMuteListParam struct {
	// 用户帐号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户关系托管?#查看指定用户的黑名单和静音列表
// 查看指定用户的黑名单和静音列表
func (y *YunxinIM) ApiUserListBlackAndMuteList(accid string) *ImResp {
	return y.PostFrom(_API_USER_LIST_BLACK_AND_MUTE_LIST_URL, UserListBlackAndMuteListParam{accid})
}
