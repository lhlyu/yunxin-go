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

// doc: https://dev.yunxin.163.com/docs/product/IM即时通讯/服务端API文档/群组功能（超大群）?#获取群成员信息
// 获取群成员信息
func (y *YunxinIM) ApiSuperteamGetTlists(param *SuperteamGetTlistsParam) *ImResp {
	return y.PostFrom(_API_SUPERTEAM_GET_TLISTS_URL, param)
}
