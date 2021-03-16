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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户设置?#账号全局禁言
// 账号全局禁言
func (y *YunxinIM) ApiUserMute(accid string, mute bool) *ImResp {
	return y.PostFrom(_API_USER_MUTE_URL, UserMuteParam{accid, mute})
}
