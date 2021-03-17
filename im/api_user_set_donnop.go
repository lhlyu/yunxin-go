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

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E8%AE%BE%E7%BD%AE?#%E8%AE%BE%E7%BD%AE%E6%A1%8C%E9%9D%A2%E7%AB%AF%E5%9C%A8%E7%BA%BF%E6%97%B6%EF%BC%8C%E7%A7%BB%E5%8A%A8%E7%AB%AF%E6%98%AF%E5%90%A6%E4%B8%8D%E6%8E%A8%E9%80%81
// 设置桌面端在线时，移动端是否需要推送
func (y *YunxinIM) ApiUserSetDonnop(accid string, donnopOpen string) *ImResp {
	return y.PostFrom(_API_USER_SET_DONNOP_URL, UserSetDonnopParam{accid, donnopOpen})
}
