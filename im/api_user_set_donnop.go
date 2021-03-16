package im

const _API_USER_SET_DONNOP_URL = "https://api.netease.im/nimserver/user/setDonnop.action"

type UserSetDonnopParam struct {
	// 用户帐号
	// 是否必须: true
	Accid string `json:"accid"`

	// 桌面端在线时，移动端是否不推送：
	// true:移动端不需要推送，false:移动端需要推送
	// 是否必须: true
	DonnopOpen string `json:"donnopOpen"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户设置?#设置桌面端在线时，移动端是否不推送
// 设置桌面端在线时，移动端是否需要推送
func (y *YunxinIM) ApiUserSetDonnop(accid string, donnopopen string) *ImResp {
	return y.PostFrom(_API_USER_SET_DONNOP_URL, UserSetDonnopParam{accid, donnopopen})
}
