package im

const _API_USER_REFRESH_TOKEN_URL = "https://api.netease.im/nimserver/user/refreshToken.action"

type UserRefreshTokenParam struct {
	// 网易云信IM账号，最大长度32字符，必须保证一个
	// APP内唯一
	// 是否必须: true
	Accid string `json:"accid"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/网易云通信ID?#重置网易云信IM token
// 重置网易云信IM token
func (y *YunxinIM) ApiUserRefreshToken(accid string) *ImResp {
	return y.PostFrom(_API_USER_REFRESH_TOKEN_URL, UserRefreshTokenParam{accid})
}
