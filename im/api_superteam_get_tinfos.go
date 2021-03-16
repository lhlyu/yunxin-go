package im

const _API_SUPERTEAM_GET_TINFOS_URL = "https://api.netease.im/nimserver/superteam/getTinfos.action"

type SuperteamGetTinfosParam struct {
	// tid列表，如["3083","3084"]
	// 是否必须: true
	Tids string `json:"tids"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#获取群信息
// 获取群信息
func (y *YunxinIM) ApiSuperteamGetTinfos(tids string) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_GET_TINFOS_URL, SuperteamGetTinfosParam{tids})
}
