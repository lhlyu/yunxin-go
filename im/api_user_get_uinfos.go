package im

const _API_USER_GET_UINFOS_URL = "https://api.netease.im/nimserver/user/getUinfos.action"

type UserGetUinfosParam struct {
	// 用户帐号（例如：JSONArray对应的accid串，如：["zhangsan"]，如果解析出错，会报414）（一次查询最多为200）
	// 是否必须: true
	Accids string `json:"accids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87?#%E8%8E%B7%E5%8F%96%E7%94%A8%E6%88%B7%E5%90%8D%E7%89%87
// 获取用户名片
func (y *YunxinIM) ApiUserGetUinfos(accids string) *ImResp {
	return y.PostFrom(_API_USER_GET_UINFOS_URL, UserGetUinfosParam{accids})
}
