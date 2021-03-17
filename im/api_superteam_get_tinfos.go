package im

const _API_SUPERTEAM_GET_TINFOS_URL = "https://api.netease.im/nimserver/superteam/getTinfos.action"

type SuperteamGetTinfosParam struct {
	// tid列表，如["3083","3084"]
	// 是否必须: true
	Tids string `json:"tids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E4%BF%A1%E6%81%AF
// 获取群信息
func (y *YunxinIM) ApiSuperteamGetTinfos(tids string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_GET_TINFOS_URL, SuperteamGetTinfosParam{tids})
}
