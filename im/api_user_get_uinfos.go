package im

const _API_USER_GET_UINFOS_URL = "https://api.netease.im/nimserver/user/getUinfos.action"

type UserGetUinfosParam struct {
	// 用户帐号（例如：JSONArray对应的accid串，如：["zhangsan"]，如果解析出错，会报414）（一次查询最多为200）
	// 是否必须: true
	Accids string `json:"accids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/用户名片?#获取用户名片
// 获取用户名片
func (y *YunxinIM) ApiUserGetUinfos(accids string) *ImResp {
	return y.PostFrom(_API_USER_GET_UINFOS_URL, UserGetUinfosParam{accids})
}
