package im

const _API_SUPERTEAM_GET_TLISTS_URL = "https://api.netease.im/nimserver/superteam/getTlists.action"

type SuperteamGetTlistsParam struct {
	// 云信服务器产生，群唯一标识，创建群时会返回，最大长度128字符
	// 是否必须: true
	Tid string `json:"tid"`

	// 时间戳，单位毫秒，查询的时间起点。
	// 是否必须: true
	Timetag string `json:"timetag"`

	// 本次查询的条数上限(最多100条)，小于等于0，或者大于100，会提示参数错误
	// 是否必须: true
	Limit string `json:"limit"`

	// 1:按时间正序排列，2:按时间降序排列。其它会提示参数错误。默认是1按时间正序排列
	// 是否必须: false
	Reverse string `json:"reverse,omitempty"`
}

// doc: https://dev.yunxin.163.com/docs/product/IM%E5%8D%B3%E6%97%B6%E9%80%9A%E8%AE%AF/%E6%9C%8D%E5%8A%A1%E7%AB%AFAPI%E6%96%87%E6%A1%A3/%E7%BE%A4%E7%BB%84%E5%8A%9F%E8%83%BD%EF%BC%88%E8%B6%85%E5%A4%A7%E7%BE%A4%EF%BC%89?#%E8%8E%B7%E5%8F%96%E7%BE%A4%E6%88%90%E5%91%98%E4%BF%A1%E6%81%AF
// 获取群成员信息
func (y *YunxinIM) ApiSuperteamGetTlists(param *SuperteamGetTlistsParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_GET_TLISTS_URL, param)
}
