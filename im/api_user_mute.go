package im

const _API_USER_MUTE_URL = "https://api.netease.im/nimserver/user/mute.action"

type UserMuteParam struct {
	// 用户帐号
	// 是否必须: true
	Accid string `json:"accid"`

	// 是否全局禁言：
	// true：全局禁言，false:取消全局禁言
	// 是否必须: true
	Mute bool `json:"mute"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E8%AE%BE%E7%BD%AE?#%E8%B4%A6%E5%8F%B7%E5%85%A8%E5%B1%80%E7%A6%81%E8%A8%80
// 账号全局禁言
func (y *YunxinIM) ApiUserMute(accid string, mute bool) *ImResp {
	return y.PostFrom(_API_USER_MUTE_URL, UserMuteParam{accid, mute})
}
