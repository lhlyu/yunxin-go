package im

const _API_USER_LIST_BLACK_AND_MUTE_LIST_URL = "https://api.netease.im/nimserver/user/listBlackAndMuteList.action"

type UserListBlackAndMuteListParam struct {
	// 用户帐号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%85%B3%E7%B3%BB%E6%89%98%E7%AE%A1?#%E6%9F%A5%E7%9C%8B%E6%8C%87%E5%AE%9A%E7%94%A8%E6%88%B7%E7%9A%84%E9%BB%91%E5%90%8D%E5%8D%95%E5%92%8C%E9%9D%99%E9%9F%B3%E5%88%97%E8%A1%A8
// 查看指定用户的黑名单和静音列表
func (y *YunxinIM) ApiUserListBlackAndMuteList(accid string) *ImResp {
	return y.PostFrom(_API_USER_LIST_BLACK_AND_MUTE_LIST_URL, UserListBlackAndMuteListParam{accid})
}
